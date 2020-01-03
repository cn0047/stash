<?php

$link = mysql_connect('localhost', 'root');
if (!$link) {
    throw new Exception(mysql_error());
}
$selectedDB = mysql_select_db('test', $link);
if (!$selectedDB) {
    throw new Exception(mysql_error());
}

$result = mysql_query('SELECT NOW()');
if (!$result) {
    throw new Exception('Invalid query: '.mysql_error());
}
$data = [];
while ($row = mysql_fetch_assoc($result)) {
    $data[] = $row;
}
var_export($data);
