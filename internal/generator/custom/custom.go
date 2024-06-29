// Package custom provides the custom generation functionality that uses a provide template to generate a file.
package custom

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

var customFileFormat = "%s.txt"

// Config is the configuration for the Custom generator.
type Config struct {
	// OutputFile is the name of the model to generate the custom file.
	OutputFile string
	// TemplateFile is the path to the template file.
	TemplateFile string
	// Format controls whether the generated file is formatted with goimports.
	Format bool
}

// Generate creates a new Custom file.
func Generate(tmplData *schema.TemplateData, outputDir string, cfg Config) error {
	outFile := cfg.OutputFile
	if outFile == "" {
		lowerName := strings.ToLower(tmplData.ModelName)
		outFile = fmt.Sprintf(customFileFormat, lowerName)
	}

	// create a new Custom file template.
	customFileTmpl, err := createCustomFileTemplate(outFile, cfg.TemplateFile)
	if err != nil {
		return err
	}

	// execute the Custom file template directly to a file.
	filePath := filepath.Clean((filepath.Join(outputDir, outFile)))
	customFileOutputFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating Custom output file: %w", err)
	}
	defer func() {
		if cerr := customFileOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	var outBuf bytes.Buffer
	err = customFileTmpl.Execute(&outBuf, &tmplData)
	if err != nil {
		return fmt.Errorf("error executing Custom template: %w", err)
	}
	if cfg.Format {
		err = codegen.FormatAndWriteToFile(outBuf.Bytes(), filePath)
		if err != nil {
			return fmt.Errorf("error writing file: %w", err)
		}
		return nil
	}
	_, _ = customFileOutputFile.Write(outBuf.Bytes())

	return nil
}

func createCustomFileTemplate(gqlmodelName, templateFile string) (*template.Template, error) {
	tmplName := path.Base(templateFile)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"GQLModelName":       func() string { return gqlmodelName },
		"GQLGenResolverName": templates.ToGo,
	}).ParseFiles(templateFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing Custom file template: %w", err)
	}
	return tmpl, nil
}
