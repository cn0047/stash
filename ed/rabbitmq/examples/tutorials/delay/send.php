<?php

require_once __DIR__ . '/../vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;
use PhpAmqpLib\Wire\AMQPTable;

$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();
$args = new AMQPTable(['x-delayed-type' => 'fanout']);
$channel->exchange_declare('delayed_exchange', 'x-delayed-message', false, true, false, false, false, $args);
$args = new AMQPTable(['x-dead-letter-exchange' => 'delayed']);
$channel->queue_declare('delayed_queue', false, true, false, false, false, $args);
$channel->queue_bind('delayed_queue', 'delayed_exchange');

$data = 'Hello World at ' . date('Y-m-d H:i:s');
$delay = 7000;
$message = new AMQPMessage($data, ['delivery_mode' => AMQPMessage::DELIVERY_MODE_PERSISTENT]);
$headers = new AMQPTable(['x-delay' => $delay]);
$message->set('application_headers', $headers);
$channel->basic_publish($message, 'delayed_exchange');
printf(' [x] Message sent: %s %s', $data, PHP_EOL);

$channel->close();
$connection->close();
