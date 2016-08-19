<?php

trait Foo3
{
    public function getName()
    {
        echo 'Foo3'.PHP_EOL;
    }
}

class Bar3
{
    public function getName()
    {
        echo 'Bar3'.PHP_EOL;
    }
}

class Boo3 extends Bar3
{
    use Foo3 {
        getName as public;
    }
}

$boo = new Boo3;
$boo->getName();
/*
Foo3
*/
