<?php

var_export([
    is_int(200),
    is_int(0b01),
    is_int(2.5),
    is_int('1'),
    is_int('2.5'),
    is_int(null),
    is_int(true),
    is_int(false),
]);

/*
array (
  0 => true,
  1 => true,
  2 => false,
  3 => false,
  4 => false,
  5 => false,
  6 => false,
  7 => false,
)
*/