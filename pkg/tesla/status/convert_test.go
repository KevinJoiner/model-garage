package status

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var baseDoc = []byte(`
{
	"subject": "did:nft:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_37",
	"source": "0x983110309620D911731Ac0932219af06091b6744",
	"data": {
		"charge_state": {
			"battery_level": 23,
			"battery_range": 341,
			"charge_energy_added": 42,
			"charge_limit_soc": 80,
			"charging_state": "Charging",
			"timestamp": 1730728800000
		},
		"climate_state": {
			"outside_temp": 19,
			"timestamp": 1730728802000
		},
		"drive_state": {
			"latitude": 38.89,
			"longitude": 77.03,
			"power": -7,
			"speed": 25,
			"timestamp": 1730738800000
		},
		"vehicle_state": {
			"odometer": 5633,
			"tpms_pressure_fl": 3.12,
			"tpms_pressure_fr": 3.09,
			"tpms_pressure_rl": 2.98,
			"tpms_pressure_rr": 2.99,
			"timestamp": 1730728805000
		}
}	}
`)

const teslaConnection = "0x983110309620D911731Ac0932219af06091b6744"

var expSignals = []vss.SignalValue{
	{Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 312},
	{Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 309},
	{Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 298},
	{Timestamp: time.UnixMilli(1730728805000), Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 299},
	{Timestamp: time.UnixMilli(1730738800000), Name: "currentLocationLatitude", ValueNumber: 38.89},
	{Timestamp: time.UnixMilli(1730738800000), Name: "currentLocationLongitude", ValueNumber: 77.03},
	{Timestamp: time.UnixMilli(1730728802000), Name: "exteriorAirTemperature", ValueNumber: 19},
	{Timestamp: time.UnixMilli(1730728800000), Name: "powertrainRange", ValueNumber: 548.7863040000001},
	{Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingAddedEnergy", ValueNumber: 42},
	{Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80},
	{Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1},
	{Timestamp: time.UnixMilli(1730738800000), Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 7000},
	{Timestamp: time.UnixMilli(1730728800000), Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 23},
	{Timestamp: time.UnixMilli(1730728805000), Name: "powertrainTransmissionTravelledDistance", ValueNumber: 9065.434752000001},
	{Timestamp: time.UnixMilli(1730738800000), Name: "speed", ValueNumber: 40.2336},
}

func TestSignalsFromTesla(t *testing.T) {
	computedSignals, err := Decode(baseDoc)
	require.Empty(t, err, "Expected no errors.")
	assert.ElementsMatch(t, computedSignals, expSignals)
}
