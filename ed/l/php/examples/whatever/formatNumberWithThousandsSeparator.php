<?php
/**
 * Format number with thousands separator.
 */
$a = 1000000;
$a = (string)$a;
$r = '';
$j = 0;
for ($i = strlen($a); $i > 0 ; $i--) {
    if ($j == 3) {
        $j = 0;
        $r = ' '.$r;
    }
    $r = $a[$i-1].$r;
    $j++;
}
var_export($r);

/*
1 000 000
*/
