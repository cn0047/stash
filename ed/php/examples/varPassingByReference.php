<?php

function func(&$r)
{
    $r++;
}

$r = 1;
func(func($r));
echo $r;
/*
2
*/

echo "\n";

function foo(&$bar)
{
    $bar *= 2;
    return $bar;
}

$x = 3;
$y = foo($x);
$x = 5;
echo $x.', '.$y;
/*
5, 6
*/
