<?php

namespace examples;

interface Foo
{
    /**
     * It's possible for interfaces to have constants.
     * Interface constants works exactly like class constants
     * except they cannot be overridden by a class/interface that inherits them.
     */
    const BAR = 'BAR';

    // private function get();
    // PHP Fatal error:  Access type for interface method Foo::get() must be omitted in ed/php/examples/interface.2.php on line 17

    // protected function getValue();
    // PHP Fatal error:  Access type for interface method Foo::getValue() must be omitted in ed/php/examples/interface.2.php on line 20

    /**
     * All methods declared in an interface must be public; this is the nature of an interface.
     */
    public function getBar();
}
