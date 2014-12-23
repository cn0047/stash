<?php

trait Foo
{
    public function getName()
    {
        echo 'Foo'.PHP_EOL;
    }
}

trait Bar
{
    public function getName()
    {
        echo 'Bar'.PHP_EOL;
    }
}

class Boo
{
    use Foo, Bar;

    public function getName()
    {
        echo 'Boo'.PHP_EOL;
    }
}

$boo = new Boo;
$boo->getName();
/*
Boo
*/