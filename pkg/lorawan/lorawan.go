package main

import (
	"encoding/json"
)

// Data represents the 'data' field in the payload
type Data struct {
	DecodedPayload json.RawMessage `json:"decodedPayload"`
	Device         Device          `json:"device"`
	ID             string          `json:"id"`
	Metadata       Metadata        `json:"metadata"`
	Payload        string          `json:"payload"`
	Timestamp      int64           `json:"timestamp"`
	Via            []Via           `json:"via"`
}

// Device represents the 'device' field in the payload
type Device struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nonce    int    `json:"nonce"`
	Protocol string `json:"protocol"`
}

// Metadata represents the 'metadata' field in the payload
type Metadata struct {
	AppEUI      string `json:"app_eui"`
	DCBalance   int    `json:"dc_balance"`
	DevAddr     string `json:"devAddr"`
	FPort       string `json:"fPort"`
	FCnt        string `json:"fcnt"`
	PayloadSize string `json:"payload_size"`
}

// Via represents the 'via' field in the payload
type Via struct {
	Channel   int        `json:"channel"`
	Frequency float64    `json:"frequency"`
	ID        string     `json:"id"`
	Location  Location   `json:"location"`
	Metadata  GWMetadata `json:"metadata"`
	Network   string     `json:"network"`
	Protocol  string     `json:"protocol"`
	Spreading string     `json:"spreading"`
	Status    string     `json:"status"`
	Timestamp int64      `json:"timestamp"`
}

// Location represents the 'location' field within 'via'
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Ref       string  `json:"ref"`
	RSSI      int     `json:"rssi"`
	SNR       float64 `json:"snr"`
}

// GWMetadata represents the 'metadata' field within 'via'
type GWMetadata struct {
	GatewayID   string `json:"gatewayId"`
	GatewayName string `json:"gatewayName"`
}
