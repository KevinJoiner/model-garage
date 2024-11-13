package schema

import (
	"fmt"
	"slices"
)

// privileges are defined on chain and copied here for validation.
var privileges = []string{"VEHICLE_NON_LOCATION_DATA", "VEHICLE_COMMANDS", "VEHICLE_CURRENT_LOCATION", "VEHICLE_ALL_TIME_LOCATION", "VEHICLE_VIN_CREDENTIAL"}

// InvalidError is an error for invalid definitions.
type InvalidError struct {
	Property string
	Name     string
	Reason   string
}

func (e InvalidError) Error() string {
	return fmt.Sprintf("signal '%s' property '%s' %s", e.Name, e.Property, e.Reason)
}

// Validate checks if the definition is valid.
func Validate(d *DefinitionInfo) error {
	if d == nil {
		return InvalidError{Property: "", Name: "", Reason: "is nil"}
	}
	if d.VspecName == "" {
		return InvalidError{Property: "vspecName", Name: d.VspecName, Reason: "is empty"}
	}
	for _, conv := range d.Conversions {
		if conv == nil {
			return InvalidError{Property: "conversion", Name: d.VspecName, Reason: "is nil"}
		}
		if conv.OriginalName == "" {
			return InvalidError{Property: "originalName", Name: d.VspecName, Reason: "is empty"}
		}
	}
	for _, priv := range d.RequiredPrivileges {
		if !slices.Contains(privileges, priv) {
			return InvalidError{Property: "requiredPrivileges", Name: d.VspecName, Reason: fmt.Sprintf("must be one of %v", privileges)}
		}
	}
	return nil
}
