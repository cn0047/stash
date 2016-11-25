<?php
/**
 * Factory method
 *
 * Defines an interface for creating an object,
 * but leaves the choice of its type to the subclasses,
 * creation being deferred at run-time.
 * 
 * @category Creational
 */

class Twitter
{
    public function share()
    {
        return $this->tweet();
    }
}

class Facebook
{
    public function share()
    {
        return $this->post();
    }
}

class Factory
{
    public static function create($name)
    {
        if (class_exists($name)) {
            return new $name();
        }
    }
}

$twitter = Factory::create('Twitter');
$facebook = Factory::create('Facebook');
