Singleton
-

Generator.

````php
<?php

class Singleton
{
    protected static $instance = null;

    protected function __construct() {}

    protected function __clone() {}

    public static function getInstance()
    {
        if (!isset(static::$instance)) {
            // late static binding
            static::$instance = new static;
        }
        return static::$instance;
    }
}

class Foobar extends Singleton {};
$foo = Foobar::getInstance();
````
