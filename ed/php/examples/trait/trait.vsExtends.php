<?php

trait Foo
{
    public function getName()
    {
        echo 'Foo'.PHP_EOL;
    }
}

class Bar
{
    public function getName()
    {
        echo 'Bar'.PHP_EOL;
    }
}

class Boo extends Bar
{
    use Foo;
}

$boo = new Boo;
$boo->getName();
/*
Foo
*/
