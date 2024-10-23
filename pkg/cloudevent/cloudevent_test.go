package cloudevent_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/cloudevent"
)

func TestCloudEventHeaderJSON(t *testing.T) {
	now := time.Now().UTC()

	// Test case with extra fields
	jsonData := []byte(`{
		"id": "123",
		"source": "test-source",
		"producer": "test-producer",
		"specversion": "1.0",
		"subject": "test-subject",
		"time": "` + now.Format(time.RFC3339Nano) + `",
		"type": "test.event",
		"datacontenttype": "application/json",
		"extra_field1": "value1",
		"extra_field2": 42,
		"extra_nested": {"key": "value"}
	}`)

	var event cloudevent.CloudEventHeader
	if err := json.Unmarshal(jsonData, &event); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify known fields
	if event.ID != "123" {
		t.Errorf("Expected ID=123, got %s", event.ID)
	}
	if event.Source != "test-source" {
		t.Errorf("Expected Source=test-source, got %s", event.Source)
	}

	// Verify extra fields
	if val, ok := event.Extras["extra_field1"].(string); !ok || val != "value1" {
		t.Errorf("Expected extra_field1=value1, got %v", event.Extras["extra_field1"])
	}
	if val, ok := event.Extras["extra_field2"].(float64); !ok || val != 42 {
		t.Errorf("Expected extra_field2=42, got %v", event.Extras["extra_field2"])
	}
	if val, ok := event.Extras["extra_nested"].(map[string]any); !ok {
		t.Errorf("Expected extra_nested to be a map, got %T", event.Extras["extra_nested"])
		if nestedVal, ok := val["key"].(string); !ok || nestedVal != "value" {
			t.Errorf("Expected extra_nested[key]=value, got %v", val["key"])
		}
		t.Errorf("Expected extra_nested[key]=value, got %v", event.Extras["extra_nested"])
	}

	// Test marshaling back to JSON
	output, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal both original and output JSON into maps for comparison
	originalMap := map[string]any{}
	outputMap := map[string]any{}
	if err := json.Unmarshal(jsonData, &originalMap); err != nil {
		t.Fatalf("Failed to unmarshal original JSON: %v", err)
	}
	if err := json.Unmarshal(output, &outputMap); err != nil {
		t.Fatalf("Failed to unmarshal output JSON: %v", err)
	}

	// Compare maps
	for k, v := range originalMap {
		outputVal, exists := outputMap[k]
		if !exists {
			t.Errorf("Field %s missing from output", k)
		}
		if reflect.TypeOf(v).Comparable() && v != outputVal { // Skip exact time comparison
			t.Errorf("Field %s value mismatch: expected %v, got %v", k, v, outputVal)
		}
	}
}
