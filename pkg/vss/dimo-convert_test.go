package vss_test

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	vehicle, err := vss.FromData([]byte(fullInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Equalf(t, fullVehicle, vehicle, "converted vehicle does not match expected vehicle")
}

var fullInputJSON = `{
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
				"vehicleID": "456",
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
			"source": "SensorXYZ",
			"subject": "Vehicle123",
			"time": "2022-01-01T12:34:56Z",
			"type": "DIMO"
		}`

var fullVehicle = &vss.Dimo{
	VehicleChassisAxleRow1WheelLeftTirePressure:          ref(uint16(30)),
	VehicleChassisAxleRow1WheelRightTirePressure:         ref(uint16(31)),
	VehicleChassisAxleRow2WheelLeftTirePressure:          ref(uint16(32)),
	VehicleChassisAxleRow2WheelRightTirePressure:         ref(uint16(33)),
	VehicleCurrentLocationAltitude:                       ref(float64(100.0)),
	VehicleCurrentLocationLatitude:                       ref(float64(37.7749)),
	VehicleCurrentLocationLongitude:                      ref(float64(-122.4194)),
	VehicleCurrentLocationTimestamp:                      ref(time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC)),
	VehiclePowertrainCombustionEngineECT:                 ref(int16(90.0)),
	VehiclePowertrainCombustionEngineEngineOilLevel:      ref("CRITICALLY_LOW"),
	VehiclePowertrainCombustionEngineSpeed:               ref(uint16(3000.0)),
	VehiclePowertrainCombustionEngineTPS:                 ref(uint8(50.0)),
	VehiclePowertrainFuelSystemAbsoluteLevel:             ref(float32(60.0)),
	VehiclePowertrainFuelSystemSupportedFuelTypes:        []string{"GASOLINE"},
	VehiclePowertrainRange:                               ref(uint32(300.0)),
	VehiclePowertrainType:                                ref("COMBUSTION"),
	VehiclePowertrainTractionBatteryChargingChargeLimit:  ref(uint8(80.0)),
	VehiclePowertrainTractionBatteryChargingIsCharging:   ref(true),
	VehiclePowertrainTractionBatteryGrossCapacity:        ref(uint16(60.0)),
	VehiclePowertrainTractionBatteryStateOfChargeCurrent: ref(float32(70.0)),
	VehiclePowertrainTransmissionTravelledDistance:       ref(float32(50000.0)),
	VehicleSpeed:                           ref(float32(60.0)),
	VehicleVehicleIdentificationBrand:      ref("Toyota"),
	VehicleVehicleIdentificationModel:      ref("Camry"),
	VehicleVehicleIdentificationYear:       ref(uint16(2020)),
	VehicleExteriorAirTemperature:          ref(float32(25.0)),
	VehicleLowVoltageBatteryCurrentVoltage: ref(float32(12.5)),
	VehicleOBDBarometricPressure:           ref(float32(1013.25)),
	VehicleOBDEngineLoad:                   ref(float32(75.0)),
	VehicleOBDIntakeTemp:                   ref(float32(30.0)),
	VehicleOBDRunTime:                      ref(float32(1200.0)),
}

func ref[T any](t T) *T {
	return &t
}
