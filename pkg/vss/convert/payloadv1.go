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
	ts, err := timestampFromV1Data(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting timestamp: %w", err)
	}
	sub, err := subjectFromV1Data(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting subject: %w", err)
	}
	tokenID, err := tokenGetter.TokenIDFromSubject(ctx, sub)
	if err != nil {
		return nil, fmt.Errorf("error getting tokenID from subject: %w", err)
	}
	sigs, err := vss.SignalsFromV1Data(tokenID, ts, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting signals from v1 data: %w", err)
	}
	return sigs, nil
}

func subjectFromV1Data(jsonData []byte) (string, error) {
	result := gjson.GetBytes(jsonData, "subject")
	if !result.Exists() {
		return "", errors.New("subject field not found")
	}
	sub, ok := result.Value().(string)
	if !ok {
		return "", errors.New("subject field is not a string")
	}
	return sub, nil
}

func timestampFromV1Data(jsonData []byte) (time.Time, error) {
	result := gjson.GetBytes(jsonData, "time")
	if !result.Exists() {
		return time.Time{}, errors.New("time field not found")
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
