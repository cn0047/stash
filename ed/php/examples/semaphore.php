<?php

$key = $argv[1];
$sem = sem_get($key, 1);
if (sem_acquire($sem, true)) {
    echo "acquire\n";
    sleep(5);
    sem_release($sem);
    echo "release\n";
} else {
    echo "already locked\n";
}

/*
+------------+----------------+
| p1         | p2             |
+------------+----------------+
|php x.php 1 |                |
|acquire     | php x.php 1    |
|            | already locked |
|release     |                |
+------------+----------------+
*/
