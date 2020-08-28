<?php

class Foo
{
    private $count = 0;
    private static $totalCount = 0;

    public function increment()
    {
        $this->count += 1;
        self::$totalCount += 1;
    }

    public function getCount()
    {
        return $this->count;
    }

    public function getTotalCount()
    {
        return self::$totalCount;
    }
}


$f1 = new Foo();
$f1->increment();
$f1->increment();

$f2 = new Foo();
$f2->increment();

print($f1->getCount()."\n");        // 2
print($f1->getTotalCount()."\n");   // 3
print($f2->getCount()."\n");        // 1
print($f2->getTotalCount()."\n");   // 3
print(Foo::getTotalCount()."\n");   // 3
