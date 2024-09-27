package ruptela_test

import (
	"cmp"
	"context"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

type tokenGetter struct{}

func (*tokenGetter) TokenIDFromSubject(context.Context, string) (uint32, error) {
	return 123, nil
}

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	getter := &tokenGetter{}
	actualSignals, err := ruptela.SignalsFromV1Payload(context.Background(), getter, []byte(fullInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.Signal) int {
		return cmp.Compare(a.Name, b.Name)
	}
	slices.SortFunc(expectedSignals, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)

	require.Equal(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var (
	fullInputJSON = `
{
	"data": {
		"pos": {
			"alt": 1048,
			"dir": 19730,
			"hdop": 6,
			"lat": 522721466,
			"lon": -9014316,
			"sat": 20,
			"spd": 0
		},
		"prt": 0,
		"signals": {
			"102": "0",
			"103": "0",
			"104": "53414C4C41414146",
			"105": "3341413534343438",
			"106": "3200000000000000",
			"107": "0",
			"108": "0",
			"114": "0",
			"135": "0",
			"136": "0",
			"137": "14",
			"173": "1",
			"205": "0",
			"207": "0",
			"29": "37FF",
			"30": "1080",
			"409": "1",
			"49": "FE",
			"50": "FA",
			"5005": "31",
			"5060": "6597",
			"51": "ED",
			"525": "A502A",
			"525_1": "A502A",
			"642": "FFFF",
			"645": "FFFFFFFF",
			"722": "FF",
			"723": "FFFF",
			"754": "FB8F",
			"92": "0",
			"93": "0",
			"94": "0",
			"95": "0",
			"950": "0",
			"96": "FF",
			"97": "FF",
			"98": "0",
			"985": "0",
			"99": "0",
			"999": "0"
		},
		"trigger": 7
	},
	"ds": "r/v0/s",
	"signature": "0x6fb5849e21e66f3e0619f148bc032153aa4c90be4cd175e83c1f959e1bc551d940d516fe74f50aed380e432406675c583e75155bf1c77b9ec0761b1dbe1ab87e1c",
	"subject": "0xf47f6579029a1c53388e4d8776413aa3f993cb94",
	"time": "2024-09-27T08:33:26Z",
	"topic": "devices/0xf47f6579029a1c53388e4d8776413aa3f993cb94/status",
	"vehicleTokenId": 0
}`
	ts = time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC)

	expectedSignals = []vss.Signal{
		{TokenID: 0, Timestamp: ts, Name: vss.FieldCurrentLocationLatitude, ValueNumber: 522721466, Source: "ruptela/TODO"},
		{TokenID: 0, Timestamp: ts, Name: vss.FieldCurrentLocationLongitude, ValueNumber: -9014316, Source: "ruptela/TODO"},
	}
)
