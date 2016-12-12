<?php

try {
    foo(200);
} catch (Exception $e) {
    echo 'CAUGHT ERROR'.PHP_EOL;
    var_export($e->getMessage().PHP_EOL);
}

function foo(array $boo)
{
}

/*
PHP Catchable fatal error:  Argument 1 passed to foo() must be of the type array, integer given...
*/
