package codegen

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func loadSignalsCSV(loadFilePath string) ([]*SignalInfo, error) {
	file, err := os.Open(filepath.Clean(loadFilePath))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	//nolint:errcheck // we don't care about the error since we are not writing to the file
	defer file.Close()

	reader := csv.NewReader(file)
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

func loadDefinitionJSON(loadFilePath string) (*Definitions, error) {
	file, err := os.Open(filepath.Clean(loadFilePath))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	//nolint:errcheck // we don't care about the error since we are not writing to the file
	defer file.Close()

	decoder := json.NewDecoder(file)
	var transInfos []*DefinitionInfo
	err = decoder.Decode(&transInfos)
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
