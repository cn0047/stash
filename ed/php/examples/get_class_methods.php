<?php

class foo
{
    public function __construct()
    {
    }

    private function bar()
    {
    }

    protected function boo()
    {
    }
}

var_export(get_class_methods('foo'));

/*
array (
  0 => '__construct',
)
*/
