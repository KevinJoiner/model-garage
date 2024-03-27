package codegen

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

func loadSignalsCSV(r io.Reader) ([]*SignalInfo, error) {
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

func loadDefinitionFile(r io.Reader) (*Definitions, error) {
	decoder := yaml.NewDecoder(r)
	var transInfos []*DefinitionInfo
	err := decoder.Decode(&transInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	definitions := &Definitions{
		FromName: map[string]*DefinitionInfo{},
	}
	for _, info := range transInfos {
		definitions.FromName[info.VspecName] = info
	}

	return definitions, nil
}
