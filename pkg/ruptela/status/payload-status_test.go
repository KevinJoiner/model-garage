package status_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/ruptela/status"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestFullFromDataConversion(t *testing.T) {
	t.Parallel()
	actualSignals, err := status.SignalsFromV1Payload([]byte(fullInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.SignalValue) int {
		return cmp.Compare(a.Name, b.Name)
	}
	for _, s := range actualSignals {
		fmt.Printf("%+v\n", s)
	}
	slices.SortFunc(expectedSignals, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)
	require.Equal(t, expectedSignals, actualSignals, "converted vehicle does not match expected vehicle")
}

var (
	fullInputJSON = `
{
	"source": "ruptela/TODO",
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
			"645": "8",
			"135": "0",
			"136": "0",
			"137": "14",
			"173": "1",
			"205": "5",
			"207": "5",
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
			"723": "EA42",
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
			"99": "1",
			"999": "0"
		},
		"trigger": 7
	},
	"ds": "r/v0/s",
	"signature": "0x6fb5849e21e66f3e0619f148bc032153aa4c90be4cd175e83c1f959e1bc551d940d516fe74f50aed380e432406675c583e75155bf1c77b9ec0761b1dbe1ab87e1c",
	"subject": "did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
	"time": "2024-09-27T08:33:26Z"
  }
}`
	ts = time.Date(2024, 9, 27, 8, 33, 26, 0, time.UTC)

	expectedSignals = []vss.SignalValue{
		{Timestamp: ts, Name: vss.FieldLowVoltageBatteryCurrentVoltage, ValueNumber: 14.335},
		{Timestamp: ts, Name: vss.FieldPowertrainFuelSystemAbsoluteLevel, ValueNumber: 5},
		{Timestamp: ts, Name: vss.FieldCurrentLocationAltitude, ValueNumber: 104.8},
		{Timestamp: ts, Name: vss.FieldDIMOAftermarketHDOP, ValueNumber: 6},
		{Timestamp: ts, Name: vss.FieldCurrentLocationLatitude, ValueNumber: 52.2721466},
		{Timestamp: ts, Name: vss.FieldCurrentLocationLongitude, ValueNumber: -0.9014316},
		{Timestamp: ts, Name: vss.FieldDIMOAftermarketNSAT, ValueNumber: 20},
		{Timestamp: ts, Name: vss.FieldPowertrainType, ValueString: "COMBUSTION"},
		{Timestamp: ts, Name: vss.FieldPowertrainFuelSystemRelativeLevel, ValueNumber: 2},
		{Timestamp: ts, Name: vss.FieldOBDDistanceWithMIL, ValueNumber: 0},
		{Timestamp: ts, Name: vss.FieldPowertrainCombustionEngineTPS, ValueNumber: 0},
		{Timestamp: ts, Name: vss.FieldPowertrainTransmissionTravelledDistance, ValueNumber: 8},
		{Timestamp: ts, Name: vss.FieldSpeed, ValueNumber: 0},
		{Timestamp: ts, Name: vss.FieldPowertrainTractionBatteryRange, ValueNumber: 59.97},
	}
)
