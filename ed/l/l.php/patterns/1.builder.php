<?php
/**
 * Builder
 *
 * @category Creational
 *
 * @see https://github.com/cn007b/my/blob/master/ed/php.symfony/examples/_bulletinBoard/src/AppBundle/Entity/UserRepository.php#L15
 */

abstract class BuilderAbstract
{
    public function buildPartA() {}

    public function buildPartB() {}

    abstract public function getResult();
}

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

class Product {}

class Director extends DirectorAbstract {}

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

$builder = new Builder();
$director = new Director($builder);
$director->construct();
$product = $builder->getResult();
