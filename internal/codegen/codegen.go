// Package codegen provides the code generation functionality for converting VSPEC signals to Go structs and ClickHouse tables.
package codegen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
)

const readAll = 0o755

// TemplateData contains the data to be used during template execution.
type TemplateData struct {
	PackageName string
	ModelName   string
	Signals     []*SignalInfo
}

// GetDefinedSignals reads the signals and definitions files and merges them.
func GetDefinedSignals(specReader, definitionReader io.Reader) (*TemplateData, error) {
	signals, err := loadSignalsCSV(specReader)
	if err != nil {
		return nil, fmt.Errorf("error reading signals: %w", err)
	}

	definitions, err := loadDefinitionFile(definitionReader)
	if err != nil {
		return nil, fmt.Errorf("error reading definition file: %w", err)
	}
	signals = definitions.DefinedSignal(signals)
	modelName := "Model"
	if len(signals) > 0 {
		idx := strings.IndexByte(signals[0].Name, '.')
		if idx > 0 {
			modelName = signals[0].Name[:idx]
			modelName = cases.Title(language.English).String(modelName)
		}
	}
	tmplData := &TemplateData{
		Signals:   signals,
		ModelName: modelName,
	}

	return tmplData, nil
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
