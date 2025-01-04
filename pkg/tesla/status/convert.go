// Package status converts Tesla CloudEvents to ClickHouse-ready slices of signals.
package status

import (
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/DIMO-Network/model-garage/pkg/tesla"
	"github.com/DIMO-Network/model-garage/pkg/vss"
)

func Decode(msgBytes []byte) ([]vss.SignalValue, error) {
	// Only interested in the top-level CloudEvent fields.

	sigs, errs := tesla.SignalsFromTesla(msgBytes)
	if len(errs) != 0 {
		return nil, convert.SignalValueConversionError{
			DecodedSignals: sigs,
			Errors:         errs,
		}
	}

	return sigs, nil
}
