<?php

$dbh = new PDO('mysql:dbname=test;host=127.0.0.1', 'root');
$pid = 1;
$qid = 2;
$text = '3';
$sql = 'SELECT * FROM tbl WHERE 1 = 1';
$params = [];
if (isset($pid)) {
    $sql .= ' AND pid = :pid';
    $params[':pid'] = $pid;
}
if (isset($qid)) {
    $sql .= ' AND qid = :qid';
    $params[':qid'] = $qid;
}
if (isset($text)) {
    $sql .= ' AND `text` = :text';
    $params[':text'] = $text;
}
$s = $dbh->prepare($sql);
$s->execute($params);
$result = $s->fetchAll(PDO::FETCH_ASSOC);
var_export($result);
