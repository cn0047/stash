<?php

class GrandFather
{
    private   $privateGrandFather   = 'privateGrandFather';
    protected $protectedGrandFather = 'protectedGrandFather';
    public    $publicGrandFather    = 'publicGrandFather';
}

class Father extends GrandFather
{
    private   $privateFather   = 'privateFather';
    protected $protectedFather = 'protectedFather';
    public    $publicFather    = 'publicFather';

    public function __construct()
    {
        var_export([
            property_exists($this, 'privateGrandFather'),
            property_exists($this, 'protectedGrandFather'),
            property_exists($this, 'publicGrandFather'),
        ]);
    }
}

class Son extends Father
{
    private   $privateSon   = 'privateSon';
    protected $protectedSon = 'protectedSon';
    public    $publicSon    = 'publicSon';

    public function __construct()
    {
        var_export([
            // GrandFather.
            property_exists($this, 'privateGrandFather'),
            property_exists($this, 'protectedGrandFather'),
            property_exists($this, 'publicGrandFather'),
            // Father.
            property_exists($this, 'privateFather'),
            property_exists($this, 'protectedFather'),
            property_exists($this, 'publicFather'),
        ]);
    }
}


new Father;
new Son;

/*
array (
  0 => false,
  1 => true,
  2 => true,
)array (
  0 => false,
  1 => true,
  2 => true,
  3 => false,
  4 => true,
  5 => true,
)
*/
