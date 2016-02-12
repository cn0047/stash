<?php

namespace AppBundle\Component;

use AppBundle\Event\Component;

class Common
{
    public function onMainMethodIsNotFound(Component $event)
    {
        $event->setResult('Processed by Common component.');
    }
}
