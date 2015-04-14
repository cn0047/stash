<?php

try {
    $dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}

$date = '2015-03-13';
$sth = $dbh->prepare('SELECT * FROM mostCurrentRow WHERE week = :date ORDER BY score DESC');
$sth->bindParam(':date', $date, PDO::PARAM_STR);
if (!$sth->execute()) {
    throw new Exception($sth->errorInfo());
}
$result = $sth->fetchAll(PDO::FETCH_ASSOC);
echo '<pre>';
var_export($result);
