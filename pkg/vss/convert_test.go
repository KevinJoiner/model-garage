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

var fullVehicle = &vss.Vehicle{
	ChassisAxleRow1WheelLeftTirePressure:          30,
	ChassisAxleRow1WheelRightTirePressure:         31,
	ChassisAxleRow2WheelLeftTirePressure:          32,
	ChassisAxleRow2WheelRightTirePressure:         33,
	CurrentLocationAltitude:                       100.0,
	CurrentLocationLatitude:                       37.7749,
	CurrentLocationLongitude:                      -122.4194,
	CurrentLocationTimestamp:                      time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC),
	PowertrainCombustionEngineECT:                 90.0,
	PowertrainCombustionEngineEngineOilLevel:      "CRITICALLY_LOW",
	PowertrainCombustionEngineSpeed:               3000.0,
	PowertrainCombustionEngineTPS:                 50.0,
	PowertrainFuelSystemAbsoluteLevel:             60.0,
	PowertrainFuelSystemSupportedFuelTypes:        []string{"GASOLINE"},
	PowertrainRange:                               300.0,
	PowertrainType:                                "COMBUSTION",
	PowertrainTractionBatteryChargingChargeLimit:  80.0,
	PowertrainTractionBatteryChargingIsCharging:   true,
	PowertrainTractionBatteryGrossCapacity:        60.0,
	PowertrainTractionBatteryStateOfChargeCurrent: 70.0,
	PowertrainTransmissionTravelledDistance:       50000.0,
	Speed:                                         60.0,
	VehicleIdentificationBrand:                    "Toyota",
	VehicleIdentificationModel:                    "Camry",
	VehicleIdentificationYear:                     2020,
	VehicleIdentificationVIN:                      "1234567890",
	DIMODefinitionID:                              "123",
	DIMOSource:                                    "SensorXYZ",
	DIMOSubject:                                   "Vehicle123",
	DIMOTimestamp:                                 time.Date(2022, 1, 1, 12, 34, 56, 0, time.UTC),
	DIMOType:                                      "DIMO",
	DIMOVehicleID:                                 "456",
	ExteriorAirTemperature:                        25.0,
	LowVoltageBatteryCurrentVoltage:               12.5,
	OBDBarometricPressure:                         1013.25,
	OBDEngineLoad:                                 75.0,
	OBDIntakeTemp:                                 30.0,
	OBDRunTime:                                    1200.0,
}
