Doctrine
-

````php
// Generate Entities from an Existing Database !!!
php bin/console doctrine:mapping:import AppBundle xml
php bin/console doctrine:mapping:convert -f annotation ./src
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
````
