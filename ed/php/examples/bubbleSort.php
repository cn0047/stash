<?php

$s = 'fswelxvda';
$c = strlen($s);
for ($i = 0; $i < $c; $i++) {
    for ($j = 0; $j < $c - $i - 1; $j++) {
        if ($s[$j] < $s[$j+1]) {
            $v = $s[$j];
            $s[$j] = $s[$j+1];
            $s[$j+1] = $v;
        }
    }
}

var_dump($s);

/*
string(9) "xwvslfeda"
*/
