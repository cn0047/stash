<?php

trait Foo
{
    public function getName()
    {
        echo 'Foo'.PHP_EOL;
    }
}

class Bar
{
    public function getName()
    {
        echo 'Bar'.PHP_EOL;
    }
}

class Boo extends Bar
{
    use Foo;
}

$boo = new Boo;
$boo->getName();
/*
Foo
*/
?>

<?php

trait Foo3
{
    public function getName()
    {
        echo 'Foo3'.PHP_EOL;
    }
}

class Bar3
{
    public function getName()
    {
        echo 'Bar3'.PHP_EOL;
    }
}

class Boo3 extends Bar3
{
    use Foo3 {getName as public;}
}

$boo = new Boo3;
$boo->getName();
/*
Foo3
*/
?>

<?php

trait Foo2
{
    public function getName()
    {
        echo 'Foo2'.PHP_EOL;
    }
}

class Bar2
{
    public function getName()
    {
        echo 'Bar2'.PHP_EOL;
    }
}

class Boo2 extends Bar2
{
    use Foo2  {getName as private;}
}

$boo = new Boo2;
$boo->getName();
/*
PHP Fatal error:  Access level to Foo2::getName() must be public (as in class Bar2)
*/
?>
