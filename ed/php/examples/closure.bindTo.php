<?php

class A {
    private $x = 'foo';
}
$getXCB = function () {
    return $this->x;
};
$getX = $getXCB->bindTo(new A, 'A');
var_export($getX());

/*
'foo'
*/
