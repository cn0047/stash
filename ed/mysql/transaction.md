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
transaction-isolation = REPEATABLE-READ
````
````
set session tx_isolation='READ-COMMITTED';
````

#### Example

````sql
create table if not exists tree (id int, title varchar(50));
insert into tree values (9, 'service 9');
insert into tree values (8, 'service 8');
insert into tree values (7, 'service 7');
insert into tree values (6, 'service 6');
````

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| select title from tree where id = 9; -- service 9  |                                                             |
| start transaction;                                 |                                                             |
| update tree set title = 'service 99' where id = 9; |                                                             |
| -- Query OK                                        |                                                             |
|                                                    | select title from tree where id = 9; -- service 9           |
|                                                    | start transaction;                                          |
|                                                    | update tree set title = 'service 999' where id = 9; -- hang |
| commit;                                            |                                                             |
|                                                    | -- Query OK                                                 |
| select title from tree where id = 9; -- service 99 |                                                             |
|                                                    | commit;                                                     |
|                                                    | select title from tree where id = 9; -- service 999         |
| select title from tree where id = 9; -- service 999|                                                             |

#### READ COMMITTED

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| select title from tree where id = 8; -- service 8  |                                                             |
|                                                    | select title from tree where id = 8; -- service 8           |
| set session tx_isolation='READ-COMMITTED';         |                                                             |
| start transaction;                                 |                                                             |
| update tree set title = 'service 88' where id = 8; |                                                             |
|                                                    | select title from tree where id = 8; -- service 8           |
| commit;                                            |                                                             |
| select title from tree where id = 8; -- service 88 |                                                             |
|                                                    | select title from tree where id = 8; -- service 88          |

#### READ UNCOMMITTED

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| select title from tree where id = 7; -- service 7  |                                                             |
|                                                    | select title from tree where id = 7; -- service 7           |
| set session tx_isolation='READ-UNCOMMITTED';       |                                                             |
| start transaction;                                 |                                                             |
| update tree set title = 'service 77' where id = 7; |                                                             |
|                                                    | select title from tree where id = 7; -- service 7           |
| commit;                                            |                                                             |
| select title from tree where id = 7; -- service 77 |                                                             |
|                                                    | select title from tree where id = 7; -- service 77

| terminal 1                                         | terminal 2                                                  |
|----------------------------------------------------|-------------------------------------------------------------|
| select title from tree where id = 6; -- service 6  |                                                             |
|                                                    | select title from tree where id = 6; -- service 6           |
| set session tx_isolation='READ-UNCOMMITTED';       |                                                             |
| start transaction;                                 |                                                             |
| update tree set title = 'service 66' where id = 6; |                                                             |
|                                                    | select title from tree where id = 6; -- service 6           |
|                                                    | set session tx_isolation='READ-UNCOMMITTED';                |
|                                                    | start transaction;                                          |
|                                                    | update tree set title = 'service 666' where id = 6; -- hang |
