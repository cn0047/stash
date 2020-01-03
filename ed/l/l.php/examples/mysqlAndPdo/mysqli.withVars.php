<?php

$db = new mysqli('localhost', 'root');
if ($db->connect_errno) {
    echo 'Connect failed: '.$db->connect_error;
    exit();
}
$db->select_db('test');
$id = 5;
$sql = 'SELECT Id, title FROM tree WHERE Id > ?';
$stmt = $db->prepare($sql);
$stmt->bind_param('s', $id);
$stmt->execute();
$stmt->bind_result($Id, $title);
$data = array();
while ($stmt->fetch()) {
   $data[] = [$Id, $title];
}
$db->close();
var_export($data);
