// Package clickhouse provides the ClickHouse table generation functionality for converting VSPEC signals to Go structs and ClickHouse tables.
package clickhouse

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/KevinJoiner/model-garage/internal/codegen"
)

// clickhouseFileName is the name of the ClickHouse table file that will be generated.
var clickhouseFileName = "%s-table.sql"

var embeddedClickhouseFileName = "%s-table.go"

//go:embed table.tmpl
var clickhouseTableTemplate string

//go:embed embeddedTable.tmpl
var embeddedClickhouseTableTemplate string

// Generate creates a new ClickHouse table file.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	setFileNamesFrom(tmplData.ModelName)

	// create a new ClickHouse table template.
	clickhouseTableTmpl, err := createClickHouseTableTemplate()
	if err != nil {
		return err
	}

	// execute the ClickHouse table template directly to a file.
	tablePath := filepath.Clean((filepath.Join(outputDir, clickhouseFileName)))
	clickhouseTableOutputFile, err := os.Create(tablePath)
	if err != nil {
		return fmt.Errorf("error creating ClickHouse table output file: %w", err)
	}
	defer func() {
		if cerr := clickhouseTableOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	err = clickhouseTableTmpl.Execute(clickhouseTableOutputFile, &tmplData)
	if err != nil {
		return fmt.Errorf("error executing ClickHouse table template: %w", err)
	}

	// create a new embedded ClickHouse table template.
	embeddedClickhouseTableTmpl, err := createEmbeddedClickHouseTableTemplate()
	if err != nil {
		return err
	}
	var outBuf bytes.Buffer
	if err = embeddedClickhouseTableTmpl.Execute(&outBuf, &tmplData); err != nil {
		return fmt.Errorf("error executing embedded ClickHouse table template: %w", err)
	}
	filePath := filepath.Clean(filepath.Join(outputDir, embeddedClickhouseFileName))
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), filePath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func setFileNamesFrom(modelName string) {
	lowerName := strings.ToLower(modelName)
	clickhouseFileName = fmt.Sprintf(clickhouseFileName, lowerName)
	embeddedClickhouseFileName = fmt.Sprintf(embeddedClickhouseFileName, lowerName)
}

func createClickHouseTableTemplate() (*template.Template, error) {
	tmpl, err := template.New("clickhouseTableTemplate").Funcs(template.FuncMap{
		"escapeDesc": func(desc string) string { return strings.ReplaceAll(desc, `'`, `\'`) },
		"lower":      strings.ToLower,
	}).Parse(clickhouseTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	return tmpl, nil
}

func createEmbeddedClickHouseTableTemplate() (*template.Template, error) {
	tmpl, err := template.New("embeddedClickhouseTableTemplate").Funcs(template.FuncMap{
		"sqlFileName": func() string { return clickhouseFileName },
	}).Parse(embeddedClickhouseTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	return tmpl, nil
}
