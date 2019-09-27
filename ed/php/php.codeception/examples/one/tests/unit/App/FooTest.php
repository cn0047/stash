<?php

use App\Foo;
use App\Boo;
use Codeception\Test\Unit;
use Codeception\Util\Stub;

class FooTest extends Unit
{
    public function testBar()
    {
        $foo = new Foo();
        $this->assertEquals($foo->bar(), 'bar');
    }

    public function testSBar()
    {
        $this->assertEquals(Foo::sbar(), 'sbar');
    }

    public function testBaz()
    {
        $foo = new Foo();
        $this->assertEquals($foo->baz(), 'baz');
    }

    public function testBazMock()
    {
        $boo = Stub::make(Boo::class, ['baz' => 'yes']);

        $foo = new Foo();
        $this->assertEquals($foo->baz(), 'yes');
    }
}