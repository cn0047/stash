<?php

class foo
{
    private $value = 'foo';

    protected function __clone()
    {
    }
}

$o = new foo;
$c = clone $o;

/*
PHP Fatal error:  Call to protected foo::__clone()
*/
