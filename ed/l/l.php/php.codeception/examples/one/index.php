<?php

require __DIR__ . '/bootstrap.php';

$foo = new App\Foo();
var_dump($foo->bar());
var_dump($foo->sbar());
var_dump($foo->baz());
var_dump($foo->sbaz());
