CREATE TABLE IF NOT EXISTS signal(
	TokenID UInt32 COMMENT 'tokenID of this device data.',
    Subject String COMMENT 'subjet of this vehicle data.',
	Timestamp DateTime('UTC') COMMENT 'timestamp of when this data was colllected.',
	Name LowCardinality(String) COMMENT 'name of the signal collected.',
	ValueString Array(String) COMMENT 'value of the signal collected.',
	ValueNumber Array(Float64) COMMENT 'value of the signal collected.',
)
ENGINE = MergeTree()
ORDER BY (TokenID, Timestamp, Name)