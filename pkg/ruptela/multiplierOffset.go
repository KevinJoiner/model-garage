package ruptela

import (
	"fmt"
	"strconv"
)

func Convert103(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(0.39215686274509803)
	rawFloat, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float: %v", err)
	}
	return rawFloat*multiplier + offset, nil
}

func Convert114(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(5)
	rawFloat, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float: %v", err)
	}
	return rawFloat*multiplier + offset, nil
}

func Convert96(rawValue string) (float64, error) {
	offset := float64(-40)
	multiplier := float64(1)
	rawFloat, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float: %v", err)
	}
	return rawFloat*multiplier + offset, nil
}

func Convert97(rawValue string) (float64, error) {
	offset := float64(-40)
	multiplier := float64(1)
	rawFloat, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float: %v", err)
	}
	return rawFloat*multiplier + offset, nil
}

func Convert98(rawValue string) (float64, error) {
	offset := float64(0)
	multiplier := float64(0.39215686274509803)
	rawFloat, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse float: %v", err)
	}
	return rawFloat*multiplier + offset, nil
}
