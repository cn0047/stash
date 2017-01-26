<?php

declare(strict_types = 1);

require_once __DIR__ . '/vendor/autoload.php';

use Aws\Sqs\SqsClient;
use Guzzle\Service\Resource\Model;

$config = require __DIR__ . '/../config.php';

$client = SqsClient::factory([
    'secret' => $config->credentials['secret'],
    'key' => $config->credentials['key'],
    'region' => $config->region,
]);

$queue = $client->createQueue(['QueueName' => 'test']);

/** @var Model $sentMessage */
$sentMessage = $client->sendMessage([
    'QueueUrl' => $queue['QueueUrl'],
    'MessageBody' => 'Job ' . date('Y-m-d H:i:s'),
]);

while (true) {
    /** @var Model $receivedMessage */
    $receivedMessage = $client->receiveMessage([
        'QueueUrl' => $queue['QueueUrl'],
        'WaitTimeSeconds' => 1,
    ]);
    /** @var array $messages */
    $messages = $receivedMessage->getPath('Messages');
    if ($messages === null) {
        return;
    } else {
        /** @var Model $message */
        foreach ($messages as $message) {
            echo var_export($message['Body'], true), PHP_EOL;
            $client->deleteMessage([
                'QueueUrl' => $queue['QueueUrl'],
                'ReceiptHandle' => $message['ReceiptHandle'],
            ]);
        }
    }
}
