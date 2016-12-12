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

class e
{
    private function __construct()
    {
    }
}

class f extends a
{
    /**
     * PHP Fatal error:  Access level to e::__construct() must be public (as in class a)
     */
    // private function __construct() {}
}

class g
{
    private function __construct()
    {
    }
}

new b;
new c;
new d;
new e;
/*
construct a.
construct c.
construct a.
PHP Fatal error:  Call to private e::__construct() from invalid context
*/
