// Package codegen provides the code generation functionality for converting VSPEC signals to Go structs and ClickHouse tables.
package codegen

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/tools/imports"
)

const (
	ClickhouseFileName  = "vss-table.sql"
	StructFileName      = "vss-structs.go"
	ConvertFileName     = "vss-convert.go"
	ConvertFuncFileName = "vss-convert-funcs.go"
	readAll             = 0755
)

// TemplateData contains the data to be used during template execution.
type TemplateData struct {
	PackageName string
	DataSignals []*SignalInfo
}

// GetMigratedSignals reads the signals and migrations files and merges them.
func GetMigratedSignals(specFile, migrationFile string) ([]*SignalInfo, error) {
	signals, err := loadSignalsCSV(specFile)
	if err != nil {
		return nil, fmt.Errorf("error reading signals: %w", err)
	}

	migrations, err := loadMigrationJSON(migrationFile)
	if err != nil {
		return nil, fmt.Errorf("error reading migration file: %w", err)
	}

	migratedSignals := migrations.MigratedSignal(signals)
	return migratedSignals, nil
}

// EnsureDir ensures the output directory exists.
func EnsureDir(dir string) error {
	info, err := os.Stat(dir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("output directory is not a directory")
		}
		return nil
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("error checking output directory: %w", err)
	}
	// create the output directory
	err = os.MkdirAll(dir, readAll)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}
	return nil
}

// FormatAndWriteToFile formats the go source with goimports and writes it to the output file.
func FormatAndWriteToFile(goData []byte, outputPath string) error {
	formatted, err := imports.Process(outputPath, goData, &imports.Options{
		AllErrors: true,
		Comments:  true,
	})
	if err != nil {
		// print the error and continue
		fmt.Printf("error formatting go source: %v\n", err)
		formatted = goData
	}
	goOutputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer func() {
		if cerr := goOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()
	_, err = goOutputFile.Write(formatted)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}
