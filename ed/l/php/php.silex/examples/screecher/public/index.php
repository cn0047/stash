<?php

use Doctrine\ORM\Tools\Setup;
use Doctrine\ORM\EntityManager;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

require_once '../vendor/autoload.php';

$app = new Silex\Application();

$app->register(new Silex\Provider\ValidatorServiceProvider());

$app['debug'] = true;
$app['database'] = new PDO("mysql:host=localhost;dbname=mysql", 'root', '');
$entityManager = EntityManager::create(
    [
        'driver' => 'pdo_mysql',
        'dbname' => 'mysql',
        'host' => 'localhost',
        'user' => 'root',
        'password' => null
    ],
    Setup::createAnnotationMetadataConfiguration([__DIR__.'/src'], $app['debug'])
);
$app['entityManager'] = $entityManager;

$app->mount('/apis', require __DIR__.'/../src/Controller/ApiController.php');
$app->mount('/maintainers', require __DIR__.'/../src/Controller/MaintainerController.php');
$app->mount('/spam_maintainer', require __DIR__.'/../src/Controller/SpamMaintainerController.php');

$app->get('/', function() use($app) {
    return 'Hello Igor! Your Silex application is up and running. Woot!';
});

$app->post('/', function (Request $request) {
    $message = $request->get('message');
    return new Response('You said: "' . $message . '". Thank you!', 201);
});

$app->error(function (\Exception $e, $code) {
    return new Response('We are sorry, but something went terribly wrong.');
});

$app->run();