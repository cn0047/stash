<?php

class a
{
    public function __construct()
    {
        echo 'construct a.'.PHP_EOL;
    }
}

class b extends a
{
}

class c extends a
{
    public function __construct()
    {
        echo 'construct c.'.PHP_EOL;
    }
}

class d extends b
{
}

new b;
new c;
new d;
/*
construct a.
construct c.
construct a.
*/
