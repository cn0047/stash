<?php

namespace Bee;

use VO\NotEmptyString;
use VO\PositiveInteger;

class Factory
{
    public function create(NotEmptyString $className, PositiveInteger $lifespan, PositiveInteger $deduceStep)
    {
        $name = "Bee\\$className";
        if (!class_exists($name)) {
            throw new \InvalidArgumentException('Invalid bee class name.');
        }
        return new $name($lifespan, $deduceStep);
    }
}
