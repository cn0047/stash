<?php

abstract class Foo
{
    abstract public function bar(array $params = null);
}

class Moo extends Foo
{
    public function bar(array $params)
    {
        echo count($params);
    }
}

$c = new Moo;
$c->bar(array(5));

/*
PHP Fatal error:  Declaration of Moo::bar() must be compatible with Foo::bar(array $params = NULL)
*/

class AppException extends Exception
{
}

abstract class Error
{
    abstract public function cast(Exception $item);
}

class MyError extends Error
{
    public function cast(AppException $item)
    {
    }
}

/*
PHP Fatal error:  Declaration of MyError::cast() must be compatible with Error::cast(Exception $item)
*/
