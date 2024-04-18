package convert

import (
	"errors"
	"fmt"
	"time"

	"github.com/DIMO-Network/model-garage/internal/convert"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

func SignalsFromV2Payload(jsonData []byte) ([]vss.Signal, error) {
	var errs error

	signals := gjson.GetBytes(jsonData, "data.signals")
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
		ts, err := timestampFromV2Data(sigData)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		originalName, err := signalNameFromV2Data(sigData)
		if err != nil {
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
	tokenID := gjson.GetBytes(jsonData, "data.tokenId")
	if !tokenID.Exists() {
		return 0, errors.New("tokenID field not found")
	}
	id, ok := tokenID.Value().(float64)
	if !ok {
		return 0, errors.New("tokenID field is not a number")
	}
	return convert.Float64toUint32(id), nil
}

func timestampFromV2Data(sigResult gjson.Result) (time.Time, error) {
	timestamp := sigResult.Get("timestamp")
	if !timestamp.Exists() {
		return time.Time{}, errors.New("timestamp field not found")
	}
	return time.Parse(time.RFC3339, timestamp.String())
}

func signalNameFromV2Data(sigResult gjson.Result) (string, error) {
	signalName := sigResult.Get("name")
	if !signalName.Exists() {
		return "", errors.New("signalName field not found")
	}
	return signalName.String(), nil
}
