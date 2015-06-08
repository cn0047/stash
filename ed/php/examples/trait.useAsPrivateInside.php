<?php

trait Foo
{
    public function getName()
    {
        return 'Foo';
    }
}

class Boo
{
    use Foo {
        getName as private;
    }

    public function get()
    {
        return $this->getName();
    }
}

$boo = new Boo;
var_dump($boo->get());
/*
string(3) "Foo"
*/
