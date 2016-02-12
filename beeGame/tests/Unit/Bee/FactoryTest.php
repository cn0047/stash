<?php

namespace Test\Unit\Bee;

use Bee\Factory;
use Bee\Worker;
use VO\NotEmptyString;
use VO\PositiveInteger;

class FactoryTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $factory = new Factory();
        $lifespan = new PositiveInteger(100);
        $deduceStep = new PositiveInteger(8);
        $bee = $factory->create(new NotEmptyString('Worker'), $lifespan, $deduceStep);
        $worker= new Worker($lifespan, $deduceStep);
        static::assertEquals($worker, $bee);
    }

    /**
     * @expectedException \InvalidArgumentException
     */
    public function testCreateFail()
    {
        $factory = new Factory();
        $lifespan = new PositiveInteger(100);
        $deduceStep = new PositiveInteger(8);
        $factory->create(new NotEmptyString('Bumblebee'), $lifespan, $deduceStep);
    }
}
