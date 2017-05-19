<?php

try {
    $dbh = new PDO('mysql:host=mysql-master;port=3306;dbname=test', 'dbu', 'dbp');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}
$s = $dbh->prepare("SELECT NOW() AS 'date from master'");
if (!$s->execute()) {
    throw new Exception($s->errorInfo());
}
$result = $s->fetch();
var_export($result);

try {
    $dbh = new PDO('mysql:host=mysql-slave-1;port=3306;dbname=test', 'dbu', 'dbp');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}
$s = $dbh->prepare("SELECT NOW() AS 'date from slave'");
if (!$s->execute()) {
    throw new Exception($s->errorInfo());
}
$result = $s->fetch();
var_export($result);
