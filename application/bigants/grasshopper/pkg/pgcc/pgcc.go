package pgcc

import (
	"bytes"
	"io"
	"log"
	"text/template"
)

// Options defines required and optional settings for building connection query
type Options struct {
	// Column used for cursor
	Cursor string
	// Table to paginate.
	// Should be valid table reference(https://www.postgresql.org/docs/current/queries-table-expressions.html)
	From string
	// Options for sorting the table
	SortKeys []SortKey
}

// SortKey defines sort order as select and order
type SortKey struct {
	Expr  string
	Order string
}

// Render renders SQL query to w with op
func Render(w io.Writer, op Options) error {
	return tmpl.Execute(w, op)
}

// RenderString renders SQL to string with op
func RenderString(op Options) string {
	var buff bytes.Buffer
	err := tmpl.Execute(&buff, op)
	if err != nil {
		log.Fatal(err)
	}
	return string(buff.Bytes())
}

var tmpl = template.Must(template.New("ConnectionTemplate").Parse(`
WITH __params__ AS (
	SELECT $1::int, $3::int
), __after__ AS (
	SELECT {{template "initialSelection" .}} FROM {{.From}} WHERE {{.Cursor}} = $2 LIMIT 1
), __before__ AS (
	SELECT {{template "initialSelection" .}} FROM {{.From}} WHERE {{.Cursor}} = $4 LIMIT 1
), __forward__ AS (
	SELECT {{template "initialSelection" .}} FROM {{.From}}
	WHERE NOT ($1 IS NULL AND $3 IS NOT NULL)
		AND CASE WHEN (SELECT TRUE FROM __after__) THEN ({{template "afterPredicate" .}}) ELSE TRUE END
		AND CASE WHEN (SELECT TRUE FROM __before__) THEN ({{template "beforePredicate" .}}) ELSE TRUE END
	ORDER BY {{ range $i, $key := .SortKeys}}{{- if $i}}, {{end}}{{$key.Expr}} {{if eq $key.Order "ASC"}}ASC{{else}}DESC{{end}}{{end}}
	LIMIT $1 + 1
), __backward__ AS (
	SELECT * FROM (
		SELECT {{template "initialSelection" .}} FROM {{.From}}
		WHERE CASE WHEN (SELECT TRUE FROM __after__) THEN ({{template "afterPredicate" .}}) ELSE TRUE END
			AND CASE WHEN (SELECT TRUE FROM __before__) THEN ({{template "beforePredicate" .}}) ELSE TRUE END
		ORDER BY {{range $i, $key := .SortKeys}}{{- if $i}}, {{end}}{{$key.Expr}} {{if eq $key.Order "ASC"}}DESC{{else}}ASC{{end}}{{end}}
		LIMIT $3 + 1
	) __backward_reversed__
	WHERE $1 IS NULL AND $3 IS NOT NULL
	ORDER BY {{range $i, $key := .SortKeys}}{{if $i}}, {{end}}_sk{{$i}}_ {{$key.Order -}}{{- end}}
), __edges__ AS (
	SELECT * FROM __forward__ UNION SELECT * FROM __backward__
	OFFSET CASE WHEN ($1 > 0 AND $3 > 0) THEN GREATEST(0 - $3 + (SELECT count(*) FROM __forward__), 0) WHEN ($3 > 0) THEN 1 ELSE 0 END
), __prior_to_before__ AS (
	SELECT EXISTS(SELECT {{.Cursor}} FROM {{.From}} WHERE ({{template "priorToBeforePredicate" .}}) LIMIT 1) more FROM __before__
), __prior_to_after__ AS (
	SELECT EXISTS(SELECT {{.Cursor}} FROM {{.From}} WHERE ({{template "priorToAfterPredicate" .}}) LIMIT 1) more FROM __after__
), __page_info__ AS (
	SELECT
		COALESCE((SELECT count(*) > $3 FROM __edges__), (SELECT more FROM __prior_to_after__), FALSE) AS has_previous_page,
		COALESCE((SELECT count(*) > $1 FROM __edges__), (SELECT more FROM __prior_to_before__), FALSE) AS has_next_page
)
SELECT __page_info__.*, __cursor__ FROM __page_info__, __edges__
ORDER BY {{range $i, $key := .SortKeys}}{{if $i}}, {{end}}_sk{{$i}}_ {{$key.Order -}}{{- end}}
LIMIT COALESCE($1, $3)

{{- define "afterPredicate"}}
	{{- range $i, $key := .SortKeys }}
		{{- if $i}} OR ({{end -}}
		{{- range $j, $key := $.SortKeys -}}
		  {{- if ge $i $j -}}
				{{- if $j}} AND {{end -}}
				{{- if gt $i $j}}
					{{- $key.Expr}} = (select _sk{{$j}}_ FROM __after__)
				{{- else}}
					{{- $key.Expr}} {{if eq $key.Order "ASC"}}>{{else}}<{{end}} (select _sk{{$i}}_ FROM __after__)
				{{- end}}
			{{- end}}
		{{- end}}{{if $i}}){{end}}
	{{- end}}
{{- end}}

{{- define "beforePredicate"}}
	{{- range $i, $key := .SortKeys }}
		{{- if $i}} OR ({{end -}}
		{{- range $j, $key := $.SortKeys -}}
		  {{- if ge $i $j -}}
				{{- if $j}} AND {{end -}}
				{{- if gt $i $j}}
					{{- $key.Expr}} = (select _sk{{$j}}_ FROM __before__)
				{{- else}}
					{{- $key.Expr}} {{if eq $key.Order "ASC"}}<{{else}}>{{end}} (select _sk{{$i}}_ FROM __before__)
				{{- end}}
			{{- end}}
		{{- end}}{{if $i}}){{end}}
	{{- end}}
{{- end}}

{{- define "priorToAfterPredicate"}}
	{{- range $i, $key := .SortKeys }}
		{{- if $i}} OR ({{end -}}
		{{- range $j, $key := $.SortKeys -}}
		  {{- if ge $i $j -}}
				{{- if $j}} AND {{end -}}
				{{- if gt $i $j}}
					{{- $key.Expr}} = (select _sk{{$j}}_ FROM __after__)
				{{- else}}
					{{- $key.Expr}} {{if eq $key.Order "ASC"}}<{{else}}>{{end}} (select _sk{{$i}}_ FROM __after__)
				{{- end}}
			{{- end}}
		{{- end}}{{if $i}}){{end}}
	{{- end}}
{{- end}}

{{- define "priorToBeforePredicate"}}
	{{- range $i, $key := .SortKeys }}
		{{- if $i}} OR ({{end -}}
		{{- range $j, $key := $.SortKeys -}}
		  {{- if ge $i $j -}}
				{{- if $j}} AND {{end -}}
				{{- if gt $i $j}}
					{{- $key.Expr}} = (select _sk{{$j}}_ FROM __before__)
				{{- else}}
					{{- $key.Expr}} {{if eq $key.Order "ASC"}}>{{else}}<{{end}} (select _sk{{$i}}_ FROM __before__)
				{{- end}}
			{{- end}}
		{{- end}}{{if $i}}){{end}}
	{{- end}}
{{- end}}

{{- define "initialSelection" -}}
{{.Cursor}} AS __cursor__{{range $i, $key := .SortKeys}}, {{$key.Expr}} AS _sk{{$i}}_{{end}}
{{- end -}}
`))
