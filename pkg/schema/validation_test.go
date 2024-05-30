package schema

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		d        *DefinitionInfo
		expected error
	}{
		{
			name: "Valid Definition",
			d: &DefinitionInfo{
				VspecName:          "Vehicle",
				GoType:             "float64",
				Conversions:        []*ConversionInfo{{OriginalName: "OriginalName"}},
				RequiredPrivileges: []string{"VEHICLE_NON_LOCATION_DATA"},
			},
			expected: nil,
		},
		{
			name: "Nil Definition",
			d:    nil,
			expected: ErrInvalid{
				Property: "",
				Name:     "",
				Reason:   "is nil",
			},
		},
		{
			name: "Empty VspecName",
			d: &DefinitionInfo{
				VspecName: "",
			},
			expected: ErrInvalid{
				Property: "vspecName",
				Name:     "",
				Reason:   "is empty",
			},
		},
		{
			name: "Invalid GoType",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				GoType:    "int",
			},
			expected: ErrInvalid{
				Property: "goType",
				Name:     "int",
				Reason:   "must be one of [float64 string]",
			},
		},
		{
			name: "No Conversions",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				GoType:    "float64",
			},
			expected: ErrInvalid{
				Property: "conversions",
				Name:     "Vehicle",
				Reason:   "at least one conversion is required",
			},
		},
		{
			name: "Nil Conversion",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				GoType:    "float64",
				Conversions: []*ConversionInfo{
					nil,
				},
			},
			expected: ErrInvalid{
				Property: "conversion",
				Name:     "Vehicle",
				Reason:   "is nil",
			},
		},
		{
			name: "Empty OriginalName",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
				GoType:    "float64",
				Conversions: []*ConversionInfo{
					{OriginalName: ""},
				},
			},
			expected: ErrInvalid{
				Property: "originalName",
				Name:     "Vehicle",
				Reason:   "is empty",
			},
		},
		{
			name: "No RequiredPrivileges",
			d: &DefinitionInfo{
				VspecName:   "Vehicle",
				GoType:      "float64",
				Conversions: []*ConversionInfo{{OriginalName: "OriginalName"}},
			},
			expected: ErrInvalid{
				Property: "requiredPrivileges",
				Name:     "Vehicle",
				Reason:   "at least one privilege is required",
			},
		},
		{
			name: "Invalid RequiredPrivilege",
			d: &DefinitionInfo{
				VspecName:          "Vehicle",
				GoType:             "float64",
				Conversions:        []*ConversionInfo{{OriginalName: "OriginalName"}},
				RequiredPrivileges: []string{"INVALID_PRIVILEGE"},
			},
			expected: ErrInvalid{
				Property: "requiredPrivileges",
				Name:     "Vehicle",
				Reason:   "must be one of [VEHICLE_NON_LOCATION_DATA VEHICLE_COMMANDS VEHICLE_CURRENT_LOCATION VEHICLE_ALL_TIME_LOCATION VEHICLE_VIN_CREDENTIAL]",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Validate(test.d)
			if result != test.expected {
				t.Errorf("Unexpected result. Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}
