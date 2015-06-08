<?php

trait Foo
{
    abstract public function getName();
}

class Boo
{
    use Foo;

    public function get()
    {
    }
}

/*
PHP Fatal error:  Class Boo contains 1 abstract method and must therefore be declared abstract or implement the remaining methods (Boo::getName)
*/
