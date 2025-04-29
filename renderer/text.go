package renderer

import (
	"os"
	"path/filepath"
	"text/template"
)

func Text(tmplFilePath string, data any) {
	tmplName := filepath.Base(tmplFilePath)

	// Create template
	tmpl, err := template.New(tmplName).ParseFiles(tmplFilePath)
	if err != nil {
		panic(err)
	}

	// Fill template with data
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
