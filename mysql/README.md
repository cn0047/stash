MySql
-

*5.1.72-rel14.10*

````sql
COLLATE UTF8_GENERAL_CI LIKE

SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SHOW COLUMNS FROM table LIKE '%';
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
