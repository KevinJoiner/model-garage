package convert

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

const (
	specV1 = "1.0"
	specV2 = "2.0"
)

// SignalsFromPayload extracts signals from a payload.
// It detects the payload version and calls the appropriate function.
func SignalsFromPayload(ctx context.Context, tokenGetter TokenIDGetter, jsonData []byte) ([]vss.Signal, error) {
	specVersion := gjson.GetBytes(jsonData, "specversion").String()
	switch {
	case specVersion == specV1:
		return SignalsFromV1Payload(ctx, tokenGetter, jsonData)
	case specVersion == specV2:
		return SignalsFromV2Payload(jsonData)
	default:
		return nil, errors.New("unsupported specversion")
	}
}

// SignalsFromV2Payload extracts signals from a V2 payload.
func SignalsFromV2Payload(jsonData []byte) ([]vss.Signal, error) {
	var errs error

	signals := gjson.GetBytes(jsonData, "data.vehicle.signals")
	if !signals.Exists() {
		return nil, errors.New("signals field not found")
	}
	if !signals.IsArray() {
		return nil, errors.New("signals field is not an array")
	}
	tokenID, err := tokenIDFromV2Data(jsonData)
	if err != nil {
		return nil, fmt.Errorf("error getting tokenID: %w", err)
	}
	retSignals := []vss.Signal{}
	for _, sigData := range signals.Array() {
		originalName, err := signalNameFromV2Data(sigData)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		ts, err := timestampFromV2Data(sigData)
		if err != nil {
			err = fmt.Errorf("error for '%s': %w", originalName, err)
			errs = errors.Join(errs, err)
			continue
		}
		sigs, err := vss.SignalsFromV2Data(tokenID, ts, originalName, sigData)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		retSignals = append(retSignals, sigs...)
	}
	return retSignals, errs
}

func tokenIDFromV2Data(jsonData []byte) (uint32, error) {
	tokenID := gjson.GetBytes(jsonData, "vehicleTokenID")
	if !tokenID.Exists() {
		return 0, errors.New("vehicleTokenID field not found")
	}
	id, ok := tokenID.Value().(float64)
	if !ok {
		return 0, errors.New("vehicleTokenID field is not a number")
	}
	return float64toUint32(id), nil
}

func timestampFromV2Data(sigResult gjson.Result) (time.Time, error) {
	timestamp := sigResult.Get("timestamp")
	if !timestamp.Exists() {
		return time.Time{}, errors.New("timestamp field not found")
	}
	return time.UnixMilli(int64(timestamp.Uint())).UTC(), nil
}

func signalNameFromV2Data(sigResult gjson.Result) (string, error) {
	signalName := sigResult.Get("name")
	if !signalName.Exists() {
		return "", errors.New("signalName field not found")
	}
	return signalName.String(), nil
}

// float64toUint32 converts float64 to uint32.
func float64toUint32(val float64) uint32 {
	if val > math.MaxUint32 {
		return math.MaxUint32
	}
	if val < 0 {
		return 0
	}
	return uint32(val)
}
