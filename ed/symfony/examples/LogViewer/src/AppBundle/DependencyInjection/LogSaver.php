<?php

namespace AppBundle\DependencyInjection;

use AppBundle\Entity\Log;

/**
 * Class that provides methods to save logs at db.
 */
class LogSaver
{
    private $em;
    private $batchSize = 0;
    private $batchThreshold = 10;

    public function __construct($em)
    {
        $this->em = $em;
    }

    /**
     * Save log at db.
     *
     * Add entity to batch,
     * and save all entities when threshold will exceeded.
     * @param array $args Parameters that will setted as entity properties.
     */
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

    /**
     * Save all entities from batch at db.
     */
    public function flush()
    {
        $this->em->flush();
        $this->em->clear();
        $this->batchSize = 0;
    }
}
