<?php

interface Queue
{
    public function add($item);
}

/**
 * Interfaces can be extended like classes using the extends operator.
 */
interface MFU extends Queue
{
    public function scan();
    public function printToPaper();
}

class JetM1212 implements MFU
{
    public function scan()
    {
    }

    public function printToPaper()
    {
    }
}

new JetM1212;

/*
PHP Fatal error:  Class JetM1212 contains 1 abstract method and must therefore be declared abstract or implement the remaining methods (Queue::add)
*/
