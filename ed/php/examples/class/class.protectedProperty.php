<?php

/**
 * Members declared protected
 * can be accessed only within the class itself and by inherited and parent classes.
 */
class A
{
    protected $p = 200;
}

class B extends A
{
    function expose($x)
    {
        return $x->p;
    }
}

$a = new A();
$b = new B();

var_dump($b->expose($a));

/*
int(200)
*/
