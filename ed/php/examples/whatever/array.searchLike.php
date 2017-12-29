<?php

$a = [
    'mike_45',
    'peter_23',
    'jim_12',
];

var_export(preg_grep("/^mike.*/", $a));
/*
array (
  0 => 'mike_45',
)
*/
