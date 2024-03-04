package vss_test

import (
	"testing"
	"time"

	"github.com/KevinJoiner/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestFullFromDataConversion(t *testing.T) {
	vehicle, err := vss.FromData([]byte(fullInputJSON), false)
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
var fullVehicle = &vss.Vehicle{
	VehicleChassisAxleRow1WheelLeftTirePressure:          30,
	VehicleChassisAxleRow1WheelRightTirePressure:         31,
	VehicleChassisAxleRow2WheelLeftTirePressure:          32,
	VehicleChassisAxleRow2WheelRightTirePressure:         33,
	VehicleCurrentLocationAltitude:                       100.0,
	VehicleCurrentLocationLatitude:                       37.7749,
	VehicleCurrentLocationLongitude:                      -122.4194,
	VehicleCurrentLocationTimestamp:                      time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC),
	VehiclePowertrainCombustionEngineECT:                 90.0,
	VehiclePowertrainCombustionEngineEngineOilLevel:      "CRITICALLY_LOW",
	VehiclePowertrainCombustionEngineSpeed:               3000.0,
	VehiclePowertrainCombustionEngineTPS:                 50.0,
	VehiclePowertrainFuelSystemAbsoluteLevel:             60.0,
	VehiclePowertrainFuelSystemSupportedFuelTypes:        []string{"Gasoline"},
	VehiclePowertrainRange:                               300.0,
	VehiclePowertrainTractionBatteryChargingChargeLimit:  80.0,
	VehiclePowertrainTractionBatteryChargingIsCharging:   true,
	VehiclePowertrainTractionBatteryGrossCapacity:        60.0,
	VehiclePowertrainTractionBatteryStateOfChargeCurrent: 70.0,
	VehiclePowertrainTransmissionTravelledDistance:       50000.0,
	VehicleSpeed:                           60.0,
	VehicleVehicleIdentificationBrand:      "Toyota",
	VehicleVehicleIdentificationModel:      "Camry",
	VehicleVehicleIdentificationYear:       2020,
	VehicleVehicleIdentificationVIN:        "1234567890",
	VehicleDIMODefinitionID:                "123",
	VehicleDIMOSource:                      "SensorXYZ",
	VehicleDIMOSubject:                     "Vehicle123",
	VehicleDIMOTimestamp:                   time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC),
	VehicleDIMOType:                        "DIMO",
	VehicleDIMOVehicleID:                   "456",
	VehicleExteriorAirTemperature:          25.0,
	VehicleLowVoltageBatteryCurrentVoltage: 12.5,
	VehicleOBDBarometricPressure:           1013.25,
	VehicleOBDEngineLoad:                   75.0,
	VehicleOBDIntakeTemp:                   30.0,
	VehicleOBDRunTime:                      1200.0,
}
