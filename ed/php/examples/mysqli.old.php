<?php

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
