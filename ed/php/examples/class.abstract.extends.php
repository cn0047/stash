<?php

abstract class AppException extends Exception
{
}

class MyException extends AppException
{
}

new MyException;
