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

/**
 * OLD style.
 */
$db = mysqli_connect('localhost', 'root');
$sql = "SELECT 200";
if ($result = mysqli_query($db, $sql)) {
    while ($row = mysqli_fetch_assoc($result)) {
        var_export($row);
    }
    mysqli_free_result($result);
} else {
    var_export(mysqli_error($db));
}
mysqli_close($db);
