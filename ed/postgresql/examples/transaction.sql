CREATE TABLE IF NOT EXISTS task_progress (
    action CHARACTER VARYING(100) NOT NULL,
    total INTEGER NOT NULL DEFAULT 0,
    done INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expire INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (action)
);

-- INIT:
DELETE FROM task_progress;

INSERT INTO task_progress (action, total, expire)
VALUES ('t', 1000100, 0)
ON CONFLICT (action) DO UPDATE
SET total = task_progress.total + 1000100, expire = task_progress.expire + 0
;

SELECT * FROM task_progress WHERE action = 't';

-- CASE 1:
| Terminal 1                                            | Terminal 2                                                    |
|-------------------------------------------------------|---------------------------------------------------------------|
| BEGIN;                                                |                                                               |
|                                                       | BEGIN;                                                        |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 0                                  |                                                               |
| UPDATE task_progress SET done = 1 WHERE action = 't'; |                                                               |
| COMMIT;                                               |                                                               |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 1                                  |                                                               |
|                                                       | SELECT * FROM task_progress WHERE action = 't';               |
|                                                       | -- Result: done eq 1                                          |
|                                                       | COMMIT;                                                       |

-- CASE 2:
| Terminal 1                                            | Terminal 2                                                    |
|-------------------------------------------------------|---------------------------------------------------------------|
|                                                       | BEGIN;                                                        |
| BEGIN;                                                |                                                               |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 0                                  |                                                               |
| UPDATE task_progress SET done = 1 WHERE action = 't'; |                                                               |
| COMMIT;                                               |                                                               |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 1                                  |                                                               |
|                                                       | SELECT * FROM task_progress WHERE action = 't';               |
|                                                       | -- Result: done eq 1                                          |
|                                                       | COMMIT;                                                       |

-- CASE 3 (⭐️):
| Terminal 1                                            | Terminal 2                                                    |
|-------------------------------------------------------|---------------------------------------------------------------|
|                                                       | BEGIN;                                                        |
|                                                       | SELECT * FROM task_progress WHERE action = 't';               |
|                                                       | -- Result: done eq 0                                          |
| BEGIN;                                                |                                                               |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 0                                  |                                                               |
| UPDATE task_progress SET done = 1 WHERE action = 't'; |                                                               |
| COMMIT;                                               |                                                               |
| SELECT * FROM task_progress WHERE action = 't';       |                                                               |
| -- Result: done eq 1                                  |                                                               |
|                                                       | UPDATE task_progress SET done = 1 WHERE action = 't';         |
|                                                       | COMMIT;                                                       |
|                                                       | SELECT * FROM task_progress WHERE action = 't';               |
|                                                       | -- Result: done eq 1                                          |

-- CASE 4 (⭐️):
| Terminal 1                                              | Terminal 2                                                      |
| ------------------------------------------------------- | --------------------------------------------------------------- |
| BEGIN;                                                  |                                                                 |
|                                                         | BEGIN;                                                          |
| SELECT * FROM task_progress WHERE action = 't';         |                                                                 |
| -- Result: done eq 0                                    |                                                                 |
|                                                         | SELECT * FROM task_progress WHERE action = 't';                 |
|                                                         | -- Result: done eq 0                                            |
| UPDATE task_progress SET done = 1 WHERE action = 't';   |                                                                 |
|                                                         | UPDATE task_progress SET done = 1 WHERE action = 't';           |
|                                                         | -- hanging                                                      |
| SELECT * FROM task_progress WHERE action = 't';         |                                                                 |
| -- Result: done eq 1                                    |                                                                 |
| COMMIT;                                                 |                                                                 |
|                                                         | -- stopped hang                                                 |
|                                                         | SELECT * FROM task_progress WHERE action = 't';                 |
|                                                         | -- Result: done eq 1                                            |
|                                                         | COMMIT;                                                         |
|                                                         | SELECT * FROM task_progress WHERE action = 't';                 |
|                                                         | -- Result: done eq 1                                            |
