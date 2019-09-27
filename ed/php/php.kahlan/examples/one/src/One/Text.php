<?php

namespace One;

class Text
{
    public function get()
    {
        return (new Foo())->bar();
    }

    public function time()
    {
        return time();
    }

    public function dt()
    {
        return \DateTime::createFromFormat('Y-m-d', '2001-01-01');
    }
}
