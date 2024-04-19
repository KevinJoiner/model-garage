// Package graphql provides the Graphql table generation functionality for converting VSPEC signals to Go structs and Graphql tables.
package graphql

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

var graphqlFileFormat = "%s.graphqls"

// Config is the configuration for the Graphql generator.
type Config struct {
	// OutputFile is the name of the model to generate the graphql table.
	OutputFile string

	// TemplateFile is the path to the template file.
	TemplateFile string
}

// Generate creates a new Graphql table file.
func Generate(tmplData *codegen.TemplateData, outputDir string, cfg Config) error {
	outFile := cfg.OutputFile
	if outFile == "" {
		lowerName := strings.ToLower(tmplData.ModelName)
		outFile = fmt.Sprintf(graphqlFileFormat, lowerName)
	}

	// create a new Graphql table template.
	graphqlTableTmpl, err := createGraphqlTableTemplate(outFile, cfg.TemplateFile)
	if err != nil {
		return err
	}

	// execute the Graphql table template directly to a file.
	tablePath := filepath.Clean((filepath.Join(outputDir, outFile)))
	graphqlTableOutputFile, err := os.Create(tablePath)
	if err != nil {
		return fmt.Errorf("error creating Graphql table output file: %w", err)
	}
	defer func() {
		if cerr := graphqlTableOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	err = graphqlTableTmpl.Execute(graphqlTableOutputFile, &tmplData)
	if err != nil {
		return fmt.Errorf("error executing Graphql table template: %w", err)
	}

	return nil
}

func createGraphqlTableTemplate(gqlmodelName, templateFile string) (*template.Template, error) {
	tmplName := path.Base(templateFile)
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"GQLModelName": func() string { return gqlmodelName },
	}).ParseFiles(templateFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing Graphql table template: %w", err)
	}
	return tmpl, nil
}
