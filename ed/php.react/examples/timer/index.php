<?php

use React\Promise\Timer;

require __DIR__ . '/vendor/autoload.php';

$promise = accessSomeRemoteResource();
Timer\timeout($promise, 2.0, $loop)->then(function ($value) {
    var_dump(204);
});
