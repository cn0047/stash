<?php

$dbh = new PDO('mysql:host=127.0.0.1;dbname=test', 'root');
$sth = $dbh->prepare('SELECT NOW()');
$sth->execute();
$result = $sth->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
