<?php

class Script
{
    private $id;
    private $foo = 'foo';
    protected $boo = 'boo';
    public $bar = 'bar';

    public function getId()
    {
        return __METHOD__;
    }

    public function __get($name)
    {
        return "Magic: $name";
    }
}

$o = new Script();
var_export([
    $o->id,
    $o->foo,
    $o->boo,
    $o->bar,
]);

/*
array (
  0 => 'Magic: id',
  1 => 'Magic: foo',
  2 => 'Magic: boo',
  3 => 'bar',
)
*/
