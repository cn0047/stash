<?php

namespace Test\Unit\VO;

use VO\PositiveInteger;

class PositiveIntegerTest extends \PHPUnit_Framework_TestCase
{
    public function testOk()
    {
        $int = new PositiveInteger(200);
        static::assertSame(200, $int->get());
    }

    /**
     * @expectedException \InvalidArgumentException
     */
    public function testFail()
    {
        new PositiveInteger('200');
    }
}
