package convert_test

import (
	"cmp"
	"context"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/DIMO-Network/model-garage/pkg/vss/convert"
	"github.com/stretchr/testify/require"
)

type tokenGetter struct{}

func (*tokenGetter) TokenIDFromSubject(context.Context, string) (uint32, error) {
	return 123, nil
}

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	getter := &tokenGetter{}
	actualSignals, err := convert.SignalsFromPayload(context.Background(), getter, []byte(fullInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	slices.SortFunc(expectedSignals, func(i, j vss.Signal) int {
		return cmp.Compare(i.Name, j.Name)
	})
	require.Equalf(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var fullInputJSON = `{
	"id": "randomIDnumber",
	"specversion": "1.0",
	"source": "SensorXYZ",
	"subject": "Vehicle123",
	"time": "2022-01-01T12:34:56Z",
	"type": "DIMO"
	"data": {
		"tires": {
			"frontLeft": 30.5,
			"frontRight": 31.0,
			"backLeft": 32.2,
			"backRight": 33.1
		},
		"altitude": 100.0,
		"latitude": 37.7749,
		"longitude": -122.4194,
		"timestamp": "2022-01-01T12:34:56Z",
		"definitionID": "123",
		"iD": "456",
		"ambientTemp": 25.0,
		"batteryVoltage": 12.5,
		"barometricPressure": 1013.25,
		"engineLoad": 75.0,
		"intakeTemp": 30.0,
		"runTime": 1200.0,
		"coolantTemp": 90.0,
		"oil": 0.10,
		"engineSpeed": 3000.0,
		"throttlePosition": 50.0,
		"fuelPercentRemaining": 60.0,
		"fuelType": "Gasoline",
		"range": 300.0,
		"chargeLimit": 80.0,
		"charging": true,
		"batteryCapacity": 60.0,
		"soc": 70.0,
		"odometer": 50000.0,
		"speed": 60.0,
		"make": "Toyota",
		"model": "Camry",
		"year": 2020
				"vin": "1234567890"
	},
}`

var ts = time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC)
var expectedSignals = []vss.Signal{
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["chassisAxleRow1WheelLeftTirePressure"],
		ValueNumber: 30.5,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["chassisAxleRow1WheelRightTirePressure"],
		ValueNumber: 31,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["chassisAxleRow2WheelLeftTirePressure"],
		ValueNumber: 32.2,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["chassisAxleRow2WheelRightTirePressure"],
		ValueNumber: 33.1,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["currentLocationAltitude"],
		ValueNumber: 100,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["currentLocationLatitude"],
		ValueNumber: 37.7749,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["currentLocationLongitude"],
		ValueNumber: -122.4194,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["currentLocationTimestamp"],
		ValueNumber: float64(ts.UTC().Unix()),
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainCombustionEngineECT"],
		ValueNumber: 90,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainCombustionEngineEngineOilLevel"],
		ValueString: "CRITICALLY_LOW",
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainCombustionEngineSpeed"],
		ValueNumber: 3000,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainCombustionEngineTPS"],
		ValueNumber: 50,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainFuelSystemAbsoluteLevel"],
		ValueNumber: 60,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainFuelSystemSupportedFuelTypes"],
		ValueString: "GASOLINE",
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainRange"],
		ValueNumber: 300,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainType"],
		ValueString: "COMBUSTION",
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainTractionBatteryChargingChargeLimit"],
		ValueNumber: 80,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainTractionBatteryChargingIsCharging"],
		ValueNumber: 1,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainTractionBatteryGrossCapacity"],
		ValueNumber: 60,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainTractionBatteryStateOfChargeCurrent"],
		ValueNumber: 70,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["powertrainTransmissionTravelledDistance"],
		ValueNumber: 50000,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["speed"],
		ValueNumber: 60,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["vehicleIdentificationBrand"],
		ValueString: "Toyota",
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["vehicleIdentificationModel"],
		ValueString: "Camry",
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["vehicleIdentificationYear"],
		ValueNumber: 2020,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["exteriorAirTemperature"],
		ValueNumber: 25,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["lowVoltageBatteryCurrentVoltage"],
		ValueNumber: 12.5,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["oBDBarometricPressure"],
		ValueNumber: 1013.25,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["oBDEngineLoad"],
		ValueNumber: 75,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["oBDIntakeTemp"],
		ValueNumber: 30,
	},
	{
		TokenID:     123,
		Timestamp:   ts,
		Name:        vss.JSONName2CHName["oBDRunTime"],
		ValueNumber: 1200,
	},
}
