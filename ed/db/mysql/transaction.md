Transaction
-

#### Transaction Isolation Levels

Isolation is the I in the acronym ACID.
The default isolation level for InnoDB is REPEATABLE READ.

* REPEATABLE READ -
  All consistent reads within the same transaction read the snapshot established by the first read.

* READ COMMITTED -
  Each consistent read, even within the same transaction, sets and reads its own fresh snapshot.

* READ UNCOMMITTED (weakest) - This is also called a dirty read.

* SERIALIZABLE (strongest)
  This level is like REPEATABLE READ,
  but InnoDB implicitly converts all plain `SELECT` statements to `SELECT ... LOCK IN SHARE MODE`.
  It therefore is known to be read only and can be serialized if performed as a consistent (nonlocking) read
  and need not block for other transactions.

`SELECT ... LOCK IN SHARE MODE` sets a shared mode lock on any rows that are read.
Other sessions can read the rows, but cannot modify them until your transaction commits.
If any of these rows were changed by another transaction that has not yet committed,
your query waits until that transaction ends and then uses the latest values.

````
[mysqld]
transaction-isolation=REPEATABLE-READ
````
````
set session tx_isolation='READ-COMMITTED';
````

#### Example

````sql
drop table if exists tree;
create table if not exists tree (id int, title varchar(50));
insert into tree values (9, 'service 9');
insert into tree values (8, 'service 8');
insert into tree values (7, 'service 7');
insert into tree values (6, 'service 6');
````

Case 1:

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| SELECT title FROM tree WHERE id = 9; -- service 9  |                                                             |
| START TRANSACTION;                                 |                                                             |
| UPDATE tree SET title = 'service 99' WHERE id = 9; |                                                             |
| -- Query OK                                        |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 9; -- service 9           |
|                                                    | START TRANSACTION;                                          |
|                                                    | UPDATE tree SET title = 'service 999' WHERE id = 9; -- hang |
| COMMIT;                                            |                                                             |
|                                                    | -- Query OK                                                 |
| SELECT title FROM tree WHERE id = 9; -- service 99 |                                                             |
|                                                    | COMMIT;                                                     |
|                                                    | SELECT title FROM tree WHERE id = 9; -- service 999         |
| SELECT title FROM tree WHERE id = 9; -- service 999|                                                             |

Case 2:

| terminal 1                                         | terminal 2                                                                    |
|----------------------------------------------------|-------------------------------------------------------------------------------|
| START TRANSACTION;                                 |                                                                               |
| UPDATE tree SET title = 'service 99' WHERE id = 9; |                                                                               |
| -- Query OK                                        |                                                                               |
|                                                    | START TRANSACTION;                                                            |
|                                                    | UPDATE tree SET title = 'service 999' WHERE id = 9; -- hang                   |
| -- leave it for a while                            |                                                                               |
|                                                    | -- ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction |

Case 3:

| terminal 1                                               | terminal 2                                                  |
|----------------------------------------------------------|-------------------------------------------------------------|
|                                                          | START TRANSACTION;                                          |
|                                                          | SELECT title FROM tree WHERE id = 9; -- service 9           |
| START TRANSACTION;                                       |                                                             |
| SELECT title FROM tree WHERE id = 9; -- service 9        |                                                             |
| UPDATE tree SET title = CONCAT(title, "a") WHERE id = 9; |                                                             |
| COMMIT;                                                  |                                                             |
|                                                          | SELECT title FROM tree WHERE id = 9; -- service 9           |
|                                                          | UPDATE tree SET title = CONCAT(title, "b") WHERE id = 9;    |
|                                                          | COMMIT;                                                     |
|                                                          | SELECT title FROM tree WHERE id = 9; -- service 9ab         |

#### READ COMMITTED

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| SELECT title FROM tree WHERE id = 8; -- service 8  |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 8; -- service 8           |
| SET session tx_isolation='READ-COMMITTED';         |                                                             |
| START TRANSACTION;                                 |                                                             |
| UPDATE tree SET title = 'service 88' WHERE id = 8; |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 8; -- service 8           |
| commit;                                            |                                                             |
| SELECT title FROM tree WHERE id = 8; -- service 88 |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 8; -- service 88          |

#### READ UNCOMMITTED

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| SELECT title FROM tree WHERE id = 7; -- service 7  |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 7; -- service 7           |
| SET session tx_isolation='READ-UNCOMMITTED';       |                                                             |
| START TRANSACTION;                                 |                                                             |
| UPDATE tree SET title = 'service 77' WHERE id = 7; |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 7; -- service 7           |
| COMMIT;                                            |                                                             |
| SELECT title FROM tree WHERE id = 7; -- service 77 |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 7; -- service 77

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| SELECT title FROM tree WHERE id = 6; -- service 6  |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 6; -- service 6           |
| SET session tx_isolation='READ-UNCOMMITTED';       |                                                             |
| START TRANSACTION;                                 |                                                             |
| UPDATE tree SET title = 'service 66' WHERE id = 6; |                                                             |
|                                                    | SELECT title FROM tree WHERE id = 6; -- service 6           |
|                                                    | SET session tx_isolation='READ-UNCOMMITTED';                |
|                                                    | START TRANSACTION;                                          |
|                                                    | UPDATE tree SET title = 'service 666' WHERE id = 6; -- hang |

#### Nested transaction

````sql
drop table if exists tree;
create table if not exists tree (id int, title varchar(50));

START TRANSACTION;
insert into tree values (1, 'service 1');

START TRANSACTION;
insert into tree values (2, 'service 2');

commit;
SELECT * FROM tree;
````
result (😮):
````
+------+-----------+
| id   | title     |
+------+-----------+
|    1 | service 1 |
|    2 | service 2 |
+------+-----------+
````
