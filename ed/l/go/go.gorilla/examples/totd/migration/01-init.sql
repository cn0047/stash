CREATE TABLE t_assignments (
  id STRING(36) NOT NULL,
  client_id STRING(100) NOT NULL,
  mfc_id STRING(100) NOT NULL,
  order_id STRING(100) NOT NULL,
  t_id STRING(100) NOT NULL,
  is_express BOOL NOT NULL,
  lane_idx INT64 NOT NULL,
  created_at INT64 NOT NULL
) PRIMARY KEY (id);

CREATE UNIQUE INDEX t_assignments_unique_idx ON t_assignments (
  client_id, mfc_id, order_id, t_id
);

CREATE TABLE t_assignments_log (
  id STRING(36) NOT NULL,
  state JSON,
  input_data JSON,
  output_data JSON,
  created_at INT64 NOT NULL
) PRIMARY KEY (id);

CREATE TABLE configs (
  client_id STRING(100) NOT NULL,
  env STRING(100) NOT NULL,
  mfc_id STRING(100) NOT NULL,
  updated_at INT64 NOT NULL,
  error_ramp INT64 NOT NULL,
  count INT64 NOT NULL,
  depth INT64 NOT NULL,
  start INT64 NOT NULL,
  id_gen STRING(100) NOT NULL,
  lane_mapping JSON,
  express_lane_mapping JSON,
  flow_racks_mapping JSON
) PRIMARY KEY (client_id, env, mfc_id);
