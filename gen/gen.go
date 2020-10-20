package main

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

// TemplateData is the data structure that is passed into the go template
type TemplateData struct {
	Codecs []Codec
}

func main() {

	table, err := os.Open(path.Join("multicodec", "table.csv"))
	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()

	records, err := csv.NewReader(table).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	codecs := make([]Codec, len(records)-1)
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
		codecs[i-1] = c
	}

	tData := TemplateData{
		Codecs: codecs,
	}

	templates := []string{"codecs", "const"}
	for _, templateName := range templates {

		tplFileName := templateName + ".go.tpl"
		t, err := template.
			New(tplFileName).
			Funcs(template.FuncMap{"ToTitle": strings.Title}).
			ParseFiles(path.Join("templates", tplFileName))
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(templateName + ".go")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(out, tData)
		if err != nil {
			out.Close()
			log.Fatal(err)
		}
		out.Close()
	}
}
