Prototype
-

Group 1.

````php
<?php

// Method 1.
class ClassA
{
    public function __construct(ClassA $prototype = null)
    {
        if (is_null($prototype)) {
        }
    }
}

$prototype = new ClassA();
$newObject = new ClassA($prototype);

// Method 2.
class ClassB
{
    public function getClone()
    {
       $object = new ClassB();
       return $object;
    }
}

$prototype = new ClassB();
$newObject = $prototype->getClone();

// Method 3.
class ClassC
{
    public function __clone() {}
}

$prototype = new ClassC();
$newObject = clone $prototype;
````
