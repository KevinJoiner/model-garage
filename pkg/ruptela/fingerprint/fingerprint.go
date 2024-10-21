// Package fingerprint provides decoding for Ruptela fingerprint payloads.
package fingerprint

import (
	"encoding/json"
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
)

// {
// 	"source": "ruptela/TODO",
// 	"data": {
// 		"pos": {
// 			"alt": 1048,
// 			"dir": 19730,
// 			"hdop": 6,
// 			"lat": 522721466,
// 			"lon": -9014316,
// 			"sat": 20,
// 			"spd": 0
// 		},
// 		"prt": 0,
// 		"signals": {
// 			"102": "0",
// 			"103": "0",
// 			"104": "53414C4C41414146",
// 			"105": "3341413534343438",
// 			"106": "3200000000000000",

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
	vin := event.Data.Signals.VINPart1 + event.Data.Signals.VINPart2 + event.Data.Signals.VINPart3
	return &cloudevent.FingerprintEvent{
		CloudEventHeader: event.CloudEventHeader,
		Data: cloudevent.Fingerprint{
			VIN: vin,
		},
	}, nil
}
