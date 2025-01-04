// Package status provides a functions for managing Autopi status payloads.
package status

import (
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/autopi"
	"github.com/tidwall/gjson"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromV2Payload extracts signals from a V2 payload.
func SignalsFromV2Payload(jsonData []byte) ([]vss.SignalValue, error) {
	signals := gjson.GetBytes(jsonData, "data.vehicle.signals")
	if !signals.Exists() {
		return nil, convert.SignalValueConversionError{
			Errors: []error{convert.FieldNotFoundError{Field: "signals", Lookup: "data.vehicle.signals"}},
		}
	}
	if !signals.IsArray() {
		if signals.Value() == nil {
			// If the signals array is NULL treat it like an empty array.
			return []vss.SignalValue{}, nil
		}
		return nil, convert.SignalValueConversionError{
			Errors: []error{errors.New("signals field is not an array")},
		}
	}
	retSignals := []vss.SignalValue{}

	SignalValueConversionErrors := convert.SignalValueConversionError{}
	for _, sigData := range signals.Array() {
		originalName, err := NameFromV2Signal(sigData)
		if err != nil {
			SignalValueConversionErrors.Errors = append(SignalValueConversionErrors.Errors, err)
			continue
		}
		ts, err := TimestampFromV2Signal(sigData)
		if err != nil {
			err = fmt.Errorf("error for '%s': %w", originalName, err)
			SignalValueConversionErrors.Errors = append(SignalValueConversionErrors.Errors, err)
			continue
		}
		sigs, err := autopi.SignalsFromV2Data(jsonData, ts, originalName, sigData)
		if err != nil {
			SignalValueConversionErrors.Errors = append(SignalValueConversionErrors.Errors, err)
			continue
		}
		retSignals = append(retSignals, sigs...)
	}

	if len(SignalValueConversionErrors.Errors) > 0 {
		SignalValueConversionErrors.DecodedSignals = retSignals
		return nil, SignalValueConversionErrors
	}
	return retSignals, nil
}

// TimestampFromV2Signal gets a timestamp from a V2 signal.
func TimestampFromV2Signal(sigResult gjson.Result) (time.Time, error) {
	lookupKey := "timestamp"
	timestamp := sigResult.Get(lookupKey)
	if !timestamp.Exists() {
		return time.Time{}, convert.FieldNotFoundError{Field: "timestamp", Lookup: lookupKey}
	}
	return time.UnixMilli(timestamp.Int()).UTC(), nil
}

// NameFromV2Signal gets a name from a V2 signal.
func NameFromV2Signal(sigResult gjson.Result) (string, error) {
	lookupKey := "name"
	signalName := sigResult.Get(lookupKey)
	if !signalName.Exists() {
		return "", convert.FieldNotFoundError{Field: "name", Lookup: lookupKey}
	}
	return signalName.String(), nil
}
