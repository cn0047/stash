<?php

class Foo
{
    private $bar;

    public function __construct()
    {
        $this->bar = 200;
    }

    public function getBar()
    {
        return $this->bar;
    }
}

class Boo extends Foo
{
}

$o = new Boo;
echo $o->getBar().PHP_EOL;
