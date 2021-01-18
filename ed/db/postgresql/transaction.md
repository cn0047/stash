Transaction
-

````sql
-- BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;
BEGIN;
  -- update valid statement
SAVEPOINT point_1;
  -- update error statement
ROLLBACK TO point_1;
  -- update valid statement
COMMIT;
````

### Transaction Isolation Levels:
* Read uncommitted
* Read committed (default)
* Repeatable read
* Serializable (strictest)
