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
	},
	"drive_state": {
		"latitude": 38.89,
		"longitude": 77.03,
		"timestamp": 1730738800
	},
	"vehicle_state": {
		"odometer": 5633,
		"tpms_pressure_fl": 3.12,
		"tpms_pressure_fr": 3.09,
		"tpms_pressure_rl": 2.98,
		"tpms_pressure_rr": 2.99,
		"timestamp": 1730728805
	}
}
`)

var expSignals = []vss.Signal{
	{TokenID: 7, Timestamp: time.Unix(1730728800, 0), Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 23, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730728805, 0), Name: "powertrainTransmissionTravelledDistance", ValueNumber: 9065.434752000001, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730728805, 0), Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 312, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730728805, 0), Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 309, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730728805, 0), Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 298, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730728805, 0), Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 299, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730738800, 0), Name: "currentLocationLatitude", ValueNumber: 38.89, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
	{TokenID: 7, Timestamp: time.Unix(1730738800, 0), Name: "currentLocationLongitude", ValueNumber: 77.03, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"},
}

func TestSignalsFromTesla(t *testing.T) {
	baseSignal := vss.Signal{TokenID: 7, Source: "dimo/integration/26A5Dk3vvvQutjSyF0Jka2DP5lg"}
	computedSignals, err := SignalsFromTesla(baseSignal, baseDoc)
	require.Empty(t, err, "Expected no errors.")
	require.ElementsMatch(t, computedSignals, expSignals)
}
