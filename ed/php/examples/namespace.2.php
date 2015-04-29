<?php

namespace Foo;

function strlen($string) {
    return 5;
}

function strlen2($string) {
    return 5;
}

namespace Bar;
use Foo;

var_dump([
    strlen('baz'),  // int(3)
    strlen2('baz'), // PHP Fatal error:  Call to undefined function Bar\strlen2() in ...
]);
