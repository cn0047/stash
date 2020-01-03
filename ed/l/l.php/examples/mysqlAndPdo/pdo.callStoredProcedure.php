<?php

$pdo = new PDO('mysql:host=xmysql;port=3306;dbname=test', 'dbu', 'dbp'); // OK
$s = $pdo->prepare('CALL getCountry(:p1)');
$params = [':p1' => 'ua'];
$s->execute($params);
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
