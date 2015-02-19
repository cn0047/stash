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

new b;
new c;
/*
construct a.
construct c.
*/
