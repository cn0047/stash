Lock tables
-

Table-level Lock Modes:
* ACCESS SHARE
* ROW SHARE
* ROW EXCLUSIVE
* SHARE UPDATE EXCLUSIVE
* SHARE
* SHARE ROW EXCLUSIVE
* EXCLUSIVE
* ACCESS EXCLUSIVE

Row-level Locks:
* FOR UPDATE
* FOR NO KEY UPDATE
* FOR SHARE
* FOR KEY SHARE

PostgreSQL provides a means for creating locks
that have application-defined meanings, these are called advisory locks.
