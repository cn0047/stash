package spec

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/test"
)

func getCtx(clientID, env, mfcID string) context.Context {
	ctx := context.Background()

	ctx = context.WithValue(ctx, payload.ContextKeyToken, "tkn")
	ctx = context.WithValue(ctx, payload.ContextKeyRetailer, clientID)
	ctx = context.WithValue(ctx, payload.ContextKeyEnv, env)
	ctx = context.WithValue(ctx, payload.ContextKeyMfc, mfcID)

	return ctx
}

func TestSpecForToteAssignmentCreate(t *testing.T) {
	Convey("Given: clientWithCount2AndDepth2AndStart1", t, func() {
		clientId := "clientWithCount2AndDepth2AndStart1"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			OrderID:  "o1",
			ToteID:   "tt1",
			DryRun:   true,
		}

		Convey("When: to MFCWithEmptyRamp add tt1 for o1", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWithEmptyRamp"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith1Tote"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		Convey("When: to MFCWith1Tote t1 for o1 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith1Tote"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith2TotesFor2Orders"
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o2-tt2", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith2TotesFor2Orders"
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 | o3-t3 |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o1-tt1", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 | o3-t3 |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o2-tt2", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 | o3-t3 |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o3-tt3", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o3"
				input.ToteID = "tt3"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})
	})

	Convey("Given: clientWithCount2AndDepth2AndStart2", t, func() {
		clientId := "clientWithCount2AndDepth2AndStart2"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			OrderID:  "o1",
			ToteID:   "tt1",
			DryRun:   true,
		}

		Convey("When: to MFCWithEmptyRamp add tt1 for o1", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWithEmptyRamp"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith1Tote"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith1Tote t1 for o1 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith1Tote"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o2-t2 |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.MfcID = "MFCWith2TotesFor2Orders"
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o2-t2 |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o2-tt2", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith2TotesFor2Orders"
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o2-t2 |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 | o3-t3 |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o1-tt1", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o2-t2 |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 | o3-t3 |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o2-tt2", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o2-t2 |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 | o3-t3 |
		// +--------+---------------+
		Convey("When: to MFCWith3TotesFor3Orders o1-t1, o2-t2, o3-t3 add o3-tt3", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.MfcID = "MFCWith3TotesFor3Orders"
				input.OrderID = "o3"
				input.ToteID = "tt3"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})
	})

	Convey("Given: clientWithCount3AndDepth2AndStart1", t, func() {
		clientId := "clientWithCount3AndDepth2AndStart1"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWith2TotesFor2Orders",
			DryRun:   true,
		}

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 |       |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o1-tt1", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 |       |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o2-tt2", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 |       |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o3-tt3", func() {
			Convey("Then: tote goes to lane 3", func() {
				input.OrderID = "o3"
				input.ToteID = "tt3"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH03")
			})
		})
	})

	Convey("Given: clientWithCount3AndDepth2AndStart3", t, func() {
		clientId := "clientWithCount3AndDepth2AndStart3"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWith2TotesFor2Orders",
			DryRun:   true,
		}

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 3", func() {
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH03")
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add o2-tt2", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2Orders"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o2-t2 |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2Orders o1-t1, o2-t2 add tt3 for o3", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.OrderID = "o3"
				input.ToteID = "tt3"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})
	})

	Convey("Given: clientWithCount3AndDepth2AndStart1", t, func() {
		clientId := "clientWithCount3AndDepth2AndStart1"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWith2TotesFor2OrdersWitGap",
			DryRun:   true,
		}

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 1", func() {
				// @TODO
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt2 for o2", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o1-t1 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt3 for o3", func() {
			Convey("Then: tote goes to lane 3", func() {
				// @TODO
			})
		})
	})

	Convey("Given: clientWithCount3AndDepth2AndStart3", t, func() {
		clientId := "clientWithCount3AndDepth2AndStart3"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWith2TotesFor2OrdersWitGap",
			DryRun:   true,
		}
		_ = input

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o3-t3 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 3", func() {
				// @TODO
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o3-t3 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt2 for o2", func() {
			Convey("Then: tote goes to lane 2", func() {
				// @TODO
			})
		})

		// Given: "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 | o3-t3 |       |
		// +--------+---------------+
		// | lane 2 |       |       |
		// +--------+---------------+
		// | lane 3 | o1-t1 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt3 for o3", func() {
			Convey("Then: tote goes to lane 3", func() {
				// @TODO
			})
		})
	})

	Convey("Given: clientWithCount3AndDepth2AndLaneMapping", t, func() {
		clientId := "clientWithCount3AndDepth2AndLaneMapping"
		input := payload.CreateToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWith2TotesFor2OrdersWitGap",
			DryRun:   true,
		}

		// Given: "clientWithCount3AndDepth2AndLaneMapping", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt1 for o1", func() {
			Convey("Then: tote goes to lane 2", func() {
				input.OrderID = "o1"
				input.ToteID = "tt1"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH02")
			})
		})

		// Given: "clientWithCount3AndDepth2AndLaneMapping", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt2 for o2", func() {
			Convey("Then: tote goes to lane 1", func() {
				input.OrderID = "o2"
				input.ToteID = "tt2"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH01")
			})
		})

		// Given: "clientWithCount3AndDepth2AndLaneMapping", "MFCWith2TotesFor2OrdersWitGap"
		// +--------+---------------+
		// |        | Totes         |
		// +--------+---------------+
		// | lane 1 |       |       |
		// +--------+---------------+
		// | lane 2 | o1-t1 |       |
		// +--------+---------------+
		// | lane 3 | o3-t3 |       |
		// +--------+---------------+
		Convey("When: to MFCWith2TotesFor2OrdersWitGap o1-t1, o3-t3 add tt3 for o3", func() {
			Convey("Then: tote goes to lane 3", func() {
				input.OrderID = "o3"
				input.ToteID = "tt3"
				out, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), input)

				So(err, ShouldBeNil)
				So(out.ToteAssignment.LaneID, ShouldEqual, "DISPATCH03")
			})
		})
	})
}

func TestSpecForToteAssignmentDelete(t *testing.T) {
	Convey("Given: clientWithTotesToDelete", t, func() {
		clientId := "clientWithTotesToDelete"
		input := payload.DeleteToteAssignmentInput{
			ClientID: clientId,
			MfcID:    "MFCWithTotesToDelete",
		}
		listInput := payload.ListToteAssignmentsInput{
			ClientID: input.ClientID,
			MfcID:    input.MfcID,
		}
		createInput := payload.CreateToteAssignmentInput{
			ClientID: input.ClientID,
			MfcID:    input.MfcID,
		}

		// Given: "clientWithTotesToDelete", "MFCWithTotesToDelete"
		// +--------+---------------------+
		// |        | Totes               |
		// +--------+---------------------+
		// | lane 1 | o1-t1 |      |      |
		// +--------+---------------------+
		// | lane 2 | o2-t1 |      |      |
		// +--------+---------------------+
		// | lane 3 | o3-t1 |      |      |
		// +--------+---------------------+
		Convey("When: delete 2 totes from MFCWithTotesToDelete", func() {
			// Init testcase, so DB will be consistent even if tests are running in parallel.
			{
				// Insert tote to delete 1.
				createInput.OrderID = "o1"
				createInput.ToteID = "tote-to-delete-1"
				_, err := test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), createInput)
				So(err, ShouldBeNil)
				// Insert tote to delete 2.
				createInput.OrderID = "o2"
				createInput.ToteID = "tote-to-delete-2"
				_, err = test.GetApp(t).ToteAssignment.Create(getCtx(clientId, "", input.MfcID), createInput)
				So(err, ShouldBeNil)
			}

			// Given: "clientWithTotesToDelete", "MFCWithTotesToDelete"
			// +--------+--------------------------------+
			// |        | Totes                          |
			// +--------+--------------------------------+
			// | lane 1 | o1t1 | tote-to-delete-1 |      |
			// +--------+--------------------------------+
			// | lane 2 | o2t1 | tote-to-delete-1 |      |
			// +--------+--------------------------------+
			// | lane 3 | o3t1 |                  |      |
			// +--------+--------------------------------+
			Convey("Then: 2 totes will be deleted", func() {
				input.ToteIDs = []string{"tote-to-delete-1", "tote-to-delete-2"}
				_, err := test.GetApp(t).ToteAssignment.Delete(getCtx(clientId, "", input.MfcID), input)
				So(err, ShouldBeNil)

				res, err := test.GetApp(t).ToteAssignment.List(getCtx(clientId, "", input.MfcID), listInput)
				So(err, ShouldBeNil)
				r, ok := res.(payload.ListToteAssignmentsGroupedByOrderOutput)
				So(ok, ShouldBeTrue)
				for _, actualToteLists := range r.Orders {
					for _, toteIDWhichShouldBeDeleted := range input.ToteIDs {
						So(actualToteLists, ShouldNotContain, toteIDWhichShouldBeDeleted)
					}
				}
			})
		})

		Convey("When: delete non existing totes from MFCWithTotesToDelete", func() {
			Convey("Then: nothing will be deleted", func() {
				resBefore, err := test.GetApp(t).ToteAssignment.List(getCtx(clientId, "", input.MfcID), listInput)
				So(err, ShouldBeNil)
				expectedData, ok := resBefore.(payload.ListToteAssignmentsGroupedByOrderOutput)
				So(ok, ShouldBeTrue)

				input.ToteIDs = []string{"none-existing-tote-1"}
				_, err = test.GetApp(t).ToteAssignment.Delete(getCtx(clientId, "", input.MfcID), input)
				So(err, ShouldBeNil)

				resAfter, err := test.GetApp(t).ToteAssignment.List(getCtx(clientId, "", input.MfcID), listInput)
				So(err, ShouldBeNil)
				actualData, ok := resAfter.(payload.ListToteAssignmentsGroupedByOrderOutput)
				So(ok, ShouldBeTrue)

				So(actualData, ShouldResemble, expectedData)
			})
		})
	})
}
