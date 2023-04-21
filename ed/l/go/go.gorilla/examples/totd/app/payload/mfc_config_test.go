package payload

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_MFCConfig_GetErrorRampStringID(ts *testing.T) {
	m := MFCConfig{IDGen: "case1-%d"}

	Convey("case with blank value", ts, func() {
		actual := m.GetErrorRampStringID()
		expected := "case1-0"
		So(actual, ShouldResemble, expected)
	})

	Convey("case with actual value", ts, func() {
		m.ErrorRamp = 99
		actual := m.GetErrorRampStringID()
		expected := "case1-99"
		So(actual, ShouldResemble, expected)
	})
}

func Test_MFCConfig_ConvertLaneIdxToID(ts *testing.T) {
	Convey("ConvertLaneIdxToID should return valid lane ID", ts, func() {
		stubbedConfigs := map[string]MFCConfig{
			// Regular configs.
			"c1_m1": {IDGen: "%v", ErrorRamp: 4, Start: 1, Count: 3},
			"c1_m2": {IDGen: "%v", ErrorRamp: 4, Start: 3, Count: 3},
			"c1_m3": {IDGen: "%v", ErrorRamp: 4, LaneMapping: map[int64]int64{1: 2, 2: 3, 3: 1}},
			// Express configs.
			"c2_m1": {
				IDGen: "%v", ErrorRamp: 4, Start: 1, Count: 3,
				ExpressLaneMapping: map[int64]int64{1: 5, 2: 6, 3: 7},
			},
			"c2_m2": {
				IDGen: "%v", ErrorRamp: 4, Start: 1, Count: 3,
				ExpressLaneMapping: map[int64]int64{1: 6, 2: 7, 3: 5},
			},
		}

		type testCase struct {
			name      string
			clientID  string
			mfcID     string
			isExpress bool
			idx       int64
			expected  string
		}
		testCases := []testCase{
			// Test cases for regular configs.
			// Order A-Z.
			{clientID: "c1", mfcID: "m1", isExpress: false, idx: 1, expected: "1"},
			{clientID: "c1", mfcID: "m1", isExpress: false, idx: 2, expected: "2"},
			{clientID: "c1", mfcID: "m1", isExpress: false, idx: 3, expected: "3"},
			// Order Z-A.
			{clientID: "c1", mfcID: "m2", isExpress: false, idx: 1, expected: "3"},
			{clientID: "c1", mfcID: "m2", isExpress: false, idx: 2, expected: "2"},
			{clientID: "c1", mfcID: "m2", isExpress: false, idx: 3, expected: "1"},
			// Order determined by mapping.
			{clientID: "c1", mfcID: "m3", isExpress: false, idx: 1, expected: "2"},
			{clientID: "c1", mfcID: "m3", isExpress: false, idx: 2, expected: "3"},
			{clientID: "c1", mfcID: "m3", isExpress: false, idx: 3, expected: "1"},

			// Test cases for express configs.
			// Order Z-A.
			{clientID: "c2", mfcID: "m1", isExpress: true, idx: 1, expected: "5"},
			{clientID: "c2", mfcID: "m1", isExpress: true, idx: 2, expected: "6"},
			{clientID: "c2", mfcID: "m1", isExpress: true, idx: 3, expected: "7"},
			// Order determined by mapping.
			{clientID: "c2", mfcID: "m2", isExpress: true, idx: 1, expected: "6"},
			{clientID: "c2", mfcID: "m2", isExpress: true, idx: 2, expected: "7"},
			{clientID: "c2", mfcID: "m2", isExpress: true, idx: 3, expected: "5"},
		}
		for i, tc := range testCases {
			Convey(fmt.Sprintf("ConvertLaneIdxToID test case %d: %s", i, tc.name), func(c C) {
				key := tc.clientID + "_" + tc.mfcID
				conf := stubbedConfigs[key]
				actual, err := conf.ConvertLaneIdxToID(tc.idx, tc.isExpress)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, tc.expected)
			})
		}
	})
}
