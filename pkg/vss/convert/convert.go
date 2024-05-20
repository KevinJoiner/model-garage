package convert

import (
	"context"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/tidwall/gjson"
)

const (
	// StatusV1 is the version string for payloads with the version 1.0 schema.
	StatusV1 = "1.0"
	// StatusV1Converted is the version string for payloads that have been converted to the 1.0 schema.
	StatusV1Converted = "1.1"
	// StatusV2 is the version string for payloads with the version 2.0 schema.
	StatusV2 = "2.0"
)

// SignalsFromPayload extracts signals from a payload.
// It detects the payload version and calls the appropriate function.
func SignalsFromPayload(ctx context.Context, tokenGetter TokenIDGetter, jsonData []byte) ([]vss.Signal, error) {
	version := GetSchemaVersion(jsonData)
	switch {
	case version == StatusV1 || version == StatusV1Converted:
		return SignalsFromV1Payload(ctx, tokenGetter, jsonData)
	case version == StatusV2:
		return SignalsFromV2Payload(jsonData)
	default:
		return nil, VersionError{Version: version}
	}
}

// GetSchemaVersion returns the version string of the schema used in the payload.
func GetSchemaVersion(jsonData []byte) string {
	dataSchema := gjson.GetBytes(jsonData, "dataschema")
	if dataSchema.Exists() {
		// get version string at the end of the URI
		// Ex. dimo.zone.status/v1.1
		schemaString := dataSchema.String()
		version := schemaString[strings.LastIndex(schemaString, "/")+1:]
		return strings.TrimPrefix(version, "v")
	}
	return gjson.GetBytes(jsonData, "specversion").String()
}
