<?php

trait Foo
{
    private $name = 'Foo';

    public function getName()
    {
        return $this->name;
    }
}

class Boo
{
    use Foo;
}

$boo = new Boo;
var_dump($boo->getName());
/*
string(3) "Foo"
*/
