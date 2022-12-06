package main

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	types := []string{
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"float32",
		"float64",
		"byte",
		"rune",
	}

	type dataType struct {
		Name string
		Type string
	}

	dataTypes := []dataType{}

	for _, t := range types {
		dataTypes = append(dataTypes, dataType{
			Name: cases.Title(language.Und, cases.NoLower).String(t),
			Type: t,
		})
	}

	tmpl, err := template.ParseGlob(path.Join("generator", "*.tpl"))
	if err != nil {
		panic(err)
	}

	for _, fileName := range []string{
		"validator_number.gen.go",
		"validator_number.gen_test.go",
		"validator_number_p.gen.go",
		"validator_number_p.gen_test.go",
	} {
		templateName := fmt.Sprintf("%s.tpl", fileName)

		output, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(output, templateName, dataTypes)
		if err != nil {
			panic(err)
		}
	}
}
