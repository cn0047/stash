<?php

$pdo = new PDO('mysql:host=localhost;port=3306;dbname=test', 'root');
$s = $pdo->prepare('CALL getCountry(:p1)');
$params = [':p1' => 'ua'];
$s->execute($params);
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
