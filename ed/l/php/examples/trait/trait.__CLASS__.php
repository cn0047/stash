<?php

trait Foo
{
    public function get()
    {
        return __CLASS__;
    }
}

Class Bar
{
    use Foo;
}

$o = new Bar;
var_dump($o->get());
/*
string(3) "Bar"
*/
