<?php
/**
 * State
 *
 * @category Behaviour
 */

interface EngineState
{
    public function startEngine();
    public function moveForward();
}

class EngineTurnedOnState implements EngineState
{
    public function startEngine()
    {
        echo 'Engine already started'.PHP_EOL;
    }

    public function moveForward()
    {
        echo 'Moved Car forward...'.PHP_EOL;
        return $this;
    }
}

class EngineTurnedOffState implements EngineState
{
    public function startEngine()
    {
        echo 'Started Engine.'.PHP_EOL;
        return new EngineTurnedOnState;
    }

    public function moveForward()
    {
        echo 'Have to start engine first!'.PHP_EOL;
    }
}

class Car implements EngineState
{
    protected $state;

    public function __construct()
    {
        $this->state = new EngineTurnedOffState;
    }

    public function startEngine()
    {
        $this->state = $this->state->startEngine();
    }

    public function moveForward()
    {
        $this->state = $this->state->moveForward();
    }
}

$car = new Car;
$car->moveForward();
$car = new Car;
$car->startEngine();
$car->moveForward();

/*
Have to start engine first!
Started Engine.
Moved Car forward...
*/
