// Package status provides a funcions for managing Ruptela status payloads.
package status

import (
	"fmt"

	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
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
	sigs, errs := ruptela.SignalsFromV1Data(baseSignal, jsonData)
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
