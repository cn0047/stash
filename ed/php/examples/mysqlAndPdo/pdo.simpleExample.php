<?php

try {
    $dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
} catch (PDOException $e) {
    echo 'Connection failed: '.$e->getMessage();
}

$date = '2015-03-13';
$sth = $dbh->prepare('SELECT :date AS `date`');
$sth->bindParam(':date', $date, PDO::PARAM_STR);
if (!$sth->execute()) {
    throw new Exception($sth->errorInfo());
}
$result = $sth->fetchColumn();
print($result);

/**
 * In loop.
 */
foreach (['one', 'tow', 'three'] as $value) {
    $sth->bindParam(':date', $value, PDO::PARAM_STR);
    if (!$sth->execute()) {
        throw new Exception($sth->errorInfo());
    }
    $result = $sth->fetch(PDO::FETCH_ASSOC);
    var_export($result);
}
