package convert_test

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/DIMO-Network/model-garage/pkg/vss/convert"
	"github.com/stretchr/testify/require"
)

const tokenID = uint32(123)

func TestFullFromV2DataConversion(t *testing.T) {
	t.Parallel()
	actualSignals, err := convert.SignalsFromPayload(nil, nil, []byte(fullV2InputJSON)) //nolint:staticcheck // we want this to fail not v2
	require.NoErrorf(t, err, "error converting full input data: %v", err)
	require.Equalf(t, expectedV2Signals, actualSignals, "converted vehicle does not match expected vehicle")
}

var fullV2InputJSON = `{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "2.0",
    "subject": "0x98D78d711C0ec544F6fb5d54fcf6559CF41546a9",
    "time": "2024-04-18T17:20:46.436008782Z",
    "type": "com.dimo.device.status",
    "signature": "0x72208df3282c890ec72afe22abbcffb76ec73dc3e1ce8becd158469126f10c35245289e02ad41782e07376f9b9092a2fec96477a6e453fed1ca3860639e776f31b",
    "data": {
        "timestamp": 1713460846435,
        "device": {
            "rpiUptimeSecs": 218,
            "batteryVoltage": 12.28
        },
        "vehicle": {
            "signals": [
                {
                    "timestamp": 1713460823243,
                    "name": "longFuelTrim",
                    "value": 25
                },
                {
                    "timestamp": 1713460826633,
                    "name": "coolantTemp",
                    "value": 107
                },
                {
                    "timestamp": 1713460827173,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460829314,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460829844,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460830382,
                    "name": "shortFuelTrim",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460837235,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460842256,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460844422,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460844962,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460845497,
                    "name": "shortFuelTrim",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460846435,
                    "name": "longitude",
                    "value": -56.50151833333334
                },
                {
                    "timestamp": 1713460846435,
                    "name": "latitude",
                    "value": 56.27014
                },
                {
                    "timestamp": 1713460846435,
                    "name": "hdop",
                    "value": 1.4
                },
                {
                    "timestamp": 1713460846435,
                    "name": "nsat",
                    "value": 6
                },
                {
                    "timestamp": 1713460846435,
                    "name": "wpa_state",
                    "value": "COMPLETED"
                },
                {
                    "timestamp": 1713460846435,
                    "name": "ssid",
                    "value": "foo"
                }
            ]
        }
    },
    "vehicleTokenId": 123,
    "make": "",
    "model": "",
    "year": 0
}`

var expectedV2Signals = []vss.Signal{
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 26, 633000000, time.UTC), Name: "powertrainCombustionEngineECT", ValueNumber: 107, ValueString: "", Source: "dimo/integration/123"}, //nolint // false positive
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 27, 173000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 314000000, time.UTC), Name: "oBDEngineLoad", ValueNumber: 12.54912, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 29, 844000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 37, 235000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 42, 256000000, time.UTC), Name: "powertrainCombustionEngineMAF", ValueNumber: 475.79, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 422000000, time.UTC), Name: "oBDEngineLoad", ValueNumber: 12.54912, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 44, 962000000, time.UTC), Name: "powertrainCombustionEngineTPS", ValueNumber: 23.529600000000002, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLongitude", ValueNumber: -56.50151833333334, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "currentLocationLatitude", ValueNumber: 56.27014, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dIMOAftermarketHDOP", ValueNumber: 1.4, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dIMOAftermarketNSAT", ValueNumber: 6, ValueString: "", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dIMOAftermarketWPAState", ValueNumber: 0, ValueString: "COMPLETED", Source: "dimo/integration/123"},
	{TokenID: tokenID, Timestamp: time.Date(2024, time.April, 18, 17, 20, 46, 435000000, time.UTC), Name: "dIMOAftermarketSSID", ValueNumber: 0, ValueString: "foo", Source: "dimo/integration/123"},
}
