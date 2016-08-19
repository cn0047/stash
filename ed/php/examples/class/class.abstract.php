<?php

abstract class Foo
{
    abstract public function bar(array $args = null);

    public function baz()
    {
        return 'baz';
    }
}

class Boo extends Foo
{
    public function bar(array $args)
    {
        echo count($args);
    }
}

$o = new Boo;
$o->bar([5]);
/*
PHP Fatal error:  Declaration of Boo::bar() must be compatible with Foo::bar(array $args = NULL)
*/
