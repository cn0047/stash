<?php

namespace Bee;

/**
 * Class Drone.
 */
class Drone extends Bee
{
    /**
     * {@inheritdoc}
     */
    public function getIsQueen()
    {
        return false;
    }
}
