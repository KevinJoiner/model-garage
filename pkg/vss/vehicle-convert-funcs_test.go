package vss_test

import (
	"fmt"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestToPowertrainFuelSystemSupportedFuelTypes0(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Gasoline",
			input:         "Gasoline",
			expected:      "GASOLINE",
			expectedError: false,
		},
		{
			name:          "Ethanol",
			input:         "Ethanol",
			expected:      "E85",
			expectedError: false,
		},
		{
			name:          "Diesel",
			input:         "Diesel",
			expected:      "DIESEL",
			expectedError: false,
		},
		{
			name:          "LPG",
			input:         "LPG",
			expected:      "LPG",
			expectedError: false,
		},
		{
			name:          "Unknown Fuel Type",
			input:         "UnknownFuelType",
			expected:      "OTHER",
			expectedError: false,
		},
		{
			name:          "Electric",
			input:         "Electric",
			expected:      "OTHER",
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToPowertrainFuelSystemSupportedFuelTypes0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToPowertrainType0(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError bool
	}{
		{
			name:          "Gasoline",
			input:         "Gasoline",
			expected:      "COMBUSTION",
			expectedError: false,
		},
		{
			name:          "Ethanol",
			input:         "Ethanol",
			expected:      "COMBUSTION",
			expectedError: false,
		},
		{
			name:          "Diesel",
			input:         "Diesel",
			expected:      "COMBUSTION",
			expectedError: false,
		},
		{
			name:          "LPG",
			input:         "LPG",
			expected:      "COMBUSTION",
			expectedError: false,
		},
		{
			name:          "Unknown Fuel Type",
			input:         "UnknownFuelType",
			expected:      "COMBUSTION",
			expectedError: false,
		},
		{
			name:          "Electric",
			input:         "Electric",
			expected:      "ELECTRIC",
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToPowertrainType0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToPowertrainCombustionEngineEngineOilLevel0(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         float64
		expected      string
		expectedError bool
	}{
		{
			name:          "Critically Low",
			input:         0.2,
			expected:      "CRITICALLY_LOW",
			expectedError: false,
		},
		{
			name:          "Low",
			input:         0.4,
			expected:      "LOW",
			expectedError: false,
		},
		{
			name:          "Normal",
			input:         0.6,
			expected:      "NORMAL",
			expectedError: false,
		},
		{
			name:          "High",
			input:         0.9,
			expected:      "HIGH",
			expectedError: false,
		},
		{
			name:          "Critically High",
			input:         1.0,
			expected:      "CRITICALLY_HIGH",
			expectedError: false,
		},
		{
			name:          "Above 1.0",
			input:         1.1,
			expected:      "CRITICALLY_HIGH",
			expectedError: false,
		},
		{
			name:          "Negative Value",
			input:         -0.1,
			expected:      "CRITICALLY_LOW",
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToPowertrainCombustionEngineEngineOilLevel0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToCurrentLocationTimestamp1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		expected      float64
		expectedError bool
	}{
		{
			name:          "Valid Value",
			input:         "2022-01-01T12:34:56Z",
			expected:      1641040496,
			expectedError: false,
		},
		{
			name:          "Zero Value",
			input:         "1970-01-01T00:00:00Z",
			expected:      0.0,
			expectedError: false,
		},
		{
			name:          "Negative Value",
			input:         "1969-12-31T23:59:50Z",
			expected:      -10,
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToCurrentLocationTimestamp0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToPowertrainTractionBatteryCurrentPower1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         float64
		expected      float64
		expectedError bool
	}{
		{
			name:          "Positive (charging) value",
			input:         113.0,
			expected:      113000.0,
			expectedError: false,
		},
		{
			name:          "Zero Value",
			input:         0.0,
			expected:      0.0,
			expectedError: false,
		},
		{
			name:          "Negative (driving) value",
			input:         -11.0,
			expected:      -11000.0,
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToPowertrainTractionBatteryCurrentPower0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToDIMOIsLocationRedacted0(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         bool
		expected      float64
		expectedError bool
	}{
		{
			name:          "True to 1",
			input:         true,
			expected:      1,
			expectedError: false,
		},
		{
			name:          "False to 0",
			input:         false,
			expected:      0,
			expectedError: false,
		},
	}

	for i := range tests {
		test := tests[i]
		name := test.name
		if name == "" {
			name = fmt.Sprintf("Input: %v", test.input)
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := vss.ToDIMOIsLocationRedacted0(test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}
