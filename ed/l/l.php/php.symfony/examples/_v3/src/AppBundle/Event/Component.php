<?php

namespace AppBundle\Event;

use Symfony\Component\EventDispatcher\Event;

class Component extends Event
{
    private $result;

    public function setResult($result)
    {
        $this->result = $result;
    }

    public function getResult()
    {
        return $this->result;
    }
}