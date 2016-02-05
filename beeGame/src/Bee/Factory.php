<?php

namespace Bee;

use VO\NotEmptyString;
use VO\PositiveInteger;

/**
 * Bee Factory.
 *
 * Provides ability loosely create new bee.
 */
class Factory
{
    /**
     * Creates new bee.
     *
     * @param NotEmptyString $className Name of bee class.
     * @param PositiveInteger $lifespan Value for bee constructor.
     * @param PositiveInteger $deduceStep Value for bee constructor.
     *
     * @throws \InvalidArgumentException In case when received unknown bee class name.
     *
     * @return Bee Instance of new bee.
     */
    public function create(NotEmptyString $className, PositiveInteger $lifespan, PositiveInteger $deduceStep)
    {
        $name = "Bee\\$className";
        if (!class_exists($name)) {
            throw new \InvalidArgumentException('Invalid bee class name.');
        }
        return new $name($lifespan, $deduceStep);
    }
}
