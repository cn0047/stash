<?php
/**
 * Flyweight
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
