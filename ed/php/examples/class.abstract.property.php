<?php

abstract class CAppException
{
    private $code = 404;
}

class MyException extends CAppException
{
    public $code = 200;
}

$o = new MyException;
var_export($o->code);

/*
200
*/
