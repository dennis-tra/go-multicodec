package main

//go:generate go run ./gen/gen.go
//go:generate gofmt -w codecs.go

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

var templates = []string{
	"codecs",
	"const",
	"tag",
}

// TemplateData is the data structure that is passed into the go template
type TemplateData struct {
	Codecs      []Codec
	CodecsByTag map[string][]Codec
	Tags        []string
}

func main() {

	table, err := os.Open("multicodec/table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	reader := csv.NewReader(table)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	tags := []string{}
	codecsByTag := map[string][]Codec{}
	codecs := []Codec{}
	for i, record := range records {

		// Skip header line
		if i == 0 {
			continue
		}

		c := Codec{
			Name:        strings.TrimSpace(record[0]),
			Tag:         strings.TrimSpace(record[1]),
			Code:        strings.TrimSpace(record[2]),
			Description: strings.TrimSpace(record[3]),
		}
		codecs = append(codecs, c)

		if _, exists := codecsByTag[c.Tag]; !exists {
			tags = append(tags, c.Tag)
			codecsByTag[c.Tag] = []Codec{}
		}
		codecsByTag[c.Tag] = append(codecsByTag[c.Tag], c)
	}

	sort.Strings(tags)

	tData := TemplateData{
		Codecs:      codecs,
		CodecsByTag: codecsByTag,
		Tags:        tags,
	}

	funcMap := template.FuncMap{
		"ToTitle": strings.Title,
	}

	for _, templateName := range templates {

		t, err := template.New(templateName + ".go.tpl").Funcs(funcMap).ParseFiles("templates/" + templateName + ".go.tpl")
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(templateName + ".go")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		err = t.Execute(out, tData)
		if err != nil {
			log.Fatal(err)
		}
	}
}
