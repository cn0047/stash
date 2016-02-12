<?php

namespace AppBundle\Component;

use AppBundle\Event\Component;

class Main
{
    private $dispatcher;

    public function __construct($dispatcher)
    {
        $this->dispatcher = $dispatcher;
    }

    public function __call($method, $arguments)
    {
        $event = new Component();
        $this->dispatcher->dispatch('component.main.method_is_not_found', $event);
        return $event->getResult();
    }
}
