// Code generated by github.com/DIMO-Network/model-garage.
package convert

import "math"

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.

// ToChassisAxleRow1WheelLeftTirePressure0 converts data from field 'tires.frontLeft' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelLeftTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelRightTirePressure0 converts data from field 'tires.frontRight' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelRightTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow1WheelRightTirePressure1 converts data from field 'tiresFrontRight' of type float64 to 'Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row1.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow1WheelRightTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelLeftTirePressure0 converts data from field 'tires.backLeft' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelLeftTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelLeftTirePressure1 converts data from field 'tiresBackLeft' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Left.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelLeftTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelRightTirePressure0 converts data from field 'tires.backRight' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelRightTirePressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToChassisAxleRow2WheelRightTirePressure1 converts data from field 'tiresBackRight' of type float64 to 'Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure' of type float64.
// Vehicle.Chassis.Axle.Row2.Wheel.Right.Tire.Pressure: Tire pressure in kilo-Pascal.
// Unit: 'kPa'
func ToChassisAxleRow2WheelRightTirePressure1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationAltitude0 converts data from field 'altitude' of type float64 to 'Vehicle.CurrentLocation.Altitude' of type float64.
// Vehicle.CurrentLocation.Altitude: Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
// Unit: 'm'
func ToCurrentLocationAltitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationIsRedacted0 converts data from field 'isRedacted' of type bool to 'Vehicle.CurrentLocation.IsRedacted' of type float64.
// Vehicle.CurrentLocation.IsRedacted: Indicates if the latitude and longitude signals at the current timestamp have been redacted using a privacy zone.
func ToCurrentLocationIsRedacted0(originalDoc []byte, val bool) (float64, error) {
	if val {
		return 1, nil
	}
	return 0, nil
}

// ToCurrentLocationLatitude0 converts data from field 'latitude' of type float64 to 'Vehicle.CurrentLocation.Latitude' of type float64.
// Vehicle.CurrentLocation.Latitude: Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-90' Max: '90'
func ToCurrentLocationLatitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToCurrentLocationLongitude0 converts data from field 'longitude' of type float64 to 'Vehicle.CurrentLocation.Longitude' of type float64.
// Vehicle.CurrentLocation.Longitude: Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
// Unit: 'degrees' Min: '-180' Max: '180'
func ToCurrentLocationLongitude0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketHDOP0 converts data from field 'hdop' of type float64 to 'Vehicle.DIMO.Aftermarket.HDOP' of type float64.
// Vehicle.DIMO.Aftermarket.HDOP: Horizontal dilution of precision of GPS
func ToDIMOAftermarketHDOP0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketNSAT0 converts data from field 'nsat' of type float64 to 'Vehicle.DIMO.Aftermarket.NSAT' of type float64.
// Vehicle.DIMO.Aftermarket.NSAT: Number of sync satellites for GPS
func ToDIMOAftermarketNSAT0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToDIMOAftermarketSSID0 converts data from field 'ssid' of type string to 'Vehicle.DIMO.Aftermarket.SSID' of type string.
// Vehicle.DIMO.Aftermarket.SSID: Service Set Identifier for the wifi.
func ToDIMOAftermarketSSID0(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketSSID1 converts data from field 'wifi.ssid' of type string to 'Vehicle.DIMO.Aftermarket.SSID' of type string.
// Vehicle.DIMO.Aftermarket.SSID: Service Set Identifier for the wifi.
func ToDIMOAftermarketSSID1(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketWPAState0 converts data from field 'wpa_state' of type string to 'Vehicle.DIMO.Aftermarket.WPAState' of type string.
// Vehicle.DIMO.Aftermarket.WPAState: Indicate the current WPA state for the device's wifi
func ToDIMOAftermarketWPAState0(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToDIMOAftermarketWPAState1 converts data from field 'wifi.wpaState' of type string to 'Vehicle.DIMO.Aftermarket.WPAState' of type string.
// Vehicle.DIMO.Aftermarket.WPAState: Indicate the current WPA state for the device's wifi
func ToDIMOAftermarketWPAState1(originalDoc []byte, val string) (string, error) {
	return val, nil
}

// ToExteriorAirTemperature0 converts data from field 'ambientAirTemp' of type float64 to 'Vehicle.Exterior.AirTemperature' of type float64.
// Vehicle.Exterior.AirTemperature: Air temperature outside the vehicle.
// Unit: 'celsius'
func ToExteriorAirTemperature0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToExteriorAirTemperature1 converts data from field 'ambientTemp' of type float64 to 'Vehicle.Exterior.AirTemperature' of type float64.
// Vehicle.Exterior.AirTemperature: Air temperature outside the vehicle.
// Unit: 'celsius'
func ToExteriorAirTemperature1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToLowVoltageBatteryCurrentVoltage0 converts data from field 'batteryVoltage' of type float64 to 'Vehicle.LowVoltageBattery.CurrentVoltage' of type float64.
// Vehicle.LowVoltageBattery.CurrentVoltage: Current Voltage of the low voltage battery.
// Unit: 'V'
func ToLowVoltageBatteryCurrentVoltage0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDBarometricPressure0 converts data from field 'barometricPressure' of type float64 to 'Vehicle.OBD.BarometricPressure' of type float64.
// Vehicle.OBD.BarometricPressure: PID 33 - Barometric pressure
// Unit: 'kPa'
func ToOBDBarometricPressure0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDEngineLoad0 converts data from field 'engineLoad' of type float64 to 'Vehicle.OBD.EngineLoad' of type float64.
// Vehicle.OBD.EngineLoad: PID 04 - Engine load in percent - 0 = no load, 100 = full load
// Unit: 'percent'
func ToOBDEngineLoad0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		return val * 100, nil
	}
	return val, nil
}

// ToOBDIntakeTemp0 converts data from field 'intakeTemp' of type float64 to 'Vehicle.OBD.IntakeTemp' of type float64.
// Vehicle.OBD.IntakeTemp: PID 0F - Intake temperature
// Unit: 'celsius'
func ToOBDIntakeTemp0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToOBDRunTime0 converts data from field 'runTime' of type float64 to 'Vehicle.OBD.RunTime' of type float64.
// Vehicle.OBD.RunTime: PID 1F - Engine run time
// Unit: 's'
func ToOBDRunTime0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineECT0 converts data from field 'coolantTemp' of type float64 to 'Vehicle.Powertrain.CombustionEngine.ECT' of type float64.
// Vehicle.Powertrain.CombustionEngine.ECT: Engine coolant temperature.
// Unit: 'celsius'
func ToPowertrainCombustionEngineECT0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineEngineOilLevel0 converts data from field 'oil' of type float64 to 'Vehicle.Powertrain.CombustionEngine.EngineOilLevel' of type string.
// Vehicle.Powertrain.CombustionEngine.EngineOilLevel: Engine oil level.
func ToPowertrainCombustionEngineEngineOilLevel0(originalDoc []byte, val float64) (string, error) {
	switch {
	case val < 0.25:
		return "CRITICALLY_LOW", nil
	case val < 0.5:
		return "LOW", nil
	case val < 0.75:
		return "NORMAL", nil
	case val < .99:
		return "HIGH", nil
	default:
		return "CRITICALLY_HIGH", nil
	}
}

// ToPowertrainCombustionEngineEngineOilRelativeLevel0 converts data from field 'oil' of type float64 to 'Vehicle.Powertrain.CombustionEngine.EngineOilRelativeLevel' of type float64.
// Vehicle.Powertrain.CombustionEngine.EngineOilRelativeLevel: Engine oil level as a percentage.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainCombustionEngineEngineOilRelativeLevel0(originalDoc []byte, val float64) (float64, error) {
	// oil comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainCombustionEngineMAF0 converts data from field 'maf' of type float64 to 'Vehicle.Powertrain.CombustionEngine.MAF' of type float64.
// Vehicle.Powertrain.CombustionEngine.MAF: Grams of air drawn into engine per second.
// Unit: 'g/s'
func ToPowertrainCombustionEngineMAF0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineSpeed0 converts data from field 'rpm' of type float64 to 'Vehicle.Powertrain.CombustionEngine.Speed' of type float64.
// Vehicle.Powertrain.CombustionEngine.Speed: Engine speed measured as rotations per minute.
// Unit: 'rpm'
func ToPowertrainCombustionEngineSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineSpeed1 converts data from field 'engineSpeed' of type float64 to 'Vehicle.Powertrain.CombustionEngine.Speed' of type float64.
// Vehicle.Powertrain.CombustionEngine.Speed: Engine speed measured as rotations per minute.
// Unit: 'rpm'
func ToPowertrainCombustionEngineSpeed1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainCombustionEngineTPS0 converts data from field 'throttlePosition' of type float64 to 'Vehicle.Powertrain.CombustionEngine.TPS' of type float64.
// Vehicle.Powertrain.CombustionEngine.TPS: Current throttle position.
// Unit: 'percent'  Max: '100'
func ToPowertrainCombustionEngineTPS0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		return val * 100, nil
	}
	return val, nil
}

// ToPowertrainFuelSystemRelativeLevel0 converts data from field 'fuelLevel' of type float64 to 'Vehicle.Powertrain.FuelSystem.RelativeLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.RelativeLevel: Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainFuelSystemRelativeLevel0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainFuelSystemRelativeLevel1 converts data from field 'fuelPercentRemaining' of type float64 to 'Vehicle.Powertrain.FuelSystem.RelativeLevel' of type float64.
// Vehicle.Powertrain.FuelSystem.RelativeLevel: Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainFuelSystemRelativeLevel1(originalDoc []byte, val float64) (float64, error) {
	// fuelPercentRemaining comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainFuelSystemSupportedFuelTypes0 converts data from field 'fuelType' of type string to 'Vehicle.Powertrain.FuelSystem.SupportedFuelTypes' of type string.
// Vehicle.Powertrain.FuelSystem.SupportedFuelTypes: High level information of fuel types supported
func ToPowertrainFuelSystemSupportedFuelTypes0(originalDoc []byte, val string) (string, error) {
	switch val {
	case "Gasoline":
		return "GASOLINE", nil
	case "Ethanol":
		return "E85", nil
	case "Diesel":
		return "DIESEL", nil
	case "LPG":
		return "LPG", nil
	default:
		return "OTHER", nil
	}
}

// ToPowertrainRange0 converts data from field 'range' of type float64 to 'Vehicle.Powertrain.Range' of type float64.
// Vehicle.Powertrain.Range: Remaining range in meters using all energy sources available in the vehicle.
// Unit: 'm'
func ToPowertrainRange0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTractionBatteryChargingChargeLimit0 converts data from field 'chargeLimit' of type float64 to 'Vehicle.Powertrain.TractionBattery.Charging.ChargeLimit' of type float64.
// Vehicle.Powertrain.TractionBattery.Charging.ChargeLimit: Target charge limit (state of charge) for battery.
// Unit: 'percent' Min: '0' Max: '100'
func ToPowertrainTractionBatteryChargingChargeLimit0(originalDoc []byte, val float64) (float64, error) {
	// chargeLimit comes in as a value between 0 and 1, convert to percentage.
	return val * 100, nil
}

// ToPowertrainTractionBatteryChargingIsCharging0 converts data from field 'charging' of type bool to 'Vehicle.Powertrain.TractionBattery.Charging.IsCharging' of type float64.
// Vehicle.Powertrain.TractionBattery.Charging.IsCharging: True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
func ToPowertrainTractionBatteryChargingIsCharging0(originalDoc []byte, val bool) (float64, error) {
	if val {
		return 1, nil
	}
	return 0, nil
}

// ToPowertrainTractionBatteryCurrentPower0 converts data from field 'charger.power' of type float64 to 'Vehicle.Powertrain.TractionBattery.CurrentPower' of type float64.
// Vehicle.Powertrain.TractionBattery.CurrentPower: Current electrical energy flowing in/out of battery. Positive = Energy flowing in to battery, e.g. during charging. Negative = Energy flowing out of battery, e.g. during driving.
// Unit: 'W'
func ToPowertrainTractionBatteryCurrentPower0(originalDoc []byte, val float64) (float64, error) {
	// V1 field is in kilowatts (kW), VSS field is in watts (W).
	return 1000 * val, nil
}

// ToPowertrainTractionBatteryGrossCapacity0 converts data from field 'batteryCapacity' of type float64 to 'Vehicle.Powertrain.TractionBattery.GrossCapacity' of type float64.
// Vehicle.Powertrain.TractionBattery.GrossCapacity: Gross capacity of the battery.
// Unit: 'kWh'
func ToPowertrainTractionBatteryGrossCapacity0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToPowertrainTractionBatteryStateOfChargeCurrent0 converts data from field 'soc' of type float64 to 'Vehicle.Powertrain.TractionBattery.StateOfCharge.Current' of type float64.
// Vehicle.Powertrain.TractionBattery.StateOfCharge.Current: Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
// Unit: 'percent' Min: '0' Max: '100.0'
func ToPowertrainTractionBatteryStateOfChargeCurrent0(originalDoc []byte, val float64) (float64, error) {
	schemaVersion := GetSchemaVersion(originalDoc)
	if hasV1Schema(schemaVersion) {
		// soc comes in as a value between 0 and 1, convert to percentage.
		return val * 100, nil
	}
	return val, nil
}

// ToPowertrainTransmissionTravelledDistance0 converts data from field 'odometer' of type float64 to 'Vehicle.Powertrain.Transmission.TravelledDistance' of type float64.
// Vehicle.Powertrain.Transmission.TravelledDistance: Odometer reading, total distance travelled during the lifetime of the transmission.
// Unit: 'km'
func ToPowertrainTransmissionTravelledDistance0(originalDoc []byte, val float64) (float64, error) {
	if val > 999999 {
		// if the value is absurdly high, it is likely in meters, convert to kilometers
		// TODO: find a reliable way to determine if the value is in meters
		const metersToKilometers = 1000
		return math.Round(val / metersToKilometers), nil
	}
	return val, nil
}

// ToPowertrainType0 converts data from field 'fuelType' of type string to 'Vehicle.Powertrain.Type' of type string.
// Vehicle.Powertrain.Type: Defines the powertrain type of the vehicle.
func ToPowertrainType0(originalDoc []byte, val string) (string, error) {
	// possible arguments Gasoline, Ethanol, Diesel, Not available, Electric, LPG
	// deault to combustion
	if val == "Electric" {
		return "ELECTRIC", nil
	}
	return "COMBUSTION", nil
}

// ToSpeed0 converts data from field 'vehicleSpeed' of type float64 to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed0(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}

// ToSpeed1 converts data from field 'speed' of type float64 to 'Vehicle.Speed' of type float64.
// Vehicle.Speed: Vehicle speed.
// Unit: 'km/h'
func ToSpeed1(originalDoc []byte, val float64) (float64, error) {
	return val, nil
}
