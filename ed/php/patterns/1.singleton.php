<?php
/**
 * Singleton
 *
 * @category Creational
 */

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

class One
{
    public $key;

    public function __construct()
    {
        $this->key = uniqid();
    }
}

class Singleton2
{
    protected static $instance = null;

    public static function getInstance()
    {
        if (!isset(self::$instance)) {
            self::$instance = new One;
        }
        return self::$instance;
    }
}

$foo = Singleton2::getInstance();
$boo = Singleton2::getInstance();
$foo2 = new Singleton2;
$foo2 = $foo2::getInstance();
$boo2 = new Singleton2;
$boo2 = $boo2::getInstance();
var_export([
    $foo->key,
    $boo->key,
    $foo2->key,
    $boo2->key,
]);

/*
array (
  0 => '55669fc7ec6a0',
  1 => '55669fc7ec6a0',
  2 => '55669fc7ec6a0',
  3 => '55669fc7ec6a0',
)
*/
