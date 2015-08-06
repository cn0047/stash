<?php

class foo
{
    protected $propery1;
}

class boo extends foo
{
    private $propery2;
    protected $propery3;
    public $propery4;
}

$reflect = new ReflectionClass('boo');
$props = $reflect->getProperties();
$ownProps = [];
foreach ($props as $prop) {
    if ($prop->class === 'boo') {
        $ownProps[] = $prop->getName();
    }
}

var_export($ownProps);

/*
array (
  0 => 'propery2',
  1 => 'propery3',
  2 => 'propery4',
)
*/
