<?php

try {
    $dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}

$date = '2015-03-13';
$s = $dbh->prepare('SELECT :date AS `date`');
$s->bindParam(':date', $date, PDO::PARAM_STR);
if (!$s->execute()) {
    throw new Exception($s->errorInfo());
}
$result = $s->fetchColumn();
print($result);

/**
 * In loop.
 */
foreach (['one', 'tow', 'three'] as $value) {
    $s->bindParam(':date', $value, PDO::PARAM_STR);
    if (!$s->execute()) {
        throw new Exception($s->errorInfo());
    }
    $result = $s->fetch(PDO::FETCH_ASSOC);
    var_export($result);
}
