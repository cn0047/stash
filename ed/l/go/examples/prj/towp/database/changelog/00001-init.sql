--liquibase formatted sql
--preconditions onFail:HALT onError:HALT

-- changeset changelog/00001-init.sql::1::V.Kovpak

CREATE TABLE wave_plan (
  wave_plan_id STRING(36) NOT NULL,
  retailer_id STRING(64) NOT NULL,
  mfc_id STRING (64) NOT NULL,
  timezone STRING (200) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by STRING(36) NOT NULL
) PRIMARY KEY (wave_plan_id);

CREATE UNIQUE INDEX wave_plan_retailer_mfc_unique_idx ON wave_plan (
  retailer_id, mfc_id
);

CREATE TABLE wave (
  wave_plan_id STRING(36) NOT NULL,
  wave_id STRING(36) NOT NULL,
  cutoff_time STRING(5) NOT NULL,
  from_time STRING(5) NOT NULL,
  to_time STRING(5) NOT NULL
) PRIMARY KEY (wave_plan_id, wave_id), INTERLEAVE IN PARENT wave_plan ON DELETE CASCADE;

CREATE UNIQUE INDEX wave_wave_plan_id_cutoff_time_unique_idx ON wave (
  wave_plan_id, cutoff_time
);

CREATE TABLE schedule (
  wave_plan_id STRING(36) NOT NULL,
  wave_id STRING(36) NOT NULL,
  schedule_id STRING(36) NOT NULL,
  schedule_type STRING(50) NOT NULL,
  schedule_time STRING(10) NOT NULL
) PRIMARY KEY (wave_plan_id, wave_id, schedule_id), INTERLEAVE IN PARENT wave ON DELETE CASCADE;

CREATE TABLE trigger (
  wave_plan_id STRING(36) NOT NULL,
  wave_id STRING(36) NOT NULL,
  schedule_id STRING(36) NOT NULL,
  trigger_at TIMESTAMP NOT NULL,
  cutoff_datetime TIMESTAMP NOT NULL,
  created_at TIMESTAMP,
  fired_at TIMESTAMP
) PRIMARY KEY (wave_plan_id, wave_id, schedule_id, trigger_at), INTERLEAVE IN PARENT schedule ON DELETE CASCADE;

--rollback DROP TABLE trigger;
--rollback DROP TABLE schedule;
--rollback DROP INDEX wave_wave_plan_id_cutoff_time_unique_idx;
--rollback DROP TABLE wave;
--rollback DROP INDEX wave_plan_retailer_mfc_unique_idx;
--rollback DROP TABLE wave_plan;
