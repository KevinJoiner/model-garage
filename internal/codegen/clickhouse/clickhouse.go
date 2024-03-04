package clickhouse

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/KevinJoiner/model-garage/internal/codegen"
)

//go:embed vssTable.tmpl
var clickhouseTableTemplate string

// Generate creates a new ClickHouse table file.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	// create a new ClickHouse table template.
	clickhouseTableTmpl, err := createClickHouseTableTemplate()
	if err != nil {
		return err
	}

	// execute the ClickHouse table template directly to a file.
	clickhouseTableOutputFile, err := os.Create(filepath.Join(outputDir, codegen.ClickhouseFileName))
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
	return nil
}

func createClickHouseTableTemplate() (*template.Template, error) {
	tmpl, err := template.New("clickhouseTableTemplate").Funcs(template.FuncMap{
		"escapeDesc": func(desc string) string { return strings.ReplaceAll(desc, `'`, `\'`) },
	}).Parse(clickhouseTableTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	return tmpl, nil
}
