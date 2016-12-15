<?php

$dbh = new PDO('mysql:host=127.0.0.1;dbname=test', 'root');
$s = $dbh->prepare('SELECT NOW()');
$s->execute();
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
