// Package fingerprint provides decoding for Ruptela fingerprint payloads.
package fingerprint

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
)

const vinLength = 17

type fingerPrintSignals struct {
	Signals signals `json:"signals"`
}
type signals struct {
	VINPart1 string `json:"104"`
	VINPart2 string `json:"105"`
	VINPart3 string `json:"106"`
}

// DecodeFingerprint decodes a fingerprint payload into a FingerprintEvent.
func DecodeFingerprint(payload []byte) (*cloudevent.FingerprintEvent, error) {
	event := cloudevent.CloudEvent[fingerPrintSignals]{}
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal payload: %w", err)
	}
	if event.Data.Signals.VINPart1 == "" || event.Data.Signals.VINPart2 == "" || event.Data.Signals.VINPart3 == "" {
		return nil, fmt.Errorf("missing fingerprint data")
	}
	part1, err := hex.DecodeString(event.Data.Signals.VINPart1)
	if err != nil {
		return nil, fmt.Errorf("could not decode VIN part 1: %w", err)
	}
	part2, err := hex.DecodeString(event.Data.Signals.VINPart2)
	if err != nil {
		return nil, fmt.Errorf("could not decode VIN part 2: %w", err)
	}
	part3, err := hex.DecodeString(event.Data.Signals.VINPart3)
	if err != nil {
		return nil, fmt.Errorf("could not decode VIN part 3: %w", err)
	}
	vinBytes := append(append(part1, part2...), part3...)
	vin := string(vinBytes[:vinLength])
	return &cloudevent.FingerprintEvent{
		CloudEventHeader: event.CloudEventHeader,
		Data: cloudevent.Fingerprint{
			VIN: vin,
		},
	}, nil
}
