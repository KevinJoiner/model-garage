// Code generated by "model-garage" DO NOT EDIT.
package vss

const (
	// FieldChassisAxleRow1WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow1WheelLeftTirePressure = "Chassis_Axle_Row1_Wheel_Left_Tire_Pressure"
	// FieldChassisAxleRow1WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow1WheelRightTirePressure = "Chassis_Axle_Row1_Wheel_Right_Tire_Pressure"
	// FieldChassisAxleRow2WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow2WheelLeftTirePressure = "Chassis_Axle_Row2_Wheel_Left_Tire_Pressure"
	// FieldChassisAxleRow2WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldChassisAxleRow2WheelRightTirePressure = "Chassis_Axle_Row2_Wheel_Right_Tire_Pressure"
	// FieldCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationAltitude = "CurrentLocation_Altitude"
	// FieldCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationLatitude = "CurrentLocation_Latitude"
	// FieldCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldCurrentLocationLongitude = "CurrentLocation_Longitude"
	// FieldCurrentLocationTimestamp Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	FieldCurrentLocationTimestamp = "CurrentLocation_Timestamp"
	// FieldDIMOAftermarketHDOP Horizontal dilution of precision of GPS
	FieldDIMOAftermarketHDOP = "DIMO_Aftermarket_HDOP"
	// FieldDIMOAftermarketNSAT Number of sync satellites for GPS
	FieldDIMOAftermarketNSAT = "DIMO_Aftermarket_NSAT"
	// FieldDIMOAftermarketSSID Service Set Ientifier for the wifi.
	FieldDIMOAftermarketSSID = "DIMO_Aftermarket_SSID"
	// FieldDIMOAftermarketWPAState Indicate the current wpa state for the devices wifi
	FieldDIMOAftermarketWPAState = "DIMO_Aftermarket_WPAState"
	// FieldExteriorAirTemperature Air temperature outside the vehicle.
	FieldExteriorAirTemperature = "Exterior_AirTemperature"
	// FieldLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	FieldLowVoltageBatteryCurrentVoltage = "LowVoltageBattery_CurrentVoltage"
	// FieldOBDBarometricPressure PID 33 - Barometric pressure
	FieldOBDBarometricPressure = "OBD_BarometricPressure"
	// FieldOBDEngineLoad PID 04 - Engine load in percent - 0 = no load, 100 = full load
	FieldOBDEngineLoad = "OBD_EngineLoad"
	// FieldOBDIntakeTemp PID 0F - Intake temperature
	FieldOBDIntakeTemp = "OBD_IntakeTemp"
	// FieldOBDRunTime PID 1F - Engine run time
	FieldOBDRunTime = "OBD_RunTime"
	// FieldPowertrainCombustionEngineECT Engine coolant temperature.
	FieldPowertrainCombustionEngineECT = "Powertrain_CombustionEngine_ECT"
	// FieldPowertrainCombustionEngineEngineOilLevel Engine oil level.
	FieldPowertrainCombustionEngineEngineOilLevel = "Powertrain_CombustionEngine_EngineOilLevel"
	// FieldPowertrainCombustionEngineMAF Grams of air drawn into engine per second.
	FieldPowertrainCombustionEngineMAF = "Powertrain_CombustionEngine_MAF"
	// FieldPowertrainCombustionEngineSpeed Engine speed measured as rotations per minute.
	FieldPowertrainCombustionEngineSpeed = "Powertrain_CombustionEngine_Speed"
	// FieldPowertrainCombustionEngineTPS Current throttle position.
	FieldPowertrainCombustionEngineTPS = "Powertrain_CombustionEngine_TPS"
	// FieldPowertrainFuelSystemAbsoluteLevel Current available fuel in the fuel tank expressed in liters.
	FieldPowertrainFuelSystemAbsoluteLevel = "Powertrain_FuelSystem_AbsoluteLevel"
	// FieldPowertrainFuelSystemSupportedFuelTypes High level information of fuel types supported
	FieldPowertrainFuelSystemSupportedFuelTypes = "Powertrain_FuelSystem_SupportedFuelTypes"
	// FieldPowertrainRange Remaining range in meters using all energy sources available in the vehicle.
	FieldPowertrainRange = "Powertrain_Range"
	// FieldPowertrainTractionBatteryChargingChargeLimit Target charge limit (state of charge) for battery.
	FieldPowertrainTractionBatteryChargingChargeLimit = "Powertrain_TractionBattery_Charging_ChargeLimit"
	// FieldPowertrainTractionBatteryChargingIsCharging True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
	FieldPowertrainTractionBatteryChargingIsCharging = "Powertrain_TractionBattery_Charging_IsCharging"
	// FieldPowertrainTractionBatteryGrossCapacity Gross capacity of the battery.
	FieldPowertrainTractionBatteryGrossCapacity = "Powertrain_TractionBattery_GrossCapacity"
	// FieldPowertrainTractionBatteryStateOfChargeCurrent Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
	FieldPowertrainTractionBatteryStateOfChargeCurrent = "Powertrain_TractionBattery_StateOfCharge_Current"
	// FieldPowertrainTransmissionTravelledDistance Odometer reading, total distance travelled during the lifetime of the transmission.
	FieldPowertrainTransmissionTravelledDistance = "Powertrain_Transmission_TravelledDistance"
	// FieldPowertrainType Defines the powertrain type of the vehicle.
	FieldPowertrainType = "Powertrain_Type"
	// FieldSpeed Vehicle speed.
	FieldSpeed = "Speed"
	// FieldVehicleIdentificationBrand Vehicle brand or manufacturer.
	FieldVehicleIdentificationBrand = "VehicleIdentification_Brand"
	// FieldVehicleIdentificationModel Vehicle model.
	FieldVehicleIdentificationModel = "VehicleIdentification_Model"
	// FieldVehicleIdentificationYear Model year of the vehicle.
	FieldVehicleIdentificationYear = "VehicleIdentification_Year"
)

// JSONName2CHName maps the JSON field names to the Clickhouse column names.
var JSONName2CHName = map[string]string{
	"chassisAxleRow1WheelLeftTirePressure":          "Chassis_Axle_Row1_Wheel_Left_Tire_Pressure",
	"chassisAxleRow1WheelRightTirePressure":         "Chassis_Axle_Row1_Wheel_Right_Tire_Pressure",
	"chassisAxleRow2WheelLeftTirePressure":          "Chassis_Axle_Row2_Wheel_Left_Tire_Pressure",
	"chassisAxleRow2WheelRightTirePressure":         "Chassis_Axle_Row2_Wheel_Right_Tire_Pressure",
	"currentLocationAltitude":                       "CurrentLocation_Altitude",
	"currentLocationLatitude":                       "CurrentLocation_Latitude",
	"currentLocationLongitude":                      "CurrentLocation_Longitude",
	"currentLocationTimestamp":                      "CurrentLocation_Timestamp",
	"dIMOAftermarketHDOP":                           "DIMO_Aftermarket_HDOP",
	"dIMOAftermarketNSAT":                           "DIMO_Aftermarket_NSAT",
	"dIMOAftermarketSSID":                           "DIMO_Aftermarket_SSID",
	"dIMOAftermarketWPAState":                       "DIMO_Aftermarket_WPAState",
	"exteriorAirTemperature":                        "Exterior_AirTemperature",
	"lowVoltageBatteryCurrentVoltage":               "LowVoltageBattery_CurrentVoltage",
	"oBDBarometricPressure":                         "OBD_BarometricPressure",
	"oBDEngineLoad":                                 "OBD_EngineLoad",
	"oBDIntakeTemp":                                 "OBD_IntakeTemp",
	"oBDRunTime":                                    "OBD_RunTime",
	"powertrainCombustionEngineECT":                 "Powertrain_CombustionEngine_ECT",
	"powertrainCombustionEngineEngineOilLevel":      "Powertrain_CombustionEngine_EngineOilLevel",
	"powertrainCombustionEngineMAF":                 "Powertrain_CombustionEngine_MAF",
	"powertrainCombustionEngineSpeed":               "Powertrain_CombustionEngine_Speed",
	"powertrainCombustionEngineTPS":                 "Powertrain_CombustionEngine_TPS",
	"powertrainFuelSystemAbsoluteLevel":             "Powertrain_FuelSystem_AbsoluteLevel",
	"powertrainFuelSystemSupportedFuelTypes":        "Powertrain_FuelSystem_SupportedFuelTypes",
	"powertrainRange":                               "Powertrain_Range",
	"powertrainTractionBatteryChargingChargeLimit":  "Powertrain_TractionBattery_Charging_ChargeLimit",
	"powertrainTractionBatteryChargingIsCharging":   "Powertrain_TractionBattery_Charging_IsCharging",
	"powertrainTractionBatteryGrossCapacity":        "Powertrain_TractionBattery_GrossCapacity",
	"powertrainTractionBatteryStateOfChargeCurrent": "Powertrain_TractionBattery_StateOfCharge_Current",
	"powertrainTransmissionTravelledDistance":       "Powertrain_Transmission_TravelledDistance",
	"powertrainType":                                "Powertrain_Type",
	"speed":                                         "Speed",
	"vehicleIdentificationBrand":                    "VehicleIdentification_Brand",
	"vehicleIdentificationModel":                    "VehicleIdentification_Model",
	"vehicleIdentificationYear":                     "VehicleIdentification_Year",
}
