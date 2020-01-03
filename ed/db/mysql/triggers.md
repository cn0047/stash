Triggers
-

A trigger is activates when a statement inserts, updates, or deletes
(they do not activate for changes in views).
There cannot be multiple triggers for a given table that have
the same trigger event and action time (till version 5.7).

WARNING:
Cascaded foreign key actions do not activate triggers.
Triggers work unsafe in statement based replication.

In an INSERT trigger, only NEW.col_name can be used.
In a DELETE trigger, only OLD.col_name can be used.
In an UPDATE trigger, you can use
OLD.col_name to refer to the columns of a row before it is updated
and NEW.col_name to refer to the columns of the row after it is updated.

A column named with OLD is read only.
In a BEFORE trigger, you can SET NEW.col_name = value
if you have the UPDATE privilege for it.

In a BEFORE trigger, the NEW value for an AUTO_INCREMENT column is 0.

````sql
CREATE TABLE account (acct_num INT, amount DECIMAL(10, 2)) ;
CREATE TRIGGER ins_sum BEFORE INSERT ON account
FOR EACH ROW SET @sum = @sum + NEW.amount;

SET @sum = 0;
INSERT INTO account VALUES(137, 14.98), (141, 1937.50), (97, -100.00);
SELECT @sum AS 'Total amount inserted';

DROP TRIGGER test.ins_sum;
-- If you drop a table, any triggers for the table are also dropped.

CREATE TRIGGER ins_transaction BEFORE INSERT ON account
FOR EACH ROW PRECEDES ins_sum
SET
@deposits = @deposits + IF(NEW.amount>0, NEW.amount, 0),
@withdrawals = @withdrawals + IF(NEW.amount<0, -NEW.amount, 0);
````

By using the BEGIN ... END construct,
you can define a trigger that executes multiple statements.
````sql
delimiter //
CREATE TRIGGER upd_check BEFORE UPDATE ON account
FOR EACH ROW
BEGIN
    IF NEW.amount < 0 THEN
        SET NEW.amount = 0;
    ELSEIF NEW.amount > 100 THEN
        SET NEW.amount = 100;
    END IF;
END;//
delimiter ;
````

With purpose abort action (insert, update, etc) from trigger
you need use this: `SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'My error.';`.

````sql
CREATE TABLE test1(a1 INT);
CREATE TABLE test2(a2 INT);
CREATE TABLE test3(a3 INT NOT NULL AUTO_INCREMENT PRIMARY KEY);
CREATE TABLE test4(
  a4 INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  b4 INT DEFAULT 0
);

delimiter |
CREATE TRIGGER testref BEFORE INSERT ON test1
  FOR EACH ROW
  BEGIN
    INSERT INTO test2 SET a2 = NEW.a1;
    DELETE FROM test3 WHERE a3 = NEW.a1;
    UPDATE test4 SET b4 = b4 + 1 WHERE a4 = NEW.a1;
  END;
|
delimiter ;

INSERT INTO test3 (a3) VALUES
  (NULL), (NULL), (NULL), (NULL), (NULL),
  (NULL), (NULL), (NULL), (NULL), (NULL)
;
INSERT INTO test4 (a4) VALUES
  (0), (0), (0), (0), (0), (0), (0), (0), (0), (0)
;

select * from test1;
-- Empty set (0.00 sec)

select * from test2;
-- Empty set (0.00 sec)

select * from test3;
+----+
| a3 |
+----+
|  1 |
|  2 |
|  3 |
|  4 |
|  5 |
|  6 |
|  7 |
|  8 |
|  9 |
| 10 |
+----+
-- 10 rows in set (0.00 sec)

select * from test4;
+----+------+
| a4 | b4   |
+----+------+
|  1 |    0 |
|  2 |    0 |
|  3 |    0 |
|  4 |    0 |
|  5 |    0 |
|  6 |    0 |
|  7 |    0 |
|  8 |    0 |
|  9 |    0 |
| 10 |    0 |
+----+------+
-- 10 rows in set (0.00 sec)

-- MOST INTERESTING:
INSERT INTO test1 VALUES
(1), (3), (1), (7), (1), (8), (4), (4);

select * from test1;
+------+
| a1   |
+------+
|    1 |
|    3 |
|    1 |
|    7 |
|    1 |
|    8 |
|    4 |
|    4 |
+------+
-- 8 rows in set (0.00 sec)

select * from test2;
+------+
| a2   |
+------+
|    1 |
|    3 |
|    1 |
|    7 |
|    1 |
|    8 |
|    4 |
|    4 |
+------+
-- 8 rows in set (0.00 sec)

select * from test3;
+----+
| a3 |
+----+
|  2 |
|  5 |
|  6 |
|  9 |
| 10 |
+----+
-- 5 rows in set (0.00 sec)

select * from test4;
+----+------+
| a4 | b4   |
+----+------+
|  1 |    3 |
|  2 |    0 |
|  3 |    1 |
|  4 |    2 |
|  5 |    0 |
|  6 |    0 |
|  7 |    1 |
|  8 |    1 |
|  9 |    0 |
| 10 |    0 |
+----+------+
-- 10 rows in set (0.00 sec)
````

````sql
SHOW CREATE TRIGGER ins_sum;
SHOW TRIGGERS;
````
