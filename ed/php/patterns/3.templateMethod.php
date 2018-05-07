<?php
/**
 * Template method
 *
 * @category Behaviour
 */

abstract class Operation
{
    protected abstract function getFirstNumber();

    protected abstract function getSecondNumber();

    protected abstract function operator($a, $b);

    public function getOperationResult()
    {
        $a = $this->getFirstNumber();
        $b = $this->getSecondNumber();
        return $this->operator($a, $b);
    }
}

class Sum extends Operation
{
    private $a;
    private $b;

    public function __construct($a = 0, $b = 0)
    {
        $this->a = $a;
        $this->b = $b;
    }

    protected function getFirstNumber()
    {
        return $this->a;
    }

    protected function getSecondNumber()
    {
        return $this->b;
    }

    protected function operator($a, $b)
    {
        return $a + $b;
    }
}

class NonNegativeSubtraction extends Operation
{
    private $a;
    private $b;

    public function __construct($a = 0, $b = 0)
    {
        $this->a = $a;
        $this->b = $b;
    }

    protected function getFirstNumber()
    {
        return $this->a;
    }

    protected function getSecondNumber()
    {
        return min($this->a, $this->b);
    }

    protected function operator($a, $b)
    {
        return $a - $b;
    }

}

$sum = new Sum(84, 56);
echo $sum->getOperationResult(), PHP_EOL;
$nonNegativeSubtraction = new NonNegativeSubtraction(9, 14);
echo $nonNegativeSubtraction->getOperationResult(), PHP_EOL;
$nonNegativeSubtraction = new NonNegativeSubtraction(25, 11);
echo $nonNegativeSubtraction->getOperationResult(), PHP_EOL;

/*
140
0
14
*/
