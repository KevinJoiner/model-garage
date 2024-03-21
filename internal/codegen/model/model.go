// Package model provides the functionality to generate a Go struct file to represent the vehicle struct.
package model

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

// structFileName is the name of the Go file that will contain the vehicle struct.
var structFileName = "%s-structs.go"

//go:embed vehicle.tmpl
var structTemplate string

// Generate creates a new Go struct file to represent the vehicle struct.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	structFileName = fmt.Sprintf(structFileName, strings.ToLower(tmplData.ModelName))
	modelTemplate, err := createModelTemplate()
	if err != nil {
		return err
	}

	// execute the struct template
	var outBuf bytes.Buffer
	if err = modelTemplate.Execute(&outBuf, &tmplData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	goOutputPath := filepath.Join(outputDir, structFileName)
	// format and write the go file.
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), goOutputPath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func createModelTemplate() (*template.Template, error) {
	tmpl, err := template.New("structTemplate").Parse(structTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}
