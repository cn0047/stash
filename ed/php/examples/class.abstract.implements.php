<?php

interface IAppException
{
}

abstract class CAppException extends Exception implements IAppException
{
}

class MyException extends CAppException
{
}

new MyException;
