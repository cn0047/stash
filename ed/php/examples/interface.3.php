<?php

class Faz
{
}

class Baz extends Faz
{
}

interface Foo
{
    public function du(Faz $param);
}

class Boo implements Foo
{
    public function du(Baz $param)
    {
        echo false;
    }
}

$c = new Boo;
$c->du(new Baz);

/*
PHP Fatal error:  Declaration of Boo::du() must be compatible with Foo::du(Faz $param)
*/
