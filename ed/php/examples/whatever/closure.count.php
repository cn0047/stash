<?php

$func = function () {
    $i = 0;
    return function () use (&$i) {
        $i = $i + 1;
        return $i;
    };
};
$f = $func();

var_export([
    $f(),
    $f(),
    $f(),
]);
