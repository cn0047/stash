<?php

require './vendor/autoload.php';

use Symfony\Component\HttpFoundation\JsonResponse;

$response = new JsonResponse();
$response->setData(array(
    'data' => 123
));

$response = new JsonResponse(array(
    'result' => 'error',
    'message' => 'Encrypt is invalid or missing'
));

$response->send();
