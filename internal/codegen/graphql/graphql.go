// Package graphql provides the Graphql table generation functionality for converting VSPEC signals to Go structs and Graphql tables.
package graphql

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

// graphqlFileName is the name of the Graphql table file that will be generated.
var graphqlFileName = "%s-gql.graphqls"

//go:embed gql.tmpl
var graphqlTableTemplate string

// Generate creates a new Graphql table file.
func Generate(tmplData *codegen.TemplateData, outputDir, gqlModelName string) error {
	if gqlModelName == "" {
		gqlModelName = tmplData.ModelName
	}

	setFileNamesFrom(gqlModelName)

	// create a new Graphql table template.
	graphqlTableTmpl, err := createGraphqlTableTemplate(gqlModelName)
	if err != nil {
		return err
	}

	// execute the Graphql table template directly to a file.
	tablePath := filepath.Clean((filepath.Join(outputDir, graphqlFileName)))
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

func setFileNamesFrom(modelName string) {
	lowerName := strings.ToLower(modelName)
	graphqlFileName = fmt.Sprintf(graphqlFileName, lowerName)
}

func createGraphqlTableTemplate(gqlmodelName string) (*template.Template, error) {
	tmpl, err := template.New("graphqlTableTemplate").Funcs(template.FuncMap{
		"GQLModelName": func() string { return gqlmodelName },
	}).Parse(graphqlTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing Graphql table template: %w", err)
	}
	return tmpl, nil
}
