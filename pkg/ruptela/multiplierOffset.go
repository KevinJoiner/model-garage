package ruptela

import (
	"fmt"
	"strconv"
)

// Convert102 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '65535'.
func Convert102(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert103 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert103(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(0.39215686274509803)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert107 converts the given raw value to a float64.
// Min: '0' Max: '65535'.
func Convert107(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert114 converts the given raw value to a float64.
// Unit: 'm' Min: '0' Max: '4211081215'.
func Convert114(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(5)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert207 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '250'.
func Convert207(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(0.4)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert30 converts the given raw value to a float64.
// Unit: 'mV' Min: '0' Max: '65535'.
func Convert30(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert483 converts the given raw value to a float64.
// Unit: '-' Min: '0' Max: '250'.
func Convert483(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert642 converts the given raw value to a float64.
// Unit: 'l' Min: '0' Max: '0xFFFF or 65535'.
func Convert642(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert645 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '0xFFFFFFFF'.
func Convert645(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert722 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert722(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert723 converts the given raw value to a float64.
// Unit: 'km' Min: '0' Max: '65535'.
func Convert723(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert95 converts the given raw value to a float64.
// Unit: 'km/h' Min: '0' Max: '255'.
func Convert95(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert96 converts the given raw value to a float64.
// Unit: '°C' Min: '0' Max: '255'.
func Convert96(rawValue string) (float64, error) {
	offset := float64(-40)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert97 converts the given raw value to a float64.
// Unit: '°C' Min: '0' Max: '255'.
func Convert97(rawValue string) (float64, error) {
	offset := float64(-40)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert98 converts the given raw value to a float64.
// Unit: '%' Min: '0' Max: '255'.
func Convert98(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(0.39215686274509803)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}

// Convert99 converts the given raw value to a float64.
// Unit: '-'.
func Convert99(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(1)
	rawInt, err := strconv.ParseUint(rawValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse uint: %w", err)
	}
	return float64(rawInt)*multiplier + offset, nil
}
