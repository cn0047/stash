<?php

interface Queue
{
    public function add($item);
}

new Queue;

/*
PHP Fatal error:  Cannot instantiate interface Queue
*/
