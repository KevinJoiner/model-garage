// Package vss provides a models and conversion functions for the vss vehicles.
package vss

import (
	"time"

	"github.com/DIMO-Network/model-garage/internal/convert"
)

// This file is automatically populated with conversion functions for each field of a model struct.
// any conversion functions already defined in this package will not be generated.
// Code generated by model-garage.

// ToVehicleChassisAxleRow1WheelLeftTirePressure converts data as float64 to uint16.
func ToVehicleChassisAxleRow1WheelLeftTirePressure(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehicleChassisAxleRow1WheelRightTirePressure converts data as float64 to uint16.
func ToVehicleChassisAxleRow1WheelRightTirePressure(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehicleChassisAxleRow2WheelLeftTirePressure converts data as float64 to uint16.
func ToVehicleChassisAxleRow2WheelLeftTirePressure(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehicleChassisAxleRow2WheelRightTirePressure converts data as float64 to uint16.
func ToVehicleChassisAxleRow2WheelRightTirePressure(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehicleCurrentLocationAltitude converts data as float64 to float64.
func ToVehicleCurrentLocationAltitude(val float64) (float64, error) {
	return val, nil
}

// ToVehicleCurrentLocationLatitude converts data as float64 to float64.
func ToVehicleCurrentLocationLatitude(val float64) (float64, error) {
	return val, nil
}

// ToVehicleCurrentLocationLongitude converts data as float64 to float64.
func ToVehicleCurrentLocationLongitude(val float64) (float64, error) {
	return val, nil
}

// ToVehicleCurrentLocationTimestamp converts data as string to time.Time.
func ToVehicleCurrentLocationTimestamp(val string) (time.Time, error) {
	return time.Parse(time.RFC3339, val)
}

// ToDefinitionID converts data as string to string.
func ToDefinitionID(val string) (string, error) {
	return val, nil
}

// ToSource converts data as string to string.
func ToSource(val string) (string, error) {
	return val, nil
}

// ToSubject converts data as string to string.
func ToSubject(val string) (string, error) {
	return val, nil
}

// ToTimestamp converts data as string to time.Time.
func ToTimestamp(val string) (time.Time, error) {
	return time.Parse(time.RFC3339, val)
}

// ToType converts data as string to string.
func ToType(val string) (string, error) {
	return val, nil
}

// ToVehicleID converts data as string to string.
func ToVehicleID(val string) (string, error) {
	return val, nil
}

// ToVehicleExteriorAirTemperature converts data as float64 to float32.
func ToVehicleExteriorAirTemperature(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleLowVoltageBatteryCurrentVoltage converts data as float64 to float32.
func ToVehicleLowVoltageBatteryCurrentVoltage(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleOBDBarometricPressure converts data as float64 to float32.
func ToVehicleOBDBarometricPressure(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleOBDEngineLoad converts data as float64 to float32.
func ToVehicleOBDEngineLoad(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleOBDIntakeTemp converts data as float64 to float32.
func ToVehicleOBDIntakeTemp(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleOBDRunTime converts data as float64 to float32.
func ToVehicleOBDRunTime(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehiclePowertrainCombustionEngineECT converts data as float64 to int16.
func ToVehiclePowertrainCombustionEngineECT(val float64) (int16, error) {
	return convert.Float64toInt16(val), nil
}

// ToVehiclePowertrainCombustionEngineEngineOilLevel converts data as float64 to string.
func ToVehiclePowertrainCombustionEngineEngineOilLevel(oilLevel float64) (string, error) {
	switch {
	case oilLevel < 0.25:
		return "CRITICALLY_LOW", nil
	case oilLevel < 0.5:
		return "LOW", nil
	case oilLevel < 0.75:
		return "NORMAL", nil
	case oilLevel < .99:
		return "HIGH", nil
	default:
		return "CRITICALLY_HIGH", nil
	}
}

// ToVehiclePowertrainCombustionEngineSpeed converts data as float64 to uint16.
func ToVehiclePowertrainCombustionEngineSpeed(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehiclePowertrainCombustionEngineTPS converts data as float64 to uint8.
func ToVehiclePowertrainCombustionEngineTPS(val float64) (uint8, error) {
	return convert.Float64toUint8(val), nil
}

// ToVehiclePowertrainFuelSystemAbsoluteLevel converts data as float64 to float32.
func ToVehiclePowertrainFuelSystemAbsoluteLevel(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehiclePowertrainFuelSystemSupportedFuelTypes converts data as string to []string.
func ToVehiclePowertrainFuelSystemSupportedFuelTypes(val string) ([]string, error) {
	switch val {
	case "Gasoline":
		return []string{"GASOLINE"}, nil
	case "Ethanol":
		return []string{"E85"}, nil
	case "Diesel":
		return []string{"DIESEL"}, nil
	case "LPG":
		return []string{"LPG"}, nil
	default:
		return []string{"OTHER"}, nil
	}
}

// ToVehiclePowertrainRange converts data as float64 to uint32.
func ToVehiclePowertrainRange(val float64) (uint32, error) {
	return convert.Float64toUint32(val), nil
}

// ToVehiclePowertrainTractionBatteryChargingChargeLimit converts data as float64 to uint8.
func ToVehiclePowertrainTractionBatteryChargingChargeLimit(val float64) (uint8, error) {
	return convert.Float64toUint8(val), nil
}

// ToVehiclePowertrainTractionBatteryChargingIsCharging converts data as bool to bool.
func ToVehiclePowertrainTractionBatteryChargingIsCharging(val bool) (bool, error) {
	return val, nil
}

// ToVehiclePowertrainTractionBatteryGrossCapacity converts data as float64 to uint16.
func ToVehiclePowertrainTractionBatteryGrossCapacity(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehiclePowertrainTractionBatteryStateOfChargeCurrent converts data as float64 to float32.
func ToVehiclePowertrainTractionBatteryStateOfChargeCurrent(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehiclePowertrainTransmissionTravelledDistance converts data as float64 to float32.
func ToVehiclePowertrainTransmissionTravelledDistance(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleSpeed converts data as float64 to float32.
func ToVehicleSpeed(val float64) (float32, error) {
	return convert.Float64ToFloat32(val), nil
}

// ToVehicleVehicleIdentificationBrand converts data as string to string.
func ToVehicleVehicleIdentificationBrand(val string) (string, error) {
	return val, nil
}

// ToVehicleVehicleIdentificationModel converts data as string to string.
func ToVehicleVehicleIdentificationModel(val string) (string, error) {
	return val, nil
}

// ToVehicleVehicleIdentificationYear converts data as float64 to uint16.
func ToVehicleVehicleIdentificationYear(val float64) (uint16, error) {
	return convert.Float64toUint16(val), nil
}

// ToVehicleVehicleIdentificationVIN converts data as string to string.
func ToVehicleVehicleIdentificationVIN(val string) (string, error) {
	return val, nil
}

// ToVehiclePowertrainType converts data as string to string.
func ToVehiclePowertrainType(val string) (string, error) {
	// possible arguments Gasoline, Ethanol, Diesel, Not available, Electric, LPG
	// deault to combustion
	if val == "Electric" {
		return "ELECTRIC", nil
	}
	return "COMBUSTION", nil
}
