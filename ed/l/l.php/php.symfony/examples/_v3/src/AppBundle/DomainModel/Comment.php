<?php

namespace AppBundle\DomainModel;

use \Doctrine\ORM\Query;
use Doctrine\ORM\EntityManager;

class Comment
{
    /** @var EntityManager */
    private $readEm;

    /** @var EntityManager */
    private $writeEm;

    public function __construct(EntityManager $readEm, EntityManager $writeEm)
    {
        $this->readEm = $readEm;
        $this->writeEm = $writeEm;
    }

    public function get()
    {
        $r = $this
            ->readEm
            ->getRepository('AppBundle:Comment')
            ->createQueryBuilder('c')
            ->select('c')
            ->getQuery()
            ->getResult(Query::HYDRATE_ARRAY)
        ;
        return $r;
    }
}
