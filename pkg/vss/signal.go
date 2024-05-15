// Package vss holds the data structures and functions for working with signals from DIMOs VSS schema.
package vss

import (
	"fmt"
	"time"
)

const (
	// TableName is the name of the distributed table in Clickhouse.
	TableName = "signal"
)

// Signal represents a single signal collected from a device.
// This is the data format that is stored in the database.
type Signal struct {
	// TokenID is the unique identifier of the device.
	TokenID uint32 `ch:"token_id" json:"tokenId"`

	// Timestamp is when this data was collected.
	Timestamp time.Time `ch:"timestamp" json:"timestamp"`

	// Name is the name of the signal collected.
	Name string `ch:"name" json:"name"`

	// ValueNumber is the value of the signal collected.
	ValueNumber float64 `ch:"value_number" json:"valueNumber"`

	// ValueString is the value of the signal collected.
	ValueString string `ch:"value_string" json:"valueString"`

	// Source is the source of the signal collected.
	Source string `ch:"source" json:"source"`
}

// SetValue dynamically set the appropriate value field based on the type of the value.
func (s *Signal) SetValue(val any) {
	switch typedVal := val.(type) {
	case float64:
		s.ValueNumber = typedVal
	case string:
		s.ValueString = typedVal
	default:
		s.ValueString = fmt.Sprintf("%v", val)
	}
}

// SignalToSlice converts a Signal to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `SignalColNames`.
func SignalToSlice(obj Signal) []any {
	return []any{
		obj.TokenID,
		obj.Timestamp,
		obj.Name,
		obj.ValueNumber,
		obj.ValueString,
		obj.Source,
	}
}

// SignalColNames returns the column names of the Signal struct.
func SignalColNames() []string {
	return []string{
		"TokenID",
		"Timestamp",
		"Name",
		"ValueNumber",
		"ValueString",
		"Source",
	}
}
