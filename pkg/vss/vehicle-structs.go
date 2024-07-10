// Code generated by github.com/DIMO-Network/model-garage DO NOT EDIT.
package vss

const (
	// FieldChassisAxleRow1WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow1WheelLeftTirePressure = "chassisAxleRow1WheelLeftTirePressure"
	// FieldChassisAxleRow1WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow1WheelRightTirePressure = "chassisAxleRow1WheelRightTirePressure"
	// FieldChassisAxleRow2WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow2WheelLeftTirePressure = "chassisAxleRow2WheelLeftTirePressure"
	// FieldChassisAxleRow2WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow2WheelRightTirePressure = "chassisAxleRow2WheelRightTirePressure"
	// FieldCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationAltitude = "currentLocationAltitude"
	// FieldCurrentLocationIsRedacted Indicates if the latitude and longitude signals at the current timestamp have been redacted using a privacy zone.
	FieldCurrentLocationIsRedacted = "currentLocationIsRedacted"
	// FieldCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationLatitude = "currentLocationLatitude"
	// FieldCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationLongitude = "currentLocationLongitude"
	// FieldDIMOAftermarketHDOP Horizontal dilution of precision of GPS
	FieldDIMOAftermarketHDOP = "dimoAftermarketHDOP"
	// FieldDIMOAftermarketNSAT Number of sync satellites for GPS
	FieldDIMOAftermarketNSAT = "dimoAftermarketNSAT"
	// FieldDIMOAftermarketSSID Service Set Identifier for the wifi.
	FieldDIMOAftermarketSSID = "dimoAftermarketSSID"
	// FieldDIMOAftermarketWPAState Indicate the current WPA state for the device's wifi
	FieldDIMOAftermarketWPAState = "dimoAftermarketWPAState"
	// FieldExteriorAirTemperature Air temperature outside the vehicle.
	FieldExteriorAirTemperature = "exteriorAirTemperature"
	// FieldLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	FieldLowVoltageBatteryCurrentVoltage = "lowVoltageBatteryCurrentVoltage"
	// FieldOBDBarometricPressure PID 33 - Barometric pressure
	FieldOBDBarometricPressure = "obdBarometricPressure"
	// FieldOBDEngineLoad PID 04 - Engine load in percent - 0 = no load, 100 = full load
	FieldOBDEngineLoad = "obdEngineLoad"
	// FieldOBDIntakeTemp PID 0F - Intake temperature
	FieldOBDIntakeTemp = "obdIntakeTemp"
	// FieldOBDRunTime PID 1F - Engine run time
	FieldOBDRunTime = "obdRunTime"
	// FieldPowertrainCombustionEngineECT Engine coolant temperature.
	FieldPowertrainCombustionEngineECT = "powertrainCombustionEngineECT"
	// FieldPowertrainCombustionEngineEngineOilLevel Engine oil level.
	FieldPowertrainCombustionEngineEngineOilLevel = "powertrainCombustionEngineEngineOilLevel"
	// FieldPowertrainCombustionEngineEngineOilRelativeLevel Engine oil level as a percentage.
	FieldPowertrainCombustionEngineEngineOilRelativeLevel = "powertrainCombustionEngineEngineOilRelativeLevel"
	// FieldPowertrainCombustionEngineMAF Grams of air drawn into engine per second.
	FieldPowertrainCombustionEngineMAF = "powertrainCombustionEngineMAF"
	// FieldPowertrainCombustionEngineSpeed Engine speed measured as rotations per minute.
	FieldPowertrainCombustionEngineSpeed = "powertrainCombustionEngineSpeed"
	// FieldPowertrainCombustionEngineTPS Current throttle position.
	FieldPowertrainCombustionEngineTPS = "powertrainCombustionEngineTPS"
	// FieldPowertrainFuelSystemRelativeLevel Level in fuel tank as percent of capacity. 0 = empty. 100 = full.
	FieldPowertrainFuelSystemRelativeLevel = "powertrainFuelSystemRelativeLevel"
	// FieldPowertrainFuelSystemSupportedFuelTypes High level information of fuel types supported
	FieldPowertrainFuelSystemSupportedFuelTypes = "powertrainFuelSystemSupportedFuelTypes"
	// FieldPowertrainRange Remaining range in meters using all energy sources available in the vehicle.
	FieldPowertrainRange = "powertrainRange"
	// FieldPowertrainTractionBatteryChargingChargeLimit Target charge limit (state of charge) for battery.
	FieldPowertrainTractionBatteryChargingChargeLimit = "powertrainTractionBatteryChargingChargeLimit"
	// FieldPowertrainTractionBatteryChargingIsCharging True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
	FieldPowertrainTractionBatteryChargingIsCharging = "powertrainTractionBatteryChargingIsCharging"
	// FieldPowertrainTractionBatteryCurrentPower Current electrical energy flowing in/out of battery. Positive = Energy flowing in to battery, e.g. during charging. Negative = Energy flowing out of battery, e.g. during driving.
	FieldPowertrainTractionBatteryCurrentPower = "powertrainTractionBatteryCurrentPower"
	// FieldPowertrainTractionBatteryGrossCapacity Gross capacity of the battery.
	FieldPowertrainTractionBatteryGrossCapacity = "powertrainTractionBatteryGrossCapacity"
	// FieldPowertrainTractionBatteryStateOfChargeCurrent Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
	FieldPowertrainTractionBatteryStateOfChargeCurrent = "powertrainTractionBatteryStateOfChargeCurrent"
	// FieldPowertrainTransmissionTravelledDistance Odometer reading, total distance travelled during the lifetime of the transmission.
	FieldPowertrainTransmissionTravelledDistance = "powertrainTransmissionTravelledDistance"
	// FieldPowertrainType Defines the powertrain type of the vehicle.
	FieldPowertrainType = "powertrainType"
	// FieldSpeed Vehicle speed.
	FieldSpeed = "speed"
)
