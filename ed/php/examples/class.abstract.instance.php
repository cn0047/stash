<?php

/**
 * Classes defined as abstract may not be instantiated,
 * and any class that contains at least one abstract method must also be abstract.
 */
abstract class Foo
{
}

new Foo;

/*
PHP Fatal error:  Cannot instantiate abstract class Foo
*/
