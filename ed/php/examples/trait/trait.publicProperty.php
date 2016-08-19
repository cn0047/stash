<?php

trait Foo
{
    public $name = 'Foo';
}

Class Bar
{
    use Foo;
}

$o = new Bar;
var_dump($o->name);
/*
string(3) "Foo"
*/
