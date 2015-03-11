<?php

class foo
{
    public function bar()
    {
        return 'bar';
    }
}

class MockTest extends PHPUnit_Framework_TestCase
{
    public function testOrigin()
    {
        $m = new foo;
        $this->assertSame($m->bar(), 'bar');
    }

    public function testMock()
    {
        $m = $this->getMock('foo');
        $m->expects($this->any())->method('bar')->will($this->returnValue('404'));
        $this->assertSame($m->bar(), '404');
    }
}
