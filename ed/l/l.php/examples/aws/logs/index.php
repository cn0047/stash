<?php

declare(strict_types = 1);

require_once __DIR__ . '/vendor/autoload.php';

use Aws\CloudWatchLogs\CloudWatchLogsClient;

$config = require __DIR__ . '/../config.php';
date_default_timezone_set('Europe/Dublin');

$client = CloudWatchLogsClient::factory(array(
    'secret' => $config->credentials['secret'],
    'key' => $config->credentials['key'],
    'region' => $config->region,
));
$result = $client->putLogEvents([
    'logGroupName' => 'test',
    'logStreamName' => 'try',
    'logEvents' => [
        [
            'message' => 'This is a test 1',
            'timestamp' => round(microtime(true) * 1000),
        ]
    ]
]);
