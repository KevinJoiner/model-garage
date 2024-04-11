package vss

import "time"

// Signal represents a single signal collected from a device.
// This is the data format that is stored in the database.
type Signal struct {
	// TokenID is the unique identifier of the device.
	TokenID uint32 `ch:"TokenId" json:"tokenId"`

	// Timestamp is when this data was collected.
	Timestamp time.Time `ch:"Timestamp" json:"timestamp"`

	// SignalName is the name of the signal collected.
	SignalName string `ch:"SignalName" json:"signalName"`

	// ValueNumber is the value of the signal collected.
	ValueNumber float64 `ch:"ValueNumber" json:"valueNumber"`

	// ValueString is the value of the signal collected.
	ValueString string `ch:"ValueString" json:"valueString"`

	// ValueStringArray is the value of the signal collected.
	ValueStringArray []string `ch:"ValueStringArray" json:"ValueStringArray"`
}

// SignalToSlice converts a Signal to an array of any for Clickhouse insertion.
// The order of the elements in the array is guaranteed to match the order of elements in the `SignalColNames`.
func SignalToSlice(obj Signal) []any {
	return []any{
		obj.TokenID,
		obj.Timestamp,
		obj.SignalName,
		obj.ValueNumber,
		obj.ValueString,
		obj.ValueStringArray,
	}
}

// SignalColNames returns the column names of the Signal struct.
func SignalColNames() []string {
	return []string{
		"TokenID",
		"Timestamp",
		"SignalName",
		"ValueNumber",
		"ValueString",
		"ValueStringArray",
	}
}
