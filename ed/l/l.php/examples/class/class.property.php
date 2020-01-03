<?php

class Error
{
    private $privateError;
    protected $protectedError;
    public $publicError;
}

class MyError extends Error
{
    private $protectedError;
}

/*
PHP Fatal error:  Access level to MyError::$protectedError must be protected (as in class Error) or weaker
*/
