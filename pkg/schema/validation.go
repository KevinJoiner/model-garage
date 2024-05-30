package schema

import (
	"fmt"
	"slices"
)

var (
	goTypes = []string{"float64", "string"}

	// privileges are defined on chain and copied here for validation.
	privileges = []string{"VEHICLE_NON_LOCATION_DATA", "VEHICLE_COMMANDS", "VEHICLE_CURRENT_LOCATION", "VEHICLE_ALL_TIME_LOCATION", "VEHICLE_VIN_CREDENTIAL"}
)

// ErrInvalid is an error for invalid definitions.
type ErrInvalid struct {
	Property string
	Name     string
	Reason   string
}

func (e ErrInvalid) Error() string {
	return fmt.Sprintf("signal '%s' property '%s' %s", e.Property, e.Name, e.Reason)
}

// Validate checks if the definition is valid.
func Validate(d *DefinitionInfo) error {
	if d == nil {
		return ErrInvalid{Property: "", Name: "", Reason: "is nil"}
	}
	if d.VspecName == "" {
		return ErrInvalid{Property: "vspecName", Name: d.VspecName, Reason: "is empty"}
	}
	if d.GoType != "" && !slices.Contains(goTypes, d.GoType) {
		return ErrInvalid{Property: "goType", Name: d.GoType, Reason: fmt.Sprintf("must be one of %v", goTypes)}
	}
	if len(d.Conversions) == 0 {
		return ErrInvalid{Property: "conversions", Name: d.VspecName, Reason: "at least one conversion is required"}
	}
	for _, conv := range d.Conversions {
		if conv == nil {
			return ErrInvalid{Property: "conversion", Name: d.VspecName, Reason: "is nil"}
		}
		if conv.OriginalName == "" {
			return ErrInvalid{Property: "originalName", Name: d.VspecName, Reason: "is empty"}
		}
	}
	if len(d.RequiredPrivileges) == 0 {
		return ErrInvalid{Property: "requiredPrivileges", Name: d.VspecName, Reason: "at least one privilege is required"}
	}
	for _, priv := range d.RequiredPrivileges {
		if !slices.Contains(privileges, priv) {
			return ErrInvalid{Property: "requiredPrivileges", Name: d.VspecName, Reason: fmt.Sprintf("must be one of %v", privileges)}
		}
	}
	return nil
}
