<?php

$item = 5;
$list = [1, 2, 3];
foreach ($list as $item) {
    var_dump($item);
    $item++;
}
echo PHP_EOL.$item.PHP_EOL;
/*
int(1)
int(2)
int(3)

4
*/
