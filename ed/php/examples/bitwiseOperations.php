<?php

$zero  = 0b00;
$one   = 0b01;
$two   = 0b10;
$three = 0b11;

var_dump(($one & $three) ==  $one);
var_dump(($one & $one)   ==  $one);

/*
bool(true)
bool(true)
*/
