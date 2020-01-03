<?php

namespace Foo;

class Bar
{
    private function baz() {}
}

$class = new \ReflectionClass('Foo\Bar');
$method = $class->getMethod('baz');
echo $method->getName().PHP_EOL;
/*
baz
*/
