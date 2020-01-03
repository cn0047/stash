<?php

/*
create table language(code char(2) key, name varchar(50));
insert into language values ('en', 'English'), ('es', 'Spanish');
*/

/*
GIVEN:

mysql> select * from language;
+------+---------+
| code | name    |
+------+---------+
| en   | English |
| es   | Spanish |
+------+---------+
*/

/*
WHEN:
*/
$dbh = new \PDO('mysql:host=mysql-master;port=3306;dbname=test', 'dbu', 'dbp');
$s = $dbh->prepare('SELECT * FROM language');
$s->execute();
var_export($s->fetchAll(\PDO::FETCH_KEY_PAIR));

/*
THEN:

array (
  'en' => 'English',
  'es' => 'Spanish',
)
*/
