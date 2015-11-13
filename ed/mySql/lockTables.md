Lock and unlock tables
-

Locks may be used to emulate transactions or to get more speed when updating tables.
A session holding a WRITE lock can perform table-level operations such as DROP TABLE or TRUNCATE TABLE.
For sessions holding a READ lock, DROP TABLE and TRUNCATE TABLE operations are not permitted.
LOCK TABLES is permitted (but ignored) for a TEMPORARY table.

Rules for Lock Acquisition:
* READ [LOCAL] lock:
    * The session that holds the lock can read the table (but not write it).
    * Multiple sessions can acquire a READ lock for the table at the same time.
    * Other sessions can read the table without explicitly acquiring a READ lock.
    * The LOCAL modifier enables nonconflicting INSERT statements (concurrent inserts) by other sessions
      to execute while the lock is held.
* [LOW_PRIORITY] WRITE lock:
    * The session that holds the lock can read and write the table.
    * Only the session that holds the lock can access the table.
      No other session can access it until the lock is released.
    * Lock requests for the table by other sessions block while the WRITE lock is held.

WRITE locks normally have higher priority than READ locks.

````sql
LOCK TABLES t1 READ;
SELECT COUNT(*) FROM t1;
+----------+
| COUNT(*) |
+----------+
|        3 |
+----------+
SELECT COUNT(*) FROM t2;
ERROR 1100 (HY000): Table 't2' was not locked with LOCK TABLES
````

You cannot refer to a locked table multiple times in a single query using the same name.
Use aliases instead, and obtain a separate lock for the table and each alias:
````sql
LOCK TABLE t WRITE, t AS t1 READ;
INSERT INTO t SELECT * FROM t;
ERROR 1100: Table 't' was not locked with LOCK TABLES
INSERT INTO t SELECT * FROM t AS t1;
````

Rules for Lock Release:
* A session can release its locks explicitly with UNLOCK TABLES.
* If a session issues a LOCK TABLES statement to acquire a lock while already holding locks,
  its existing locks are released implicitly before the new locks are granted.
* If a session begins a transaction (for example, with START TRANSACTION),
  an implicit UNLOCK TABLES is performed, which causes existing locks to be released.
* If the connection for a client session terminates, whether normally or abnormally,
  the server implicitly releases all table locks held by the session.
* If you use ALTER TABLE on a locked table, it may become unlocked.

Example:

| session 1 query                                      | session 1 result              | session 2 query               | session 2 result |
|------------------------------------------------------|-------------------------------|-------------------------------|------------------|
| select count(*) from product;                        | 5                             |                               |                  |
|                                                      |                               | select count(*) from product; | 5                |
| lock tables product write;                           |                               |                               |                  |
|                                                      |                               | select count(*) from product; | 5                |
| insert into product set category_id=1, name = 'boo'; |                               |                               |                  |
|                                                      | select count(*) from product; | --hang                        |                  |
| unlock tables;                                       |                               |                               |                  |
|                                                      |                               |                               | 6                |
