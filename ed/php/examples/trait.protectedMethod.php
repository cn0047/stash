<?php

trait Foo
{
    protected function get()
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
PHP Fatal error:  Call to protected method Bar::get() from context ''
*/
