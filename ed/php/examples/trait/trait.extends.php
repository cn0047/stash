<?php

trait Foo
{
}

trait Bar extends Foo
{
}

/*
PHP Parse error:  syntax error, unexpected 'extends' (T_EXTENDS), expecting '{'
*/
