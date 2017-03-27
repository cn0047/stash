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

$callback = function (AMQPMessage $message) {
    printf(' [x] Message received: %s %s', $message->body, PHP_EOL);
    $message->delivery_info['channel']->basic_nack($message->delivery_info['delivery_tag']);
};
$channel->basic_consume('delayed_queue', '', false, true, false, false, $callback);
while(count($channel->callbacks)) {
    $channel->wait();
}
$channel->close();
$connection->close();
