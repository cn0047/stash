<?php

namespace AppBundle\EventListener;

use Symfony\Component\Console\Event\ConsoleCommandEvent;
use Symfony\Component\HttpKernel\Event\FilterControllerEvent;
use Symfony\Component\HttpKernel\Event\FilterResponseEvent;

class MySymfonyListener
{
    public function onKernelController(FilterControllerEvent $event)
    {
        // Some stuff here...
    }

    public function onKernelResponse(FilterResponseEvent $event)
    {
        // Some stuff here...
    }

    public function onConsoleCommand(ConsoleCommandEvent $event)
    {
        // Some stuff here...
    }
}
