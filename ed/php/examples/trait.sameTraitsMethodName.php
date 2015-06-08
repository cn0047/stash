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
}

$o = new Boo;
$o->getName();

/*
PHP Fatal error:  Trait method getName has not been applied, because there are collisions with other trait methods on Boo
*/
