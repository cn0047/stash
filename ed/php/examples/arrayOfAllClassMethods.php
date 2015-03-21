<?php
/**
* Array of all class methods.
* @link http://runnable.com/VNIJsTEgXHsl_2XP/array-of-all-class-methods-for-php
* @link http://php.net/manual/en/reflectionclass.getmethods.php#116648
*/

class foo
{
    public function __construct()
    {
    }

    private function bar()
    {
    }

    protected function boo()
    {
    }
}

$reflection = new ReflectionClass('foo');
$methods = $reflection->getMethods();
array_walk(
    $methods,
    function (&$v) {
        $v = $v->getName();
    }
);
var_export($methods);
/*
array (
  0 => '__construct',
  1 => 'bar',
  2 => 'boo',
)
*/
