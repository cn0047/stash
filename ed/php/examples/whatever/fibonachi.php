<?php
/**
 * @deprecated
 */

// 1, 1, 2, 3, 5, 8, 13, 21 ...

function fib ($n) {
    $f[0] = 1;
    $f[1] = 1;
    for ($i = 2; $i <= $n; $i++) {
        $f[$i] = $f[$i-1] + $f[$i-2];
    }
    return $f[$n];
}

var_export([
    fib(0),
    fib(1),
    fib(2),
    fib(3),
    fib(4),
    fib(5),
]);

/*
array (
  0 => 1,
  1 => 1,
  2 => 2,
  3 => 3,
  4 => 5,
  5 => 8,
)
*/
