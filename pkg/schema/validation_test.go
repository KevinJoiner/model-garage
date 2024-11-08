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
				RequiredPrivileges: []string{"VEHICLE_NON_LOCATION_DATA"},
			},
			expected: nil,
		},
		{
			name: "Nil Definition",
			d:    nil,
			expected: InvalidError{
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
			expected: InvalidError{
				Property: "vspecName",
				Name:     "",
				Reason:   "is empty",
			},
		},
		{
			name: "No RequiredPrivileges",
			d: &DefinitionInfo{
				VspecName: "Vehicle",
			},
			expected: InvalidError{
				Property: "requiredPrivileges",
				Name:     "Vehicle",
				Reason:   "at least one privilege is required",
			},
		},
		{
			name: "Invalid RequiredPrivilege",
			d: &DefinitionInfo{
				VspecName:          "Vehicle",
				RequiredPrivileges: []string{"INVALID_PRIVILEGE"},
			},
			expected: InvalidError{
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
