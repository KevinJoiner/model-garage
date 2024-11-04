package tesla

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

var baseDoc = []byte(`
{
	"charge_state": {
		"battery_level": 23,
		"timestamp": 1730728800
	}
}
`)

var expSignals = []vss.Signal{
	{TokenID: 7, Timestamp: time.Unix(1730728800, 0), Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 23, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
}

func TestSignalsFromTesla(t *testing.T) {
	baseSignal := vss.Signal{TokenID: 7, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"}
	sigs, err := SignalsFromTesla(baseSignal, baseDoc)
	require.Empty(t, err, "Expected no errors.")
	require.ElementsMatch(t, sigs, expSignals)
}
