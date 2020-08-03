package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

func main() {
	var source []string
	var err error

	for _, glob := range []string{
		"root.graphql",
		"nodes/*.graphql",
		"mutations/*.graphql",
		"common.graphql",
	} {
		if source, err = appendGlob(source, glob); err != nil {
			log.Fatal(err)
		}
	}

	merged := strings.Join(source, "\n")

	parsedSchema, err := graphql.ParseSchema(merged, nil, graphql.UseStringDescriptions())
	if err != nil {
		log.Fatalf("%s: \n%s", err.Error(), merged)
	}
	s := schema{parsedSchema}

	schemaJSON, _ := json.Marshal(map[string]interface{}{"data": s})

	if err = ioutil.WriteFile("schema.graphql", []byte(merged), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("schema.json", schemaJSON, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

type schema struct{ *graphql.Schema }

func (s schema) MarshalJSON() ([]byte, error) { return s.Schema.ToJSON() }

func appendGlob(source []string, glob string) ([]string, error) {
	files, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	var b []byte
	for _, file := range files {
		if b, err = ioutil.ReadFile(file); err != nil {
			return nil, err
		}
		source = append(source, string(b))
	}
	return source, nil
}
