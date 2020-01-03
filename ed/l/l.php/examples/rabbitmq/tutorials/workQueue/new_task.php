<?php

require_once __DIR__ . '/../vendor/autoload.php';
use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$durable = true;
$channel->queue_declare('durable_task_queue', false, $durable, false, false);

$uid = uniqid('', false);
$data = $argv[1] ?? "Hello World! [$uid]";
$msg = new AMQPMessage($data, ['delivery_mode' => 2] /* make message persistent */);

$channel->basic_publish($msg, '', 'durable_task_queue');

echo " [v] Sent $data" . PHP_EOL;

$channel->close();
$connection->close();
