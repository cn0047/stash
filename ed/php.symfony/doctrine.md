Doctrine
-

````php
// Generate Entities from an Existing Database !!!
php bin/console doctrine:mapping:import AppBundle xml
php bin/console doctrine:mapping:convert annotation ./src
php bin/console doctrine:generate:entities --no-backup AppBundle
php bin/console doctrine:mapping:info

php bin/console doctrine:database:create
php bin/console doctrine:database:drop --force
php bin/console doctrine:generate:entity
php bin/console doctrine:generate:entities AppBundle/Entity/Product
// update db schema from entities
php bin/console doctrine:schema:update --force

// generates all entities in the AppBundle
php bin/console doctrine:generate:entities AppBundle
// generates all entities of bundles in the Acme namespace
php bin/console doctrine:generate:entities Acme

php bin/console list doctrine
php bin/console help doctrine:database:create
php bin/console doctrine:ensure-production-settings --env=prod

php bin/console doctrine:migrations:execute 20160105185037 --down

// Generate CRUD !!!
php bin/console generate:doctrine:crud --entity=AppBundle:EntityName

// in controller
$post=$this->get('doctrine')->getManager()->getRepository('AppBundle:Post')->find($id);

$em->getConnection()->exec('create table tmp (id int)');

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
