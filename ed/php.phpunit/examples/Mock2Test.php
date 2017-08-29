<?php

class Foo
{
    public function bar()
    {
        $r = $this->boo('bar');
        return "Result: $r";
    }

    public function boo($str)
    {
        return "boo $str";
    }
}

class Mock2Test extends PHPUnit_Framework_TestCase
{
    public function testOrigin()
    {
        $m = new Foo;
        $this->assertSame($m->bar(), 'Result: boo bar');
    }

    public function testMock()
    {
        $m = $this->createMock('Foo', ['boo']);
        $m->expects($this->once())
            ->method('boo')
            ->will($this->returnValue('404'));
        $this->assertSame($m->bar(), 'Result: 404');
    }
}
