<?php

interface Queue
{
    public function get();
    public function add($item);
}

interface Stack
{
    public function get();
    public function add($item);
}

/**
 * Prior to PHP 5.3.9, a class could not implement two interfaces
 * that specified a method with the same name, since it would cause ambiguity.
 */
class Dispatcher implements Queue, Stack
{
    public function get()
    {
    }

    public function add($item)
    {
    }
}
