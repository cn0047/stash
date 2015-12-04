<?php

$a = [];
foreach (range(1, 16777215) as $i) {
    $c = crc32($i);
    if (isset($a[$c])) {
        echo "$c   {$a[$c]}   $i\n";
    } else {
        $a[$c] = $i;
    }
}

var_export([
    [crc32('80623'), crc32('14746802')],
    [crc32('51742'), crc32('14797963')],
    [crc32('90868'), crc32('14756649')],
]);
