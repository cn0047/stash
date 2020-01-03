DQL (Doctrine Query Language)
-

````sql
SELECT u FROM MyProject\Model\User u WHERE u.age > 20
````

````php
<?php
// $em instanceof EntityManager
// example1: passing a DQL string
$q = $em->createQuery('select u from MyProject\Model\User u');
// example2: usin setDql
$q = $em->createQuery();
$q->setDql('select u from MyProject\Model\User u');
````

````php
$query = $em
    ->getRepository('LetterBundle:BurstSeed')
    ->createQueryBuilder('bs')
    ->select('bs.burstId')
    ->where(implode(' OR ', $clauses))
;
$r = $query->getQuery()->getArrayResult();

$query = $em
    ->createQueryBuilder()
    ->select('bs')
    ->select("
        bs.burstId, COUNT(bs.burstId) AS numSeeds,
        b.status, b.numTotal, b.numSent,
        bs.sendResult,
        SUM(CASE WHEN (bs.sendResult = 'ok_read') THEN 1 ELSE 0 END) AS numInbox,
        SUM(CASE WHEN (bs.sendResult = 'ok_spam_folder') THEN 1 ELSE 0 END) AS numSpamBox
    ")
    ->from('LetterBundle:BurstSeed', 'bs')
    ->leftJoin('BurstBundle:Burst', 'b', \Doctrine\ORM\Query\Expr\Join::WITH, 'bs.burstId = b.id')
    ->andWhere('b.id IN (:bursts)')
    ->setParameter('bursts', $bursts)
    ->groupBy('bs.burstId')
    ->orderBy('numSpamBox', 'DESC')
    ->addOrderBy('b.numTotal', 'DESC')
;
````
