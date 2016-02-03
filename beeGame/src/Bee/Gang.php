<?php

namespace Bee;

/**
 * Aggregate of bees.
 *
 * this is aggregate
 * that allows manage all bees in particular game,
 * and provides necessary methods to understand what is going on.
 */
class Gang
{
    private $bees = [];

    private $isQueenAlive = false;

    public function add(Bee $bee)
    {
        if ($bee->getIsQueen()) {
            $this->isQueenAlive = true;
        }
        $this->bees[] = $bee;
    }

    public function shuffle()
    {
        shuffle($this->bees);
    }

    public function getCount()
    {
        return count($this->bees);
    }

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

    public function getIsQueenAlive()
    {
        return $this->isQueenAlive;
    }

    public function getStatistics()
    {
        $data = [];
        foreach ($this->bees as $bee) {
            $data[get_class($bee)][] = $bee->getPoints();
        }
        return $data;
    }
}
