<?php

namespace Test\Unit\Bee;

use Bee\Drone;
use Bee\Queen;
use Bee\Worker;
use VO\PositiveInteger;

class BeeTest extends \PHPUnit_Framework_TestCase
{
    public function testGetPoints()
    {
        $bee = new Queen(new PositiveInteger(100), new PositiveInteger(8));
        static::assertSame(100, $bee->getPoints());
    }

    public function testHit()
    {
        $bee = new Queen(new PositiveInteger(100), new PositiveInteger(8));
        $bee->hit();
        static::assertSame(92, $bee->getPoints());
    }

    public function testGetIsAlive()
    {
        $bee = new Worker(new PositiveInteger(100), new PositiveInteger(8));
        static::assertTrue($bee->getIsAlive());
    }

    public function testIsNotAlive()
    {
        $bee = new Drone(new PositiveInteger(0), new PositiveInteger(8));
        static::assertFalse($bee->getIsAlive());
    }

    public function testGetIsQueen()
    {
        $bee = new Queen(new PositiveInteger(100), new PositiveInteger(8));
        static::assertTrue($bee->getIsQueen());
    }

    public function testGetIsNotQueen()
    {
        $bee = new Drone(new PositiveInteger(100), new PositiveInteger(8));
        static::assertFalse($bee->getIsQueen());
    }
}
