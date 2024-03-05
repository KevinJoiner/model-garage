// Package model provides the functionality to generate a Go struct file to represent the vehicle struct.
package model

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/KevinJoiner/model-garage/internal/codegen"
)

//go:embed vehicle.tmpl
var structTemplate string

// Generate creates a new Go struct file to represent the vehicle struct.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	modelTemplate, err := createModelTemplate()
	if err != nil {
		return err
	}

	// execute the struct template
	var outBuf bytes.Buffer
	if err = modelTemplate.Execute(&outBuf, &tmplData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	goOutputPath := filepath.Join(outputDir, codegen.StructFileName)
	// format and write the go file.
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), goOutputPath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func createModelTemplate() (*template.Template, error) {
	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sqlFileName": func() string { return codegen.ClickhouseFileName },
	}).Parse(structTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}
