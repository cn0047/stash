<?php

class Foo
{
    public function get()
    {
        return static::A;
    }
}

class Boo extends Foo
{
    protected const A = 'Foo B';
}

var_dump((new Boo())->get());

/*
string(5) "Foo B"
*/
