<?php

namespace AppBundle\EventListener;

class MyDoctrineListener
{
    public function onPostPersist($event)
    {
        var_dump($event);
    }
    public function onCustom($event)
    {
        var_dump($event);
    }
}
