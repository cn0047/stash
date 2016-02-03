<?php

namespace Test\Unit\Bee;

use Bee\Bee;
use Bee\Gang;
use Bee\Queen;
use Bee\Worker;
use VO\PositiveInteger;

class GangTest extends \PHPUnit_Framework_TestCase
{
    /** @var Gang */
    private $beeGang;

    protected function setUp()
    {
        $this->beeGang = new Gang();
    }

    public function testGetCount()
    {
        static::assertSame(0, $this->beeGang->getCount());
    }

    public function testAdd()
    {
        $this->beeGang->add(new Worker(new PositiveInteger(75), new PositiveInteger(10)));
        static::assertSame(1, $this->beeGang->getCount());
    }

    public function testGetIsQueenAlive()
    {
        $this->beeGang->add(new Queen(new PositiveInteger(100), new PositiveInteger(8)));
        static::assertTrue($this->beeGang->getIsQueenAlive());
    }

    public function testRandomHit()
    {
        $this->beeGang->add(new Worker(new PositiveInteger(75), new PositiveInteger(10)));
        $this->beeGang->randomHit();
        // One available way to break encapsulation.
        $reflection = new \ReflectionObject($this->beeGang);
        $property = $reflection->getProperty('bees');
        $property->setAccessible(true);
        /** @var Bee $bee */
        $bee = $property->getValue($this->beeGang)[0];
        static::assertSame(65, $bee->getPoints());
    }
}