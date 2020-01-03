<?php

class Father
{
    final public function getSurname()
    {
    }
}

class Son extends Father
{
    public function getSurname()
    {
    }
}

new Son;

/*
PHP Fatal error:  Cannot override final method Father::getSurname()
*/
