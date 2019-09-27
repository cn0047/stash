<?php

namespace App;

use App\Boo;

class Foo
{
    public function bar()
    {
        return 'bar';
    }

    public static function sbar()
    {
        return 'sbar';
    }

    public function baz()
    {
        return (new Boo())->baz();
    }

    public function sbaz()
    {
        return Boo::sbaz();
    }
}
