<?php

namespace Bee;

use VO\PositiveInteger;

/**
 * Class Bee.
 *
 * This class contains common behaviour for bee.
 */
abstract class Bee
{
    /** @var integer $lifespan Contains points value for initialization purposes. */
    protected $lifespan = 0;

    /** @var integer $deduceStep Step to deduce points during hit. */
    protected $deduceStep = 0;

    /** @var integer $points Points value. */
    protected $points = 0;

    /**
     * Bee constructor.
     *
     * @param PositiveInteger $lifespan This value used for init bee points when game is starting.
     * @param PositiveInteger $deduceStep Each hit will deduce exactly this value from points of bee.
     */
    public function __construct(PositiveInteger $lifespan, PositiveInteger $deduceStep)
    {
        $this->lifespan = $lifespan->get();
        $this->deduceStep = $deduceStep->get();
        $this->points = $this->lifespan;
    }

    /**
     * Gets points value.
     *
     * @return integer Points value.
     */
    public function getPoints()
    {
        return $this->points;
    }

    /**
     * Deduce points of bee.
     */
    public function hit()
    {
        $this->points -= $this->deduceStep;
    }

    /**
     * Provides ability to understand is this bee alive.
     *
     * @return boolean Is alive.
     */
    public function getIsAlive()
    {
        return $this->points > 0;
    }

    /**
     * Provides ability to understand is this bee is a Queen.
     *
     * @return boolean Is queen.
     */
    abstract public function getIsQueen();
}
