<?php

var_export([
    is_numeric('1.3'),
    is_numeric('+3'),
    is_numeric('+0123.45e6'),
    is_numeric(0b10100111001),
    is_numeric(0xf4c3b00c),
]);

/*
array (
  0 => true,
  1 => true,
  2 => true,
  3 => true,
  4 => true,
)
*/
