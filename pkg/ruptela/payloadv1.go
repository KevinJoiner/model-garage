// Package ruptela provides a funcions for managing Ruptela payloads.
package ruptela

import (
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(jsonData []byte) ([]vss.Signal, error) {
	ts, err := TimestampFromV1Data(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting timestamp: %w", err)},
		}
	}
	tokenID, err := TokenIDFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			Errors: []error{fmt.Errorf("error getting tokenId: %w", err)},
		}
	}
	source, err := SourceFromData(jsonData)
	if err != nil {
		return nil, convert.ConversionError{
			TokenID: tokenID,
			Errors:  []error{fmt.Errorf("error getting source: %w", err)},
		}
	}

	baseSignal := vss.Signal{
		TokenID:   tokenID,
		Timestamp: ts,
		Source:    source,
	}
	sigs, errs := SignalsFromV1Data(baseSignal, jsonData)
	if errs != nil {
		return nil, convert.ConversionError{
			TokenID:        tokenID,
			Source:         source,
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	return sigs, nil
}

// SubjectFromV1Data gets a subject from a v1 payload.
func SubjectFromV1Data(jsonData []byte) (string, error) {
	result := gjson.GetBytes(jsonData, "subject")
	if !result.Exists() {
		return "", convert.FieldNotFoundError{Field: "subject", Lookup: "subject"}
	}
	sub, ok := result.Value().(string)
	if !ok {
		return "", errors.New("subject field is not a string")
	}
	return sub, nil
}

// TimestampFromV1Data gets a timestamp from a v1 payload.
func TimestampFromV1Data(jsonData []byte) (time.Time, error) {
	result := gjson.GetBytes(jsonData, "time")
	if !result.Exists() {
		return time.Time{}, convert.FieldNotFoundError{Field: "timestamp", Lookup: "time"}
	}

	timeStr, ok := result.Value().(string)
	if !ok {
		ms, ok := result.Value().(float64)
		if ok {
			return time.UnixMilli(int64(ms)), nil
		}
		return time.Time{}, errors.New("time field is not a string or float64")
	}
	ts, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %w", err)
	}
	return ts, nil
}

// TokenIDFromData gets a tokenID from a V2 payload.
func TokenIDFromData(jsonData []byte) (uint32, error) {
	lookupKey := "subject"
	subject := gjson.GetBytes(jsonData, lookupKey)
	if !subject.Exists() {
		return 0, convert.FieldNotFoundError{Field: "tokenID", Lookup: lookupKey}
	}
	subjectStr, ok := subject.Value().(string)
	if !ok {
		return 0, fmt.Errorf("%s field is not a string", lookupKey)
	}
	subjectDID, err := convert.DecodeDID(subjectStr)
	if err != nil {
		return 0, fmt.Errorf("error decoding subject: %w", err)
	}
	return uint32(subjectDID.TokenID), nil
}

// SourceFromData gets a source from a V2 payload.
func SourceFromData(jsonData []byte) (string, error) {
	lookupKey := "source"
	source := gjson.GetBytes(jsonData, lookupKey)
	if !source.Exists() {
		return "", convert.FieldNotFoundError{Field: "source", Lookup: lookupKey}
	}
	src, ok := source.Value().(string)
	if !ok {
		return "", errors.New("source field is not a string")
	}
	return src, nil
}
