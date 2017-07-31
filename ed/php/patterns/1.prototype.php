<?php
/**
 * Prototype
 *
 * It is used when the type of objects to create is determined by a prototypical instance.
 * For example: daily meeting in calendar has same members, same time, etc.
 * You don't need to provide such information each time, just instantiate prototype class and use it.
 *
 * @category Creational
 */

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
