<?php

namespace AppBundle\Entity;

use AppBundle\Entity\Log;
use Doctrine\ORM\EntityRepository;
use Doctrine\ORM\Query;

class LogRepository extends EntityRepository
{
    /**
     * Provides ability to select data through using RegExp.
     *
     * @param array $filters Array with key as field name and value as RegExp pattern.
     * @throws \InvalidArgumentException When $filters contains unknown field.
     * @return array Data.
     */
    public function findByFilters($filters)
    {
        $qb = $this->createQueryBuilder('l');
        if (is_array($filters)) {
            $log = new Log();
            foreach ($filters as $field => $pattern) {
                // This is the simplest validation,
                // if entity has property with name received from request - it's ok
                if (property_exists($log, $field)) {
                    // It's hurt to correctly validate RegExp pattern,
                    // that's why we just look that it not empty.
                    if (!empty($pattern)) {
                        $parameterName = 'parameter_'.uniqid();
                        $qb
                            ->andWhere("REGEXP(l.{$field}, :$parameterName) = true")
                            ->setParameter($parameterName, $pattern)
                        ;
                    }
                } else {
                    throw new \InvalidArgumentException("Unknown field: $field.");
                }
            }
        }
        return $qb->getQuery()->getResult(Query::HYDRATE_ARRAY);
    }
}
