// Package twilio provides structures and methods to interact with Twilio data types.
package twilio

import (
	"time"
)

// ConnectionEvent represents the root structure for the connection event.
type ConnectionEvent struct {
	// EventSID is the SID of the event. This is a copy of the ce_id header field.
	EventSID string `json:"event_sid"`
	// EventType is the type of connection event. This is a copy of the ce_type header field.
	EventType string `json:"event_type"`
	// Timestamp is the UTC timestamp when the event occurred in ISO8601 format.
	Timestamp time.Time `json:"timestamp"`
	// AccountSID is the Account SID of the SuperSIM this record belongs to.
	AccountSID string `json:"account_sid"`
	// APN is the Access Point Name used to establish a data session.
	APN string `json:"apn"`
	// DataModifier indicates if the SuperSIM's data is blocked due to the system such as when the SIM has reached its data limit.
	DataModifier *string `json:"data_modifier,omitempty"`
	// DataSessionSID is the Data Session SID only associated with DataSession events.
	DataSessionSID string `json:"data_session_sid"`
	// DataSessionStartTime is the Data Session start time in UTC and in ISO8601 format.
	DataSessionStartTime *time.Time `json:"data_session_start_time,omitempty"`
	// DataSessionEndTime is the Data Session end time in UTC and in ISO8601 format.
	DataSessionEndTime *time.Time `json:"data_session_end_time,omitempty"`
	// DataSessionUpdateStartTime is the Data Session update start time in UTC and in ISO8601 format.
	DataSessionUpdateStartTime *time.Time `json:"data_session_update_start_time,omitempty"`
	// DataSessionUpdateEndTime is the Data Session update end time in UTC and in ISO8601 format.
	DataSessionUpdateEndTime *time.Time `json:"data_session_update_end_time,omitempty"`
	// DataDownload is the amount of data downloaded to the device in bytes between the data_session_update_start_time and data_session_update_end_time.
	DataDownload *int64 `json:"data_download,omitempty"`
	// DataUpload is the amount of data uploaded from the device in bytes between the data_session_update_start_time and data_session_update_end_time.
	DataUpload *int64 `json:"data_upload,omitempty"`
	// DataTotal is the total amount of data uploaded or downloaded by the device in bytes between the data_session_update_start_time and data_session_update_end_time.
	DataTotal *int64 `json:"data_total,omitempty"`
	// DataSessionDataDownload is the cumulative amount of data downloaded to the device over the data session.
	DataSessionDataDownload *int64 `json:"data_session_data_download,omitempty"`
	// DataSessionDataUpload is the cumulative amount of data uploaded by the device over the data session.
	DataSessionDataUpload *int64 `json:"data_session_data_upload,omitempty"`
	// DataSessionDataTotal is the cumulative amount of data uploaded or downloaded by the device over the data session.
	DataSessionDataTotal *int64 `json:"data_session_data_total,omitempty"`
	// IMEI is the 'international mobile equipment identity' of the device using the SIM to connect. May be null as it is not guaranteed that the visited network will pass on this information.
	IMEI *string `json:"imei,omitempty"`
	// IMSI is the IMSI used by the Super SIM to connect.
	IMSI string `json:"imsi"`
	// IPAddress is the IP address assigned to the device. This IP address is not publicly addressable.
	IPAddress string `json:"ip_address"`
	// SIMICCID is the ICCID of the SuperSIM this record belongs to.
	SIMICCID string `json:"sim_iccid"`
	// SIMSID is the SIM SID of the SuperSIM this record belongs to.
	SIMSID string `json:"sim_sid"`
	// SIMUniqueName is the unique name of the SuperSIM this record belongs to.
	SIMUniqueName string `json:"sim_unique_name"`
	// FleetSID is the SID of the Fleet to which the Super SIM is assigned.
	FleetSID *string `json:"fleet_sid,omitempty"`
	// Location is an object containing information about the location of the cell to which the device was connected. May be null as location information is not guaranteed to be sent by the visited network.
	Location *LocationInfo `json:"location,omitempty"`
	// Network is an object containing information about the network that the Super SIM attempted to connect to or is connected to.
	Network *NetworkInfo `json:"network,omitempty"`
	// RATType is the generation of wireless technology that the device was using.
	RATType string `json:"rat_type"`
	// Error is an object containing information about an error encountered.
	Error *ErrorInfo `json:"error,omitempty"`
}

// LocationInfo represents the location information of the cell tower.
type LocationInfo struct {
	// CellID is the unique ID of the cellular tower that the device was attached to.
	CellID string `json:"cell_id"`
	// LAC is the location area code (LAC) of the cell tower.
	LAC string `json:"lac"`
	// Lat is the spherical coordinate value of the estimated cellular tower location parallel to the Equator.
	Lat float64 `json:"lat"`
	// Lon is the spherical coordinate value of the estimated cellular tower location from the geographical North Pole to the geographical South Pole.
	Lon float64 `json:"lon"`
}

// NetworkInfo represents the information about the visited cellular network.
type NetworkInfo struct {
	// MCC is the mobile country code of the network that the Super SIM attempted to connect to or is connected to.
	MCC string `json:"mcc"`
	// MNC is the mobile network code of the network that the Super SIM attempted to connect to or is connected to.
	MNC string `json:"mnc"`
	// FriendlyName is the human readable name of the Network resource to which the MCC-MNC belongs.
	FriendlyName string `json:"friendly_name"`
	// ISOCountry is the ISO2 code of the Network resource to which the MCC-MNC belongs.
	ISOCountry string `json:"iso_country"`
	// SID is the SID of the Network resource to which the MCC-MNC belongs.
	SID string `json:"sid"`
}

// ErrorInfo represents the error information if any error occurred.
type ErrorInfo struct {
	// Code is the Twilio Error Code.
	Code int `json:"code"`
	// Message is a short message indicating why the error occurred. Could include standardized diameter error messages.
	Message string `json:"message"`
}
