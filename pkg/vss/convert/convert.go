package convert

import (
	"context"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/vss"
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

// SignalsFromPayload extracts signals from a payload.
// It detects the payload version and calls the appropriate function.
func SignalsFromPayload(ctx context.Context, tokenGetter TokenIDGetter, jsonData []byte) ([]vss.Signal, error) {
	version := GetSchemaVersion(jsonData)
	switch {
	case hasV1Schema(version):
		return SignalsFromV1Payload(ctx, tokenGetter, jsonData)
	case semver.Compare(StatusV2, version) == 0:
		return SignalsFromV2Payload(jsonData)
	default:
		return nil, VersionError{Version: version}
	}
}

// GetSchemaVersion returns the version string of the schema used in the payload.
func GetSchemaVersion(jsonData []byte) string {
	dataSchema := gjson.GetBytes(jsonData, "dataschema")
	if !dataSchema.Exists() {
		return ""
	}
	schemaString := dataSchema.String()
	version := schemaString[strings.LastIndex(schemaString, "/")+1:]
	return version
}

// hasV1Schema checks if the payload has the same sceham as a v1.0.0.
func hasV1Schema(version string) bool {
	return version == "" || semver.Compare(StatusV1, version) == 0 || semver.Compare(StatusV1Converted, version) == 0
}
