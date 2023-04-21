DELETE FROM configs WHERE 1=1;
DELETE FROM tote_assignments WHERE 1=1;



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "test", "", "test", 0,
    6, 5, 3, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "test", "", "test2", 0,
    4, 3, 3, 1, "DISPATCH%02d", JSON'{}', JSON'{"1":5, "2":6}'
),
(
    "clientWithCount3AndDepth2AndLaneMapping", "", "MFCWithEmptyRamp", 0,
    4, 3, 2, 0, "DISPATCH%02d", JSON'{"1":3, "2":2, "3":1}', JSON'{}'
),
(
    "clientWithCount2AndDepth2AndStart1", "", "MFCWithEmptyRamp", 0,
    3, 2, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{"1":4}'
),
(
    "clientWithCount2AndDepth2AndStart2", "", "MFCWithEmptyRamp", 0,
    3, 2, 2, 2, "DISPATCH%02d", JSON'{}', JSON'{}'
);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount2AndDepth2AndStart1", "", "MFCWith1Tote", 0,
    3, 2, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "clientWithCount2AndDepth2AndStart2", "", "MFCWith1Tote", 0,
    3, 2, 2, 2, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("1", "clientWithCount2AndDepth2AndStart1", "MFCWith1Tote", "o1", "t1", false, 1, 0),
("2", "clientWithCount2AndDepth2AndStart2", "MFCWith1Tote", "o1", "t1", false, 1, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount2AndDepth2AndStart1", "", "MFCWith2TotesFor2Orders", 0,
    3, 2, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "clientWithCount2AndDepth2AndStart2", "", "MFCWith2TotesFor2Orders", 0,
    3, 2, 2, 2, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("a1", "clientWithCount2AndDepth2AndStart1", "MFCWith2TotesFor2Orders", "o1", "t1", false, 1, 0),
("a2", "clientWithCount2AndDepth2AndStart1", "MFCWith2TotesFor2Orders", "o2", "t2", false, 2, 0),
("a3", "clientWithCount2AndDepth2AndStart2", "MFCWith2TotesFor2Orders", "o1", "t1", false, 1, 0),
("a4", "clientWithCount2AndDepth2AndStart2", "MFCWith2TotesFor2Orders", "o2", "t2", false, 2, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount2AndDepth2AndStart1", "", "MFCWith3TotesFor3Orders", 0,
    3, 2, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "clientWithCount2AndDepth2AndStart2", "", "MFCWith3TotesFor3Orders", 0,
    3, 2, 2, 2, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("t1", "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders", "o1", "t1", false, 1, 0),
("t2", "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders", "o2", "t2", false, 2, 0),
("t3", "clientWithCount2AndDepth2AndStart1", "MFCWith3TotesFor3Orders", "o3", "t3", false, 1, 0),
("t4", "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders", "o1", "t1", false, 1, 0),
("t5", "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders", "o2", "t2", false, 2, 0),
("t6", "clientWithCount2AndDepth2AndStart2", "MFCWith3TotesFor3Orders", "o3", "t3", false, 1, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount3AndDepth2AndStart1", "", "MFCWith2TotesFor2Orders", 0,
    4, 3, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "clientWithCount3AndDepth2AndStart3", "", "MFCWith2TotesFor2Orders", 0,
    4, 3, 2, 3, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("td1", "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2Orders", "o1", "t1", false, 1, 0),
("td2", "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2Orders", "o2", "t2", false, 2, 0),
("td3", "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2Orders", "o1", "t1", false, 1, 0),
("td4", "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2Orders", "o2", "t2", false, 2, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount3AndDepth2AndStart1", "", "MFCWith2TotesFor2OrdersWitGap", 0,
    4, 3, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
),
(
    "clientWithCount3AndDepth2AndStart3", "", "MFCWith2TotesFor2OrdersWitGap", 0,
    4, 3, 2, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("tb1", "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2OrdersWitGap", "o1", "t1", false, 1, 0),
("tb2", "clientWithCount3AndDepth2AndStart1", "MFCWith2TotesFor2OrdersWitGap", "o3", "t3", false, 3, 0),
("tb3", "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2OrdersWitGap", "o1", "t1", false, 1, 0),
("tb4", "clientWithCount3AndDepth2AndStart3", "MFCWith2TotesFor2OrdersWitGap", "o3", "t3", false, 3, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithCount3AndDepth2AndLaneMapping", "", "MFCWith2TotesFor2OrdersWitGap", 0,
    4, 3, 2, 0, "DISPATCH%02d", JSON'{"1":2, "2":3, "3":1}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("tc1", "clientWithCount3AndDepth2AndLaneMapping", "MFCWith2TotesFor2OrdersWitGap", "o1", "t1", false, 1, 0),
("tc2", "clientWithCount3AndDepth2AndLaneMapping", "MFCWith2TotesFor2OrdersWitGap", "o3", "t3", false, 2, 0);



INSERT INTO configs (client_id, env, mfc_id, updated_at, error_ramp, count, depth, start, id_gen, lane_mapping, express_lane_mapping)
VALUES
(
    "clientWithTotesToDelete", "", "MFCWithTotesToDelete", 0,
    4, 3, 3, 1, "DISPATCH%02d", JSON'{}', JSON'{}'
);
INSERT INTO tote_assignments (id, client_id, mfc_id, order_id, tote_id, is_express, lane_idx, created_at) VALUES
("ttd1", "clientWithTotesToDelete", "MFCWithTotesToDelete", "o1", "t1", false, 1, 0),
("ttd4", "clientWithTotesToDelete", "MFCWithTotesToDelete", "o2", "t2", false, 2, 0),
("ttd6", "clientWithTotesToDelete", "MFCWithTotesToDelete", "o3", "t3", false, 3, 0);
