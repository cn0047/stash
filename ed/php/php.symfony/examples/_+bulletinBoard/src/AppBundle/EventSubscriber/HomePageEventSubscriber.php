<?php

namespace AppBundle\EventSubscriber;

class HomePageEventSubscriber
{
    public function onCustomEvent($event)
    {
        var_dump($event->getCode());
    }
}