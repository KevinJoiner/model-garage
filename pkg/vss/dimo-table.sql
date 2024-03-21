CREATE TABLE IF NOT EXISTS dimo (
	DefinitionID String COMMENT 'ID for the vehicles definition',
	Source String COMMENT 'where the data was sourced from',
	Subject String COMMENT 'subjet of this vehicle data',
	Timestamp DateTime('UTC') COMMENT 'timestamp of when this data was colllected',
	Type String COMMENT 'type of data collected',
	Vehicle_CurrentLocation_Altitude Float64 COMMENT 'Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Latitude Float64 COMMENT 'Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Longitude Float64 COMMENT 'Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.',
	Vehicle_CurrentLocation_Timestamp DateTime('UTC') COMMENT 'Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.',
	Vehicle_LowVoltageBattery_CurrentVoltage Float32 COMMENT 'Current Voltage of the low voltage battery.',
	Vehicle_Speed Float32 COMMENT 'Vehicle speed.',
	Vehicle_VehicleIdentification_Brand String COMMENT 'Vehicle brand or manufacturer.',
	Vehicle_VehicleIdentification_Model String COMMENT 'Vehicle model.',
	Vehicle_VehicleIdentification_VIN String COMMENT '17-character Vehicle Identification Number (VIN) as defined by ISO 3779.',
	Vehicle_VehicleIdentification_Year UInt16 COMMENT 'Model year of the vehicle.',
	VehicleID String COMMENT 'unque DIMO ID for the vehicle',
)
ENGINE = MergeTree()
ORDER BY (Subject, Timestamp)