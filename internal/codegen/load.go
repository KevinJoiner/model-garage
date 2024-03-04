package codegen

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func loadSignalsCSV(filePath string) ([]*SignalInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
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

func loadMigrationJSON(filePath string) (*Migrations, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var transInfos []*MigrationInfo
	err = decoder.Decode(&transInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	migrations := &Migrations{
		FromName: map[string]*MigrationInfo{},
	}
	for _, info := range transInfos {
		migrations.FromName[info.VspecName] = info
	}

	return migrations, nil
}
