<?php

$pdo = new PDO('sqlite::memory:');
$pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
$sql = <<<"SQL"
CREATE TABLE types (
    i INTEGER
    ,n NUMERIC
    ,r REAL
    ,b BLOB
    ,t TEXT
)
SQL;
$pdo->exec($sql);
$pdo->exec("INSERT INTO types VALUES (1, 2, 3, 4, '2015-08-19')");
$s = $pdo->prepare('SELECT * FROM types');
if (!$s->execute()) {
    throw new \RuntimeException(var_export($s->errorInfo(), true));
}
$r = $s->fetchAll(\PDO::FETCH_ASSOC);
var_dump($r);

/*
array(1) {
  [0] =>
  array(5) {
    'i' => string(1) "1"
    'n' => string(1) "2"
    'r' => string(3) "3.0"
    'b' => string(1) "4"
    't' => string(1) "2015-08-19"
  }
}
*/
