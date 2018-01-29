<?php

trait Foo
{
    public function __construct()
    {
      var_dump('Constructed FOO');
    }

    public function foo()
    {
        return 'FOO';
    }
}

Class Bar
{
    use Foo;

    public function __construct()
    {
      var_dump('Constructed BAR');
    }
}

$o = new Bar;

/*
string(15) "Constructed BAR"
*/
