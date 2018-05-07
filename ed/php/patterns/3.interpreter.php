<?php
/**
 * Interpreter
 *
 * @category Behaviour
 */

interface MathExpression
{
    public function evaluate(array $values);
}

class Literal implements MathExpression
{
    private $literal;

    public function __construct($literal)
    {
        $this->literal = $literal;
    }

    public function evaluate(array $values = [])
    {
        return $this->literal;
    }
}

class Variable implements MathExpression
{
    private $variable;

    public function __construct($variable)
    {
        $this->variable = $variable;
    }

    public function evaluate(array $values)
    {
        return $values[$this->variable];
    }
}

class Sum implements MathExpression
{
    private $a;
    private $b;

    public function __construct(MathExpression $a, MathExpression $b)
    {
        $this->a = $a;
        $this->b = $b;
    }

    public function evaluate(array $values = [])
    {
        return $this->a->evaluate($values) + $this->b->evaluate($values);
    }
}

class Product implements MathExpression
{
    private $a;
    private $b;

    public function __construct(MathExpression $a, MathExpression $b)
    {
        $this->a = $a;
        $this->b = $b;
    }

    public function evaluate(array $values = [])
    {
        return $this->a->evaluate($values) * $this->b->evaluate($values);
    }
}

$o = new Sum(new Literal(2), new Literal(3));
echo $o->evaluate().PHP_EOL;

$o = new Sum(new Literal(4), new Variable('a'));
echo $o->evaluate(['a' => 5]).PHP_EOL;

$o = new Product(new Literal(2), new Literal(3));
echo $o->evaluate().PHP_EOL;

$o = new Product(new Literal(4), new Variable('b'));
echo $o->evaluate(['b' => 5]).PHP_EOL;

$o = new Product(
    new Sum(new Literal(2), new Literal(3)),
    new Sum(new Literal(4), new Variable('c'))
);
echo $o->evaluate(['c' => 5]).PHP_EOL;

/*
5
9
6
20
45
*/
