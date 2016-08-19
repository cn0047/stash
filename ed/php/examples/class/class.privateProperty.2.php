<?php

class A
{
    private $p = 200;
}

class B
{
    function expose(A $x)
    {
        return $x->p;
    }
}

$b = new B();
$a = new A();

var_dump($b->expose($a));

/*
PHP Fatal error:  Cannot access private property A::$p
*/
