<?php

class ClassTest
{
    public static function get()
    {
        echo __CLASS__.PHP_EOL;
    }
}

class ClassTestChild extends ClassTest
{
    public static function get()
    {
        echo __CLASS__.PHP_EOL;
    }

    public static function getParent()
    {
        parent::get();
    }
}

ClassTestChild::get();
ClassTestChild::getParent();
/*
ClassTestChild
ClassTest
*/
