<?php

set_error_handler(
    function ($errorCode, $errorDescription, $errorFile, $errorLine, array $errorContext = []) {
        throw new Exception($errorDescription, $errorCode);
    }
);

try {
    foo(200);
} catch (Exception $e) {
    var_export('ERROR:'.$e->getMessage().PHP_EOL);
}

function foo(array $boo)
{
}

/*
ERROR:Argument 1 passed to foo() must be of the type array, integer given...
*/
