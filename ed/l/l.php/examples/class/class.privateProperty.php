<?php

/**
 * Members declared as private may only be accessed by the class that defines the member.
 */
class A
{
    private $p = 200;

    function expose(A $x)
    {
        return $x->p;
    }
}

$a1 = new A();
$a2 = new A();

var_dump($a1->expose($a2));

/*
int(200)
*/
