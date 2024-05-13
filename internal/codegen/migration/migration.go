// Package migration provides a function for migrating a clickhouse database to a schema.
package migration

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	delemReplacer = strings.NewReplacer("_", " ", "-", " ", ".", " ")
	titleCaser    = cases.Title(language.AmericanEnglish, cases.NoLower)
	lowerCaser    = cases.Lower(language.AmericanEnglish)
)

const (
	migrationFileFormat = "%s_%s_migration.go"
	timestampFormat     = "20060102150405"
)

//go:embed migration.tmpl
var migrationFileTemplate string

// Config is the configuration for the migration generator.
type Config struct {
	// fileName is the name of the migration file.
	FileName string
}

// Generate creates a new ClickHouse table file.
func Generate(tmplData *codegen.TemplateData, outputDir string, cfg Config) error {
	version, err := getVersion(outputDir)
	if err != nil {
		return fmt.Errorf("error getting version: %w", err)
	}
	fileName := cfg.FileName
	if fileName == "" {
		fileName = tmplData.ModelName
	}
	migrationTempl, err := createMigrationTemplate(fileName)
	if err != nil {
		return err
	}

	var outBuf bytes.Buffer
	err = migrationTempl.Execute(&outBuf, &tmplData)
	if err != nil {
		return fmt.Errorf("error executing ClickHouse table template: %w", err)
	}

	fileName = delemReplacer.Replace(fileName)
	migrationFilePath := getFilePath(fileName, outputDir, version)
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), migrationFilePath)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func createMigrationTemplate(fileName string) (*template.Template, error) {
	funcName := strings.ReplaceAll(titleCaser.String(fileName), " ", "")
	tmpl, err := template.New("migrationTemplate").Funcs(template.FuncMap{
		"funcName": func() string { return funcName },
	}).Parse(migrationFileTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	return tmpl, nil
}

func getFilePath(fileName, outputDir string, version string) string {
	noSpaceName := lowerCaser.String(strings.ReplaceAll(fileName, " ", "_"))
	migrationFileName := fmt.Sprintf(migrationFileFormat, version, noSpaceName)
	return filepath.Clean(filepath.Join(outputDir, migrationFileName))
}

func getVersion(outputDir string) (string, error) {
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return "", fmt.Errorf("error reading directory: %w", err)
	}
	var maxVersion int
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if strings.HasSuffix(name, ".go") {
			versionStr := strings.Split(name, "_")[0]
			version, err := strconv.Atoi(versionStr)
			if err != nil {
				continue
			}
			if version > maxVersion {
				maxVersion = version
			}
		}
	}
	return fmt.Sprintf("%05d", maxVersion+1), nil
}
