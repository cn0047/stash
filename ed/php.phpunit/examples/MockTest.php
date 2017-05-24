<?php

class Foo
{
    public function bar($str = 'bar')
    {
        return $str;
    }
}

class MockTest extends PHPUnit_Framework_TestCase
{
    public function testOrigin()
    {
        $m = new Foo;
        $this->assertSame($m->bar(), 'bar');
    }

    public function testMock()
    {
        $m = $this->getMock('Foo');
        $m->expects($this->once())
            ->method('bar')
            ->will($this->returnValue('404'));
        $this->assertSame($m->bar(), '404');
    }

    public function test2Mock()
    {
        $m = $this->getMock('Foo');
        $m->expects($this->once())
            ->method('bar')
            ->with($this->equalTo('boo'))
            ->will($this->returnValue('404'));
        $this->assertSame($m->bar('boo'), '404');
    }
}
