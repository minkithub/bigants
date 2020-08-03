package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/knq/snaker"
)

func main() {

	in := flag.String("in", "schema.graphql", "target schema")
	out := flag.String("out", "pgmg_gen.go", "output file path")
	flag.Parse()

	b, err := ioutil.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}

	schema := graphql.MustParseSchema(string(b), nil)
	insp := schema.Inspect()
	allTypes := insp.Types()

	ss := []source{}
	includeDeprecated := struct{ IncludeDeprecated bool }{true}

	for _, t := range allTypes {
		var s source

		if strings.HasPrefix(*t.Name(), "__") {
			continue
		}
		fields := t.Fields(&includeDeprecated)
		inputFields := t.InputFields()

		if fields != nil {
			for _, f := range *fields {
				args := f.Args()
				if args == nil || len(args) == 0 {
					continue
				}
				s = source{
					Name:  *t.Name() + snaker.ForceCamelIdentifier(f.Name()) + "Args",
					Props: []*prop{},
				}
				for _, iv := range args {
					s.Props = append(s.Props, inputValueToProp(iv))
				}
				ss = append(ss, s)
			}
		}
		if inputFields != nil {
			s = source{Name: *t.Name(), Props: []*prop{}}
			for _, iv := range *inputFields {
				s.Props = append(s.Props, inputValueToProp(iv))
			}
			ss = append(ss, s)
		}

	}

	writeGoCode(*out, ss)

}

func inputValueToProp(f *introspection.InputValue) *prop {
	// log.Printf("Visiting %s", f.Name())
	return &prop{
		GoTypeName:     renderGoType(f.Type()),
		Name:           snaker.ForceCamelIdentifier(f.Name()),
		SerializedName: snaker.CamelToSnake(f.Name()),
		FieldName:      f.Name(),
	}
}

func fieldToProp(f *introspection.Field) *prop {
	// log.Printf("Visiting %s", f.Name())
	return &prop{
		GoTypeName:     renderGoType(f.Type()),
		Name:           snaker.ForceCamelIdentifier(f.Name()),
		SerializedName: snaker.CamelToSnake(f.Name()),
		FieldName:      f.Name(),
	}
}

func renderGoType(t *introspection.Type) string {
	if t.Kind() == "OBJECT" {
		return "*" + *t.Name()
	} else if t.Kind() == "INPUT_OBJECT" {
		return "*" + *t.Name()

	} else if t.Kind() == "INTERFACE" {
		return *t.Name()

	} else if t.Kind() == "SCALAR" {
		return "*" + gqlScalarToGoType(*t.Name())

	} else if t.Kind() == "ENUM" {
		return "*string"

	} else if t.Kind() == "NON_NULL" {
		ofType := t.OfType()

		if t.Kind() == "OBJECT" {
			return *t.Name()

		} else if t.Kind() == "INPUT_OBJECT" {
			return *t.Name()

		} else if ofType.Kind() == "SCALAR" {
			return gqlScalarToGoType(*ofType.Name())

		} else if t.Kind() == "ENUM" {
			return "string"

		} else if ofType.Kind() == "LIST" {
			return "[]" + renderGoType(ofType.OfType())

		}
		return *ofType.Name()

	} else if t.Kind() == "LIST" {
		return "*[]" + renderGoType(t.OfType())
	}
	log.Fatal(fmt.Errorf("Boom: %s (%s)", *t.Name(), t.Kind()))
	return ""
}

func gqlScalarToGoType(gqlScalar string) string {
	if gqlScalar == "Boolean" {
		return "bool"
	} else if gqlScalar == "String" {
		return "string"
	} else if gqlScalar == "Float" {
		return "float64"
	} else if gqlScalar == "Int" {
		return "int32"
	} else if gqlScalar == "ID" {
		return "graphql.ID"
	} else {
		log.Fatal("Invalid Type: " + gqlScalar)
	}
	return ""
}

type source struct {
	Name  string
	Props []*prop
}

type prop struct {
	Name           string
	SerializedName string
	GoTypeName     string
	FieldName      string
}

func writeGoCode(out string, ss []source) {
	log.Printf("Writing %d input types\n", len(ss))

	var res bytes.Buffer
	if err := tmpl.ExecuteTemplate(&res, "inputgen", ss); err != nil {
		log.Fatal(err)
	}
	code := res.Bytes()
	formattedSrc, err := format.Source(code)
	if err != nil {
		print(string(code))
		log.Fatal(err)
	}
	err = ioutil.WriteFile(out, formattedSrc, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
}
