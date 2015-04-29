<?php

trait Foo
{
    private $name = 'Foo';
}

class Bar
{
    use Foo;

    private $name = 'Bar';

    public function getName()
    {
        echo $this->name;
    }
}

$bar = new Bar();
$bar->getName();

/*
PHP Fatal error:  Bar and Foo define the same property ($name) in the composition of Bar. However, the definition differs and is considered incompatible. Class was composed in
*/
