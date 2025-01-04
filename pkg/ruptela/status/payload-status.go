// Package status provides a funcions for managing Ruptela status payloads.
package status

import (
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

// SignalsFromV1Payload gets a slice signals from a v1 payload.
func SignalsFromV1Payload(jsonData []byte) ([]vss.SignalValue, error) {
	ts, err := TimestampFromV1Data(jsonData)
	if err != nil {
		return nil, convert.SignalValueConversionError{
			Errors: []error{fmt.Errorf("error getting timestamp: %w", err)},
		}
	}

	sigs, errs := ruptela.SignalsFromV1Data(ts, jsonData)
	if errs != nil {
		return nil, convert.SignalValueConversionError{
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}
	return sigs, nil
}
