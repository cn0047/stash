<?php

namespace Bee;

/**
 * Aggregate of bees.
 *
 * This is aggregate
 * that allows manage all bees in particular game,
 * and provides necessary methods to understand what is going on with bees.
 */
class Gang
{
    /** @var array $bees Contains all bees inside gang. */
    private $bees = [];

    /** @var boolean $isQueenAlive Contains boolean flag that explain is queen bee alive. */
    private $isQueenAlive = false;

    /**
     * Add bee to the gang.
     *
     * @param Bee $bee Bee.
     */
    public function add(Bee $bee)
    {
        if ($bee->getIsQueen()) {
            $this->isQueenAlive = true;
        }
        $this->bees[] = $bee;
    }

    /**
     * Shuffle bee inside gang.
     *
     * This method is optional, and it provides more random behaviour to game.
     */
    public function shuffle()
    {
        shuffle($this->bees);
    }

    /**
     * Gets count of bee inside gang.
     *
     * @return integer Count.
     */
    public function getCount()
    {
        return count($this->bees);
    }

    /**
     * This method encapsulate logic for hit any single bee from gang.
     *
     * This method randomly find bee and hit her.
     * After all bees inside gang will be re-indexed.
     */
    public function randomHit()
    {
        $index = mt_rand(0, $this->getCount()-1);
        /** @var Bee $bee */
        $bee = $this->bees[$index];
        $bee->hit();
        if (!$bee->getIsAlive()) {
            if ($bee->getIsQueen()) {
                $this->isQueenAlive = false;
            }
            unset($this->bees[$index]);
            // When we do unset - we delete particular key,
            // and array will contains empty spot - it can provide bugs,
            // that's why we do re-index to avoid unexpected behaviour.
            $this->bees = array_values($this->bees);
        }
    }

    /**
     * Provides ability to understand is queen bee alive.
     *
     * When queen bee out of own points it have impact for all game,
     * this method helps to track it.
     *
     * @return boolean Is queen bee alive.
     */
    public function getIsQueenAlive()
    {
        return $this->isQueenAlive;
    }

    /**
     * Gets statistics about bees inside gang.
     *
     * @return array Statistics.
     */
    public function getStatistics()
    {
        $data = [];
        /** @var Bee $bee */
        foreach ($this->bees as $bee) {
            $data[get_class($bee)][] = $bee->getPoints();
        }
        return $data;
    }
}
