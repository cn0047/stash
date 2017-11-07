<?php

$dbh = new PDO('pgsql:host=postgres-master;port=5432;dbname=test;user=dbu;password=dbp');
$s = $dbh->prepare('SELECT NOW()');
$s->execute();
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
