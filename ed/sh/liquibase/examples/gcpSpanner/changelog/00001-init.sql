--liquibase formatted sql
--preconditions onFail:HALT onError:HALT

-- changeset changelog/00001-init.sql::1::k

CREATE TABLE test_1 (
    id STRING(64) NOT NULL,
    data JSON
) PRIMARY KEY (id);

CREATE TABLE test_2 (
    id STRING(64) NOT NULL,
    data JSON
) PRIMARY KEY (id);

--rollback DROP TABLE test_1;
--rollback DROP TABLE test_2;
