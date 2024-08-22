package convert

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// VersionError is an error for unsupported specversion.
type VersionError struct {
	Version string
}

// Error returns the error message.
func (e VersionError) Error() string {
	return fmt.Sprintf("unsupported verision: %s", e.Version)
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

// ConversionError is an error that occurs during conversion.
type ConversionError struct {
	// DecodedSignals is the list of signals that were successfully decoded.
	DecodedSignals []vss.Signal `json:"decodedSignals"`
	Errors         []error      `json:"errors"`
	TokenID        uint32       `json:"tokenId"`
	Source         string       `json:"source"`
}

// Error returns the error message.
func (e ConversionError) Error() string {
	var errBuilder strings.Builder
	errBuilder.WriteString("conversion error")
	if e.TokenID != 0 {
		errBuilder.WriteString(" tokenId '")
		errBuilder.WriteString(strconv.FormatUint(uint64(e.TokenID), 10))
		errBuilder.WriteString("'")
	}

	if e.Source != "" {
		errBuilder.WriteString(" source '")
		errBuilder.WriteString(e.Source)
		errBuilder.WriteString("'")
	}
	if len(e.Errors) != 0 {
		errBuilder.WriteString(fmt.Sprintf(": %v", e.Errors))
	}
	return errBuilder.String()
}

// Unwrap returns all errors that occurred during conversion.
func (e ConversionError) Unwrap() []error {
	return e.Errors
}
