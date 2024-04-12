package migration

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

var migrationFileFormat = "%s_%s_migration.go"

const timestampFormat = "20060102150405"

//go:embed migration.tmpl
var migrationFileTemplate string

// Generate creates a new ClickHouse table file.
func Generate(tmplData *codegen.TemplateData, outputDir string) error {
	version := time.Now().UTC().Format(timestampFormat)
	migrationTempl, err := createMigrationTemplate(version)
	if err != nil {
		return err
	}

	var outBuf bytes.Buffer
	err = migrationTempl.Execute(&outBuf, &tmplData)
	if err != nil {
		return fmt.Errorf("error executing ClickHouse table template: %w", err)
	}
	migrationFilePath := getFilePath(tmplData.ModelName, outputDir, version)
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), migrationFilePath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func createMigrationTemplate(version string) (*template.Template, error) {
	tmpl, err := template.New("migrationTemplate").Funcs(template.FuncMap{
		"escapeDesc": func(desc string) string { return strings.ReplaceAll(desc, `'`, `\'`) },
		"lower":      strings.ToLower,
		"version":    func() string { return version },
	}).Parse(migrationFileTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	return tmpl, nil
}

func getFilePath(modelName, outputDir string, version string) string {
	migrationFileName := fmt.Sprintf(migrationFileFormat, version, modelName)
	return filepath.Clean(filepath.Join(outputDir, migrationFileName))
}
