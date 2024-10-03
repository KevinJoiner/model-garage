package convert_test

import (
	"fmt"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/nativestatus/convert"
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
			result, err := convert.ToPowertrainFuelSystemSupportedFuelTypes0(nil, test.input)
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
			result, err := convert.ToPowertrainType0(nil, test.input)
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
			result, err := convert.ToPowertrainCombustionEngineEngineOilLevel0(nil, test.input)
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
			result, err := convert.ToPowertrainTractionBatteryCurrentPower0(nil, test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

func TestToCurrentLocationIsRedacted0(t *testing.T) {
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
			result, err := convert.ToCurrentLocationIsRedacted0(nil, test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}

// Powertrain.Transmission.TravelledDistance
func TestPowertrainTransmissionTravelledDistance(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         float64
		expected      float64
		expectedError bool
	}{
		{
			name:          "Positive Value",
			input:         100.0,
			expected:      100.0,
			expectedError: false,
		},
		{
			name:          "Zero Value",
			input:         0.0,
			expected:      0.0,
			expectedError: false,
		},
		{
			name:          "Negative Value",
			input:         -100.0,
			expected:      -100.0,
			expectedError: false,
		},
		{
			name:          "distance in meters",
			input:         1000.0 * 1000,
			expected:      1000.0,
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
			result, err := convert.ToPowertrainTransmissionTravelledDistance0(nil, test.input)
			if test.expectedError {
				require.Error(t, err, "Expected an error but got none")
			} else {
				require.NoError(t, err, "Unexpected error")
				require.Equal(t, test.expected, result, "Unexpected result")
			}
		})
	}
}
