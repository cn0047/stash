<?php

interface Queue
{
    final public function add($item);
}

/*
PHP Fatal error:  Access type for interface method Queue::add() must be omitted
*/
