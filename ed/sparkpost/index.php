<?php

require 'vendor/autoload.php';

use SparkPost\SparkPost;
use GuzzleHttp\Client;
use Http\Adapter\Guzzle6\Client as GuzzleAdapter;

$httpClient = new GuzzleAdapter(new Client());
$sparky = new SparkPost($httpClient, ['key'=>'']);
$promise = $sparky->transmissions->post([
    'content' => [
        'from' => ['name' => 'SparkPost Team', 'email' => 'from@sparkpostbox.com'],
        'subject' => 'First Mailing From PHP',
        'html' => '<html><body><h1>Congratulations, {{name}}!</h1><p>You just sent your very first mailing!</p></body></html>',
        'text' => 'Congratulations, {{name}}!! You just sent your very first mailing!',
    ],
    'substitution_data' => ['name' => 'James Bond'],
    'recipients' => [
        ['address' => ['name' => 'James', 'email' => 'vladimir.kovpak@dm-companies.com']],
    ],
]);
$res = $promise->wait();
var_export($res);
