<?php

trait Foo
{
    final public function get()
    {
        return 'Foo';
    }
}

Class Bar
{
    use Foo;
}

$o = new Bar;
var_dump($o->get());
/*
string(3) "Foo"
*/
