<?php

class Foo
{
    const A = 'Foo A';
}

class Boo extends Foo
{
    const A = 'Foo B';
}

$o = new Boo;
var_dump($o::A);

/*
string(5) "Foo B"
*/
