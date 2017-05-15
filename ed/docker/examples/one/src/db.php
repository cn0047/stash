<?php

try {
    $dbh = new PDO('mysql:host=mysql;port=3306;dbname=test', 'dbu', 'dbp');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}

$s = $dbh->prepare('SELECT NOW() AS `date`');
if (!$s->execute()) {
    throw new Exception($s->errorInfo());
}
$result = $s->fetchColumn();
print($result);
