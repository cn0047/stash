Decorator
-
(also known as Wrapper, an alternative naming shared with the Adapter pattern)

Behaviour.

Is a design pattern that allows behavior to be added to an individual object, either statically or dynamically,
without affecting the behavior of other objects from the same class.

````php
<?php

abstract class Component
{
    abstract public function operation();
}

class ConcreteComponent extends Component
{
    public function operation()
    {
        return 'I am component';
    }
}

abstract class Decorator extends Component
{
    protected $component = null;

    public function __construct(Component $component)
    {
        $this->component = $component;
    }

    public function operation()
    {
        return $this->component->operation();
    }
}

class ConcreteDecoratorA extends Decorator
{
    public function operation()
    {
        return '<a>' . parent::operation() . '</a>';
    }
}

class ConcreteDecoratorS extends Decorator
{
    public function operation()
    {
        return '<strong>' . parent::operation() . '</strong>';
    }
}

$element = new ConcreteComponent();
$extendedElement = new ConcreteDecoratorA($element);
$superExtendedElement = new ConcreteDecoratorS($extendedElement);
print $superExtendedElement->operation();
/*
<strong><a>I am component</a></strong>
*/
````
