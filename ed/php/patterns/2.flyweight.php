<?php
/**
 * Flyweight
 *
 * A flyweight is an object that minimizes memory use by sharing as much data as possible with other similar objects.
 *
 * @category Structural
 */

class FlyweightFactory
{
    protected static $flyweigths = array();

    public static function getFlyweight($key)
    {
        if (!isset(self::$flyweigths[$key])) {
            self::$flyweigths[$key] = new ConcreteFlyweight();
        }
        return self::$flyweigths[$key];
    }
}

abstract class Flyweight
{
    protected $intrinsicState = null;

    public function Operation($extrinsicState)
    {
    }
}

class ConcreteFlyweight extends Flyweight {}
class UnsharedFlyweight extends Flyweight {}
