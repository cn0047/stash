<?php

$a = range(1, 80);
$max = 5;
$i = 0;
$r = new SplMaxHeap();
$r->insert(33);
$r->insert(11);
$r->insert(2);
$r->insert(2);
$r->insert(22);

var_export([
    $r->count(),
    $r->top(),
]);

/*
array (
  0 => 5,
  1 => 33,
)
*/
