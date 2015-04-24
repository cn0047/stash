<?php

class Foo
{
    private static $name = 'Foo';

    public static function getName1()
    {
        return self::$name;
    }

    public static function getName2()
    {
        return static::$name;
    }

    public static function getName3()
    {
        return static::$nameAdd;
    }
}

class Bar extends Foo
{
    private static $name = 'Bar';
    protected static $nameAdd = 'BarAdd';
}

var_export([
    Bar::getName1(),
    Bar::getName2(),
    Bar::getName3(),
]);

/*
array (
  0 => 'Foo',
  1 => 'PHP Fatal error:  Cannot access private property Bar::$name in ...',
  2 => 'BarAdd',
)
*/
