package convert

import "fmt"

// VersionError is an error for unsupported specversion.
type VersionError struct {
	Version string
}

// Error returns the error message.
func (e VersionError) Error() string {
	return fmt.Sprintf("unsupported specversion: %s", e.Version)
}

// FieldNotFoundError is an error for missing fields.
type FieldNotFoundError struct {
	Field  string
	Lookup string
}

// Error returns the error message.
func (e FieldNotFoundError) Error() string {
	return fmt.Sprintf("field not found: %s (lookupKey: %s)", e.Field, e.Lookup)
}
