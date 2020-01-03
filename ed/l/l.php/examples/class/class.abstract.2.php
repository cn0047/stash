<?php

namespace examples;

/**
 * Classes defined as abstract may not be instantiated,
 * and any class that contains at least one abstract method must also be abstract.
 */
abstract class Foo
{
    const BAR = 'BAR';

    private function get()
    {
    }

    protected function getValue()
    {
    }

    public function getBar()
    {
    }

    // abstract private function set();
    // PHP Fatal error:  Abstract function examples\Foo::set() cannot be declared private in ed/php/examples/class.abstract.2.php on line 15

    /**
     * All methods marked abstract in the parent's class declaration
     * must be defined by the child,
     * additionally, these methods must be defined with the same (or a less restricted) visibility.
     * If the abstract method is defined as protected,
     * the function implementation must be defined as either protected or public, but not private.
     */
    abstract protected function setValue();

    abstract public function setBar();
}

// PHP Fatal error:  Cannot instantiate abstract class examples\Foo in gh/ed/php/examples/class.abstract.2.php on line 37
$instance = new Foo();
