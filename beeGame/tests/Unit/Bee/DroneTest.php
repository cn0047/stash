<?php

namespace Test\Unit\Bee;

use Bee\Drone;
use VO\PositiveInteger;

class DroneTest extends \PHPUnit_Framework_TestCase
{
    public function testGetIsQueen()
    {
        $bee = new Drone(new PositiveInteger(50), new PositiveInteger(12));
        static::assertFalse($bee->getIsQueen());
    }
}
