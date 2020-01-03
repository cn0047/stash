<?php

/**
 * PHP 5.5.9
 */
class foo
{
    private $value = 'foo';

    public function __clone()
    {
        echo 'Cloning...'.PHP_EOL;
    }

    public function render()
    {
        echo $this->value.PHP_EOL;
    }

    public function set($value)
    {
        $this->value = $value;
    }
}

$o = new foo;
$o->render();
echo PHP_EOL;

$o2 = $o;
$o2->set('bar');
$o->render();
$o2->render();
echo PHP_EOL;

$o3 = clone $o;
$o3->set('boo');
$o->render();
$o2->render();
$o3->render();
echo PHP_EOL;
/*
foo

bar
bar

Cloning...
bar
bar
boo
*/
