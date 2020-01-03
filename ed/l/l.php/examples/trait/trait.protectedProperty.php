<?php

trait Foo
{
    protected $name = 'Foo';
}

Class Bar
{
    use Foo;
}

$o = new Bar;
var_dump($o->name);
/*
PHP Fatal error:  Cannot access protected property Bar::$name
*/
