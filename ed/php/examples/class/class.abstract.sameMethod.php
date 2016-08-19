<?php

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
