// Package graphql provides the Graphql table generation functionality for converting VSPEC signals to Go structs and Graphql tables.
package graphql

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

// graphqlFileName is the name of the Graphql table file that will be generated.
var graphqlFileName = "%s-gql.graphql"

var goGraphqlFileName = "%s-gql.go"

//go:embed gql.tmpl
var graphqlTableTemplate string

//go:embed goTable.tmpl
var goGraphqlTableTemplate string

// Generate creates a new Graphql table file.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	setFileNamesFrom(tmplData.ModelName)

	// create a new Graphql table template.
	graphqlTableTmpl, err := createGraphqlTableTemplate()
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

	// create a new go Graphql table template.
	goGraphqlTableTmpl, err := createGoGraphqlTableTemplate()
	if err != nil {
		return err
	}
	var outBuf bytes.Buffer
	if err = goGraphqlTableTmpl.Execute(&outBuf, &tmplData); err != nil {
		return fmt.Errorf("error executing go Graphql table template: %w", err)
	}
	filePath := filepath.Clean(filepath.Join(outputDir, goGraphqlFileName))
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), filePath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func setFileNamesFrom(modelName string) {
	lowerName := strings.ToLower(modelName)
	graphqlFileName = fmt.Sprintf(graphqlFileName, lowerName)
	goGraphqlFileName = fmt.Sprintf(goGraphqlFileName, lowerName)
}

func createGraphqlTableTemplate() (*template.Template, error) {
	tmpl, err := template.New("graphqlTableTemplate").Funcs(template.FuncMap{
		"escapeDesc": func(desc string) string { return strings.ReplaceAll(desc, `'`, `\'`) },
		"lower":      strings.ToLower,
	}).Parse(graphqlTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing Graphql table template: %w", err)
	}
	return tmpl, nil
}

func createGoGraphqlTableTemplate() (*template.Template, error) {
	tmpl, err := template.New("goGraphqlTableTemplate").Funcs(template.FuncMap{
		"schemaFile": func() string { return graphqlFileName },
	}).Parse(goGraphqlTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing Graphql table template: %w", err)
	}
	return tmpl, nil
}
