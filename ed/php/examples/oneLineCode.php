<?php

echo sprintf('"%04d"', 1).PHP_EOL; // "0001"

var_dump(in_array('test', array(0))); // bool(true) - Because test converts to integer.

var_export([
    5 % 2,
    7 % 3,
    11 % 2,
    8 % 4
]);
/*
array (
  0 => 1,
  1 => 1,
  2 => 1,
  3 => 0,
)
*/
