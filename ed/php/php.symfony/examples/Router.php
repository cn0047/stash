<?php

require './vendor/autoload.php';

use Symfony\Component\Routing\Matcher\UrlMatcher;
use Symfony\Component\Routing\RequestContext;
use Symfony\Component\Routing\Route;
use Symfony\Component\Routing\RouteCollection;

$route = new Route(
    '/movie/english/{slug}/edit/{id}/{title}',
    array('controller' => 'MyController')
);
$routes = new RouteCollection();
$routes->add('route_name', $route);

$context = new RequestContext();
$matcher = new UrlMatcher($routes, $context);
$parameters = $matcher->match('/movie/english/scorpion/edit/125/E01E05');

var_dump($parameters);
