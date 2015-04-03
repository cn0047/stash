<?php

$zero  = 0b0000;
$one   = 0b0001;
$two   = 0b0010;
$three = 0b0011;

var_dump(($one & $three) ==  $one);
var_dump(($one & $one)   ==  $one);

/*
bool(true)
bool(true)
*/
