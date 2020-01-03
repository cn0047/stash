<?php

$message = 'hello';
$example = function () {
    var_dump($message);
};
echo $example();
/*
PHP Notice:  Undefined variable: message
*/

$example = function () use ($message) {
    var_dump($message);
};
echo $example();
/*
string(5) "hello"
*/

$message = 'world';
echo $example();
/*
string(5) "hello"
*/
