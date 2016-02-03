<?php

namespace Bee;

use VO\PositiveInteger;

abstract class Bee
{
    protected $lifespan = 0;

    protected $deduceStep = 0;

    protected $points = 0;

    public function __construct(PositiveInteger $lifespan, PositiveInteger $deduceStep)
    {
        $this->lifespan = $lifespan->get();
        $this->deduceStep = $deduceStep->get();
        $this->points = $this->lifespan;
    }

    public function getPoints()
    {
        return $this->points;
    }

    public function hit()
    {
        $this->points -= $this->deduceStep;
    }

    public function getIsAlive()
    {
        return $this->points > 0;
    }

    public function getIsQueen()
    {
        return false;
    }
}
