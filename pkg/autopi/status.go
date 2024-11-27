// Package status holds decoding functions for Ruptela status payloads.
package autopi

import (
	"errors"
	"fmt"
	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/tidwall/gjson"
	"golang.org/x/mod/semver"
)

const (
	// StatusV1 is the version string for payloads with the version 1.0 schema.
	StatusV1 = "v1.0.0"
	// StatusV1Converted is the version string for payloads that have been converted to the 1.0 schema.
	StatusV1Converted = "v1.1.0"
	// StatusV2 is the version string for payloads with the version 2.0 schema.
	StatusV2 = "v2.0.0"
)

// GetDataVersion returns the version string used in the payload.
func GetDataVersion(jsonData []byte) string {
	dataVersion := gjson.GetBytes(jsonData, "dataversion")
	if !dataVersion.Exists() {
		return ""
	}
	return dataVersion.String()
}

// hasV1Data checks if the payload has the same sceham as a v1.0.0.
func HasV1Data(version string) bool {
	return version == "" || semver.Compare(StatusV1, version) == 0 || semver.Compare(StatusV1Converted, version) == 0
}

// TokenIDFromData gets a tokenID from a V2 payload.
func TokenIDFromData(jsonData []byte) (uint32, error) {
	lookupKey := "subject"
	subject := gjson.GetBytes(jsonData, lookupKey)
	if !subject.Exists() {
		return 0, convert.FieldNotFoundError{Field: "tokenID", Lookup: lookupKey}
	}
	subjectStr, ok := subject.Value().(string)
	if !ok {
		return 0, fmt.Errorf("%s field is not a string", lookupKey)
	}
	subjectDID, err := cloudevent.DecodeNFTDID(subjectStr)
	if err != nil {
		return 0, fmt.Errorf("error decoding subject: %w", err)
	}
	return subjectDID.TokenID, nil
}

// SourceFromData gets a source from a V2 payload.
func SourceFromData(jsonData []byte) (string, error) {
	lookupKey := "source"
	source := gjson.GetBytes(jsonData, lookupKey)
	if !source.Exists() {
		return "", convert.FieldNotFoundError{Field: "source", Lookup: lookupKey}
	}
	src, ok := source.Value().(string)
	if !ok {
		return "", errors.New("source field is not a string")
	}
	return src, nil
}
