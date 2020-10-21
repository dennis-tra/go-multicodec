package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

// Codec represents a row in the canonical multicodec table
type Codec struct {
	Name        string
	Tag         string
	Code        string
	Description string
}

// IsDeprecated returns true if the description contains the word deprecated
func (c Codec) IsDeprecated() bool {
	return strings.Contains(strings.ToLower(c.Description), "deprecated")
}

// VarName returns the variable name to be used in Go for this codec
// Implementation adapted from https://github.com/iancoleman/strcase
func (c Codec) VarName() string {
	if c.Name == "" {
		return c.Name
	}

	// convert the string to camelcase
	n := strings.Builder{}
	n.Grow(len(c.Name))
	capNext := true
	for i, v := range []byte(c.Name) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		vIsNum := v >= '0' && v <= '9'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
			if capNext {
				n.WriteByte('_')
			}
		}
	}
	return n.String()
}

func main() {

	resp, err := http.Get("https://raw.githubusercontent.com/multiformats/multicodec/master/table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	records, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	codecs := make([]Codec, len(records)-1)
	for i, record := range records[1:] { // Skip header line
		codecs[i] = Codec{
			Name:        strings.TrimSpace(record[0]),
			Tag:         strings.TrimSpace(record[1]),
			Code:        strings.TrimSpace(record[2]),
			Description: strings.TrimSpace(record[3]),
		}
	}

	tplFileName := "const.go.tpl"
	t, err := template.
		New(tplFileName).
		Funcs(template.FuncMap{"ToTitle": strings.Title}).
		ParseFiles(path.Join("gen", tplFileName))
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("const.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = t.Execute(out, codecs)
	if err != nil {
		log.Fatal(err)
	}
}
