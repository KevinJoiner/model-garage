// Package convert provides a function to generate conversion functions for a vehicle struct.
package convert

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

// TokenIDGetter is an interface to get a tokenID from a subject.
type TokenIDGetter interface {
	TokenIDFromSubject(ctx context.Context, subject string) (uint32, error)
}

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(ctx context.Context, tokenGetter TokenIDGetter, jsonData []byte) ([]vss.Signal, error) {
	ts, err := TimestampFromV1Data(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting timestamp: %w", err)
	}
	tokenID, err := TokenIDFromV1Data(jsonData, tokenGetter)
	if err != nil {
		return nil, err
	}

	source, err := SourceFromV1Data(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting source: %w", err)
	}
	baseSignal := vss.Signal{
		TokenID:   tokenID,
		Timestamp: ts,
		Source:    source,
	}
	sigs, err := SignalsFromV1Data(baseSignal, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting signals from v1 data: %w", err)
	}
	return sigs, nil
}

// SubjectFromV1Data gets a subject from a v1 payload.
func SubjectFromV1Data(jsonData []byte) (string, error) {
	result := gjson.GetBytes(jsonData, "subject")
	if !result.Exists() {
		return "", FieldNotFoundError{Field: "subject", Lookup: "subject"}
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
		return time.Time{}, FieldNotFoundError{Field: "timestamp", Lookup: "time"}
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

// TokenIDFromV1Data gets a tokenID from a v1 payload.
// attempts to get the tokenID from the vehicleTokenId field.
// If the field is not found, it will attempt to get the tokenId using the subject field and the tokenGetter.
func TokenIDFromV1Data(jsonData []byte, tokenGetter TokenIDGetter) (uint32, error) {
	tokenID, err := TokenIDFromData(jsonData)
	if err == nil {
		return tokenID, nil
	}
	if !errors.As(err, &FieldNotFoundError{}) {
		return 0, err
	}

	sub, err := SubjectFromV1Data(jsonData)
	if err != nil {
		return 0, err
	}
	tokenID, err = tokenGetter.TokenIDFromSubject(context.Background(), sub)
	if err != nil {
		return 0, fmt.Errorf("error getting tokenID: %w", err)
	}
	return tokenID, nil
}

// SourceFromV1Data gets a source field from a v1 payload.
func SourceFromV1Data(jsonData []byte) (string, error) {
	result := gjson.GetBytes(jsonData, "source")
	if !result.Exists() {
		return "", FieldNotFoundError{Field: "source", Lookup: "source"}
	}
	source, ok := result.Value().(string)
	if !ok {
		return "", errors.New("source field is not a string")
	}
	return source, nil
}
