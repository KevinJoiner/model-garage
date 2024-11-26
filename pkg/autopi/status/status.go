// Package status holds decoding functions for Ruptela status payloads.
package status

import (
	"errors"
	"fmt"
	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/convert"
	"github.com/tidwall/gjson"
)

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
