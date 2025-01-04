package status_test

import (
	"cmp"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/ruptela/status"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestLocationPayload(t *testing.T) {
	t.Parallel()
	actualSignals, err := status.SignalsFromLocationPayload([]byte(locationInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.SignalValue) int {
		return cmp.Compare(a.Name, b.Name)
	}
	expected := expectedLocationSignals()
	slices.SortFunc(expected, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)
	require.Equal(t, expected, actualSignals, "converted vehicle does not match expected vehicle")
}

func expectedLocationSignals() []vss.SignalValue {
	ts = time.Unix(1727360340, 0).UTC()
	return []vss.SignalValue{
		{Timestamp: ts, Name: vss.FieldCurrentLocationAltitude, ValueNumber: 123.2},
		{Timestamp: ts, Name: vss.FieldCurrentLocationLatitude, ValueNumber: 43.2699983},
		{Timestamp: ts, Name: vss.FieldCurrentLocationLongitude, ValueNumber: -71.50142},
		{Timestamp: ts, Name: vss.FieldDIMOAftermarketHDOP, ValueNumber: 0},
		{Timestamp: ts.Add(time.Second), Name: vss.FieldCurrentLocationAltitude, ValueNumber: 1.2},
		{Timestamp: ts.Add(time.Second), Name: vss.FieldCurrentLocationLatitude, ValueNumber: 44.2699983},
		{Timestamp: ts.Add(time.Second), Name: vss.FieldCurrentLocationLongitude, ValueNumber: -71.50142},
		{Timestamp: ts.Add(time.Second), Name: vss.FieldDIMOAftermarketHDOP, ValueNumber: 1},
		{Timestamp: ts.Add(time.Second * 2), Name: vss.FieldCurrentLocationAltitude, ValueNumber: 0.2},
		{Timestamp: ts.Add(time.Second * 2), Name: vss.FieldCurrentLocationLatitude, ValueNumber: 45.2699983},
		{Timestamp: ts.Add(time.Second * 2), Name: vss.FieldCurrentLocationLongitude, ValueNumber: -71.50142},
		{Timestamp: ts.Add(time.Second * 2), Name: vss.FieldDIMOAftermarketHDOP, ValueNumber: 2},
	}
}

var locationInputJSON = `
	{
		"subject": "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
		"source": "ruptela/TODO",
		"data": {
			"location": [
				{
					"alt": 1232,
					"dir": 0,
					"hdop": 0,
					"lat":  432699983,
					"lon": -715014200,
					"ts": 1727360340
				},
				{
					"alt": 12,
					"dir": 0,
					"hdop": 1,
					"lat": 442699983,
					"lon": -715014200,
					"ts": 1727360341
				},
				{
					"alt": 2,
					"dir": 0,
					"hdop": 2,
					"lat": 452699983,
					"lon": -715014200,
					"ts": 1727360342
				}
			]
		},
		"ds": "r/v0/loc",
		"signature": "0xb6b130b31b4cd73182008d286fe878bc311a2259b8cfc1ba785495d9c88a028c55e489608191518b3ad26e8226c35a4dfd032f03aac930712ac038e2afeeefc81c",
		"time": "2024-09-26T14:19:14Z"
	}`
