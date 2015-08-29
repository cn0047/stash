<?php

trait Foo
{
    private $name = 'Foo';
}

class Boo
{
    use Foo;

    public function getName()
    {
        return $this->name;
    }
}

$boo = new Boo();
var_dump($boo->getName());
/*
Foo
*/
