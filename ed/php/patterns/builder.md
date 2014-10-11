Builder
-

Group 1

````php
<?php

abstract class DirectorAbstract
{
    protected $builder = null;

    public function __construct(BuilderAbstract $builder)
    {
        $this->builder = $builder;
    }

    public function construct()
    {
        $this->builder->buildPartA();
        $this->builder->buildPartB();
    }
}

class Director extends DirectorAbstract {}

abstract class BuilderAbstract
{
    public function buildPartA() {}

    public function buildPartB() {}

    abstract public function getResult();
}

class Builder extends BuilderAbstract
{
    protected $product = null;

    public function __construct()
    {
        $this->product = new Product();
    }

    public function buildPartA() {}

    public function buildPartB() {}

    public function getResult()
    {
        return $this->product;
    }
}

class Product {}

$builder = new Builder();
$director = new Director($builder);
$director->construct();
$product = $builder->getResult();
````
