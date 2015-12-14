<?php

$sem = sem_get(getmypid(), 1);
if (sem_acquire($sem, true)) {
    echo "acquire\n";
    sleep(5);
    sem_release($sem);
    echo "release\n";
} else {
    echo "already locked\n";
}
