package schema

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// LoadSignalsCSV loads the signals from a vss CSV file.
func LoadSignalsCSV(r io.Reader) ([]*SignalInfo, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read vspec: %w", err)
	}

	var signals []*SignalInfo
	for i := 1; i < len(records); i++ {
		record := records[i]
		signals = append(signals, NewSignalInfo(record))
	}

	// Sort the signals by name
	slices.SortStableFunc(signals, func(a, b *SignalInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	return signals, nil
}

// LoadDefinitionFile loads the definitions from a definitions.yaml file.
func LoadDefinitionFile(r io.Reader) (*Definitions, error) {
	decoder := yaml.NewDecoder(r)
	var defInfos []*DefinitionInfo
	err := decoder.Decode(&defInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	definitions := &Definitions{
		FromName: map[string]*DefinitionInfo{},
	}
	for _, info := range defInfos {
		definitions.FromName[info.VspecName] = info
	}

	return definitions, nil
}

// TemplateData contains the data to be used during template execution.
type TemplateData struct {
	PackageName string
	ModelName   string
	Signals     []*SignalInfo
}

// GetDefinedSignals reads the signals and definitions files and merges them.
func GetDefinedSignals(specReader, definitionReader io.Reader) (*TemplateData, error) {
	signals, err := LoadSignalsCSV(specReader)
	if err != nil {
		return nil, fmt.Errorf("error reading signals: %w", err)
	}

	definitions, err := LoadDefinitionFile(definitionReader)
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
