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
