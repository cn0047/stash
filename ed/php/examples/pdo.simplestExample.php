<?php

$dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
$sth = $dbh->prepare('SELECT NOW()');
$sth->execute();
$result = $sth->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
