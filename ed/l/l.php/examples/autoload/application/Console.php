<?php

namespace application;

use tools\Writer;

class Console
{
    private $writer;

    function __construct()
    {
        $this->writer = new Writer();
    }
}