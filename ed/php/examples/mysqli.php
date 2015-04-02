<?php

$db = new mysqli('localhost', 'root');
if ($db->connect_errno) {
    echo 'Connect failed: '.$db->connect_error;
    exit();
}
$db->select_db('test');
$cursor = $db->query('SHOW TABLES');
if ($db->error) {
    echo $db->error.PHP_EOL;
}
while($row = $cursor->fetch_assoc()){
    var_export($row);
}
$cursor->close();
