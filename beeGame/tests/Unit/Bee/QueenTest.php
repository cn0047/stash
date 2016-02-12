<?php

namespace Test\Unit\Bee;

use Bee\Queen;
use VO\PositiveInteger;

class QueenTest extends \PHPUnit_Framework_TestCase
{
    public function testGetIsQueen()
    {
        $bee = new Queen(new PositiveInteger(50), new PositiveInteger(12));
        static::assertTrue($bee->getIsQueen());
    }
}
