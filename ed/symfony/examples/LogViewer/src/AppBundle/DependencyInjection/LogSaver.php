<?php

namespace AppBundle\DependencyInjection;

use AppBundle\Entity\Log;

class LogSaver
{
    private $em;
    private $batchSize = 0;
    private $batchThreshold = 10;

    public function __construct($em)
    {
        $this->em = $em;
    }

    public function save(array $args)
    {
        $log = new Log();
        $log->init($args);
        $this->em->persist($log);
        $this->batchSize++;
        if (($this->batchSize % $this->batchThreshold) === 0) {
            $this->flush();
        }
    }

    public function flush()
    {
        $this->em->flush();
        $this->em->clear();
        $this->batchSize = 0;
    }
}
