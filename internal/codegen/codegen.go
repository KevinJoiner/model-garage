// Package codegen provides the code generation functionality for converting VSPEC signals to Go structs and ClickHouse tables.
package codegen

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/imports"
)

const (
	// ClickhouseFileName is the name of the ClickHouse table file that will be generated.
	ClickhouseFileName = "vss-table.sql"

	// StructFileName is the name of the Go file that will contain the vehicle struct.
	StructFileName = "vss-structs.go"

	// ConvertFileName is the name of the Go file that will convert JSON data to the vehicle struct.
	ConvertFileName = "vss-convert.go"

	// ConvertFuncFileName is the name of the Go file that will contain the conversion functions.
	ConvertFuncFileName = "vss-convert-funcs.go"

	readAll = 0o755
)

// TemplateData contains the data to be used during template execution.
type TemplateData struct {
	PackageName string
	Signals     []*SignalInfo
}

// GetDefinedSignals reads the signals and definitions files and merges them.
func GetDefinedSignals(specFile, definitionFile string) ([]*SignalInfo, error) {
	signals, err := loadSignalsCSV(specFile)
	if err != nil {
		return nil, fmt.Errorf("error reading signals: %w", err)
	}

	definitions, err := loadDefinitionJSON(definitionFile)
	if err != nil {
		return nil, fmt.Errorf("error reading definition file: %w", err)
	}

	definedSignals := definitions.DefinedSignal(signals)
	return definedSignals, nil
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
func FormatAndWriteToFile(goData []byte, outputFilePath string) (err error) {
	cleanPath := filepath.Clean(outputFilePath)
	formatted, fmtErr := imports.Process(cleanPath, goData, &imports.Options{
		AllErrors: true,
		Comments:  true,
	})
	if fmtErr != nil {
		// do not return early, we still want to write the file
		fmtErr = fmt.Errorf("error formatting go source: %w", fmtErr)
		formatted = goData
	}
	goOutputFile, err := os.Create(cleanPath)
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

	// return the formatting error if there is one
	return fmtErr
}
