<?php
/**
 * Memento
 *
 * Is a pattern that provides the ability to restore an object to its previous state (undo via rollback).
 *
 * @category Behaviour
 */

class Memento
{
    protected $state;

    public function __construct($state)
    {
        $this->state = $state;
    }

    public function getSavedState()
    {
        return $this->state;
    }
}

class Victim
{
    protected $state = '';

    public function setState($state)
    {
        echo 'Set state: '.$state.PHP_EOL;
        $this->state = $state;
    }

    public function saveToMemento()
    {
        return new Memento($this->state);
    }

    public function restoreFromMemento(Memento $memento)
    {
        $this->state = $memento->getSavedState();
        echo 'Get state: '.$this->state.PHP_EOL;
    }
}

$states = [];
$victim = new Victim();
$victim->setState('state one');
$states[] = $victim->saveToMemento();
$victim->setState('state two');
$states[] = $victim->saveToMemento();
$victim->setState('state three');
$states[] = $victim->saveToMemento();
$victim->restoreFromMemento($states[0]);

/*
Set state: state one
Set state: state two
Set state: state three
Get state: state one
*/
