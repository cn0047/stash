<?php

trait Foo
{
    abstract public function getName();
}

class Boo
{
    use Foo;

    public function getName()
    {
    }
}
