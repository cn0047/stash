<?php

abstract class Foo
{
    abstract public function bar(array $params = null);
}

class Moo extends Foo
{
    public function bar(array $params)
    {
        echo count($params);
    }
}

$c = new Moo;
$c->bar(array(5));

/*
PHP Fatal error:  Declaration of Moo::bar() must be compatible with Foo::bar(array $params = NULL)
*/
