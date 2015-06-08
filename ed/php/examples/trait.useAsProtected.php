<?php

trait Foo2
{
    public function getName()
    {
        echo 'Foo2'.PHP_EOL;
    }
}

class Bar2
{
    public function getName()
    {
        echo 'Bar2'.PHP_EOL;
    }
}

class Boo2 extends Bar2
{
    use Foo2  {
        getName as protected;
    }
}

$boo = new Boo2;
$boo->getName();
/*
PHP Fatal error:  Access level to Foo2::getName() must be public (as in class Bar2)
*/
