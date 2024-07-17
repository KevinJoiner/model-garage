package convert

import (
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromV2Payload extracts signals from a V2 payload.
func SignalsFromV2Payload(jsonData []byte) ([]vss.Signal, error) {
	var errs error

	signals := gjson.GetBytes(jsonData, "data.vehicle.signals")
	if !signals.Exists() {
		return nil, FieldNotFoundError{Field: "signals", Lookup: "data.vehicle.signals"}
	}
	if !signals.IsArray() {
		return nil, errors.New("signals field is not an array")
	}
	tokenID, err := TokenIDFromData(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting tokenID: %w", err)
	}
	source, err := SourceFromData(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting source: %w", err)
	}
	retSignals := []vss.Signal{}
	signalMeta := vss.Signal{
		TokenID: tokenID,
		Source:  source,
	}
	for _, sigData := range signals.Array() {
		originalName, err := NameFromV2Signal(sigData)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		ts, err := TimestampFromV2Signal(sigData)
		if err != nil {
			err = fmt.Errorf("error for '%s': %w", originalName, err)
			errs = errors.Join(errs, err)
			continue
		}
		signalMeta.Timestamp = ts
		sigs, err := SignalsFromV2Data(jsonData, signalMeta, originalName, sigData)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		retSignals = append(retSignals, sigs...)
	}
	return retSignals, errs
}

// TimestampFromV2Signal gets a timestamp from a V2 signal.
func TimestampFromV2Signal(sigResult gjson.Result) (time.Time, error) {
	lookupKey := "timestamp"
	timestamp := sigResult.Get(lookupKey)
	if !timestamp.Exists() {
		return time.Time{}, FieldNotFoundError{Field: "timestamp", Lookup: lookupKey}
	}
	return time.UnixMilli(int64(timestamp.Uint())).UTC(), nil
}

// NameFromV2Signal gets a name from a V2 signal.
func NameFromV2Signal(sigResult gjson.Result) (string, error) {
	lookupKey := "name"
	signalName := sigResult.Get(lookupKey)
	if !signalName.Exists() {
		return "", FieldNotFoundError{Field: "name", Lookup: lookupKey}
	}
	return signalName.String(), nil
}
