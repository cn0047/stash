<?php

abstract class AppException extends Exception
{
    const CODE = 404;
}

class MyException extends AppException
{
    const CODE = 200;
}

echo MyException::CODE;

/*
200
*/
