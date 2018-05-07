<?php
/**
 * Proxy
 *
 * @category Structural
 *
 * @see https://github.com/cn007b/my/blob/master/ed/nodejs/patterns/2.proxy.simple.js
 */

interface IMath
{
    function Add($x, $y);
    function Sub($x, $y);
    function Mul($x, $y);
    function Div($x, $y);
}

class Math implements IMath
{
    public function __construct()
    {
        echo 'Create object Math. Wait...'.PHP_EOL;
    }

    public function  Add($x, $y)
    {
        return $x + $y;
    }

    public function  Sub($x, $y)
    {
        return $x - $y;
    }

    public function  Mul($x, $y)
    {
        return $x * $y;
    }

    public function  Div($x, $y)
    {
        return $x / $y;
    }
}

class MathProxy implements IMath
{
    protected $math;

    public function __construct()
    {
        $this->math = null;
    }

    public function Add($x, $y)
    {
        return $x + $y;
    }

    public function  Sub($x, $y)
    {
        return $x - $y;
    }

    public function Mul($x, $y)
    {
        if ($this->math == null) {
            $this->math = new Math();
        }
        return $this->math->Mul($x, $y);
    }

    public function Div($x, $y)
    {
        if ($this->math == null) {
            $this->math = new Math();
        }
        return $this->math->Div($x, $y);
    }
}

$p = new MathProxy;
echo '4 + 2 = '.$p->Add(4, 2).PHP_EOL;
echo '4 - 2 = '.$p->Sub(4, 2).PHP_EOL;
echo '4 * 2 = '.$p->Mul(4, 2).PHP_EOL;
echo '4 / 2 = '.$p->Div(4, 2).PHP_EOL;

/*
4 + 2 = 6
4 - 2 = 2
Create object Math. Wait...
4 * 2 = 8
4 / 2 = 2
*/
