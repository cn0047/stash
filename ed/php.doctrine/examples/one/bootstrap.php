<?php

require_once "vendor/autoload.php";

use Doctrine\ORM\Tools\Setup;
use Doctrine\ORM\EntityManager;

$isDevMode = true;
$config = Setup::createAnnotationMetadataConfiguration(array(__DIR__."/src"), $isDevMode);
$conn = array(
    'driver' => 'pdo_sqlite',
    'path' => '/tmp/db.sqlite',
);
$entityManager = EntityManager::create($conn, $config);
