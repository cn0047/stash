MySql
-

*5.1.72-rel14.10*

````sql
COLLATE UTF8_GENERAL_CI LIKE

SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SHOW COLUMNS FROM table LIKE '%';

INSERT INTO tbl1 (field1) SELECT field2 FROM tbl2;

SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'tableName';
````
####Functions
````sql
UNIX_TIMESTAMP()
````

####Options
````sql
SET SQL_SAFE_UPDATES=0;
````

####Tricks
````
tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
````
