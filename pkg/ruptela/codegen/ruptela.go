package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Record defines a struct to hold the data for each oid.
type Record struct {
	OID                   string
	Name                  string
	Size                  int
	Type                  string
	MinValue              string
	MaxValue              string
	MultiplierOffset      string
	Units                 string
	ErrorValues           string
	IOExplanationAndNotes string
	Averaging             string
	EventOnChange         string
	EventOnHysteresis     string
	Eco5                  string
	Plug5                 string
	HCV5LitePro5Lite      string
	Pro5                  string
	HCV5                  string
	LCV5                  string
	Trace5LTM             string
	Trace5LTE             string
	Basic                 string
	ECO4UBI               string
	Eco4                  string
	Eco4S                 string
	Eco4T                 string
	Eco4RST               string
	Pro4                  string
	Tco4HCV               string
	Tco4LCV               string
	Plug4                 string
	Pro4BT                string
	Tco4HCVBT             string
	Tco4LCVBT             string
	Eco3                  string
	Pro3                  string
	Tco3TCO               string
	Tco3OBD               string

	// Interpreted values
	Offset     float64
	Multiplier float64
	ErrorRange *ErrorRange
	ErrorSet   []uint64
	MinBig     *big.Int
	MaxBig     *big.Int
}

type ErrorRange struct {
	Min uint64
	Max uint64
}

// Function to load CSV and parse it into a map.
func loadCSVToMap(oidCSV string) (map[string]Record, error) {
	reader := csv.NewReader(strings.NewReader(oidCSV))
	reader.TrimLeadingSpace = true

	// Read the header row to get the column names.
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("could not read headers: %w", err)
	}

	// Create a map to hold the records.
	result := make(map[string]Record)

	// Read through each row.
	for {
		row, err := reader.Read()
		if err != nil {
			break // Assume end of file or error, exit loop.
		}

		// Construct a Record struct from the row.
		size, _ := strconv.Atoi(strings.ReplaceAll(row[2], ",", "")) // Parse "Size, B" as int
		record := Record{
			OID:                   row[0],
			Name:                  row[1],
			Size:                  size,
			Type:                  row[3],
			MinValue:              row[4],
			MaxValue:              row[5],
			MultiplierOffset:      row[6],
			Units:                 row[7],
			ErrorValues:           row[8],
			IOExplanationAndNotes: row[9],
			Averaging:             row[10],
			EventOnChange:         row[11],
			EventOnHysteresis:     row[12],
			Eco5:                  row[13],
			Plug5:                 row[14],
			HCV5LitePro5Lite:      row[15],
			Pro5:                  row[16],
			HCV5:                  row[17],
			LCV5:                  row[18],
			Trace5LTM:             row[19],
			Trace5LTE:             row[20],
			Basic:                 row[21],
			ECO4UBI:               row[22],
			Eco4:                  row[23],
			Eco4S:                 row[24],
			Eco4T:                 row[25],
			Eco4RST:               row[26],
			Pro4:                  row[27],
			Tco4HCV:               row[28],
			Tco4LCV:               row[29],
			Plug4:                 row[30],
			Pro4BT:                row[31],
			Tco4HCVBT:             row[32],
			Tco4LCVBT:             row[33],
			Eco3:                  row[34],
			Pro3:                  row[35],
			Tco3TCO:               row[36],
			Tco3OBD:               row[37],

			Multiplier: 1.0,
			ErrorRange: nil,
		}

		// Use the first column as the oid (assuming the first column is the oid).
		oid := row[0]

		result[oid] = record
	}

	return result, nil
}

func getMultiplierAndOffset(multiplierOffset string) (offset float64, multiplier float64, err error) {
	if strings.Contains(multiplierOffset, "or") {
		return 0, 1, nil
	}
	parts := strings.Split(multiplierOffset, ";")
	if len(parts) == 2 {
		offsetStr := strings.ReplaceAll(parts[1], "offset", "")
		offsetStr = strings.TrimSpace(offsetStr)
		offset, err = strconv.ParseFloat(offsetStr, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse offset: %w", err)
		}
		multiplierStr := strings.TrimSpace(parts[0])
		multiplier, err = getMultiplier(multiplierStr)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse multiplier: %w", err)
		}
	} else if len(parts) == 1 {
		if strings.Contains(parts[0], "offset") {
			offsetStr := strings.ReplaceAll(parts[0], "offset", "")
			offsetStr = strings.TrimSpace(offsetStr)
			offset, err = strconv.ParseFloat(offsetStr, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("could not parse offset: %w", err)
			}
		} else {
			multipartStr := strings.TrimSpace(parts[0])
			multiplier, err = getMultiplier(multipartStr)
			if err != nil {
				return 0, 0, fmt.Errorf("could not parse multiplier: %w", err)
			}
		}
	}
	return offset, multiplier, nil
}

func getMultiplier(multiplier string) (float64, error) {
	if multiplier == "-" {
		return 1, nil
	}
	multiplier = strings.ReplaceAll(multiplier, ",", ".")
	if !strings.Contains(multiplier, "/") {
		multiplier, err := strconv.ParseFloat(multiplier, 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse multiplier: %w", err)
		}
		return multiplier, nil
	}
	// handle fraction multiplier
	fraction := strings.Split(multiplier, "/")
	numerator, err := strconv.ParseFloat(fraction[0], 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse numerator: %w", err)
	}
	denominator, err := strconv.ParseFloat(fraction[1], 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse denominator: %w", err)
	}
	return numerator / denominator, nil
}

func getErrorRange(rawErrorValues string) (*ErrorRange, []uint64, error) {
	rawParts := strings.Split(strings.Split(rawErrorValues, "-")[0], ":")
	errorValues := strings.TrimSpace(rawParts[0])
	if errorValues == "" {
		return nil, nil, nil
	}
	if strings.Contains(errorValues, ",") {
		parts := strings.Split(errorValues, ",")
		result := make([]uint64, len(parts))
		for i, part := range parts {
			value, err := strconv.ParseUint(strings.TrimSpace(part), 0, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("could not parse error value: %w", err)
			}
			result[i] = value
		}
		return nil, result, nil
	}
	if strings.Contains(errorValues, "to") {
		parts := strings.Split(errorValues, "to")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid error values: %s", errorValues)
		}
		start, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 0, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse error value: %w", err)
		}
		end, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 0, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse error value: %w", err)
		}
		errRange := &ErrorRange{
			Min: start,
			Max: end,
		}
		return errRange, nil, nil
	}
	value, err := strconv.ParseUint(errorValues, 0, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse error value: %w", err)
	}
	return nil, []uint64{value}, nil
}

func getMinOrMax(rawVal string) (*big.Int, error) {
	// 255 or 0xFF
	val := strings.TrimSpace(strings.Split(rawVal, "or")[0])
	if val == "" {
		return nil, nil
	}
	// 65,535
	val = strings.ReplaceAll(val, ",", "")
	bigInt, ok := new(big.Int).SetString(val, 0)
	if !ok {
		return nil, fmt.Errorf("could not parse value: %s", val)
	}
	return bigInt, nil
}
