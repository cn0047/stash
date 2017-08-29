<?php

interface Foo
{
    public function bar($str = 'bar');
}

class MockInterfaceTest extends PHPUnit_Framework_TestCase
{
    public function testMock()
    {
        $m = $this->createMock('Foo');
        $m->expects($this->once())
            ->method('bar')
            ->will($this->returnValue('404'));
        $this->assertSame($m->bar(), '404');
    }
}
