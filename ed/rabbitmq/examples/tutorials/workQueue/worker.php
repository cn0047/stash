<?php

require_once __DIR__ . '/../vendor/autoload.php';
use PhpAmqpLib\Connection\AMQPStreamConnection;


$connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
$channel = $connection->channel();

$durable = true;
$channel->queue_declare('durable_task_queue', false, $durable, false, false);

echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

$callback = function($msg) {
    echo " [x] Received ", $msg->body, "\n";
    sleep(substr_count($msg->body, '.'));
    echo " [x] Done", "\n";
    $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
};
// This tells RabbitMQ not to give more than one message to a worker at a time.
// Or, in other words, don't dispatch a new message to a worker
// until it has processed and acknowledged the previous one.
$channel->basic_qos(null, 1, null);
$noAck = false;
$channel->basic_consume('durable_task_queue', '', false, $noAck, false, false, $callback);

while(count($channel->callbacks)) {
    $channel->wait();
}

$channel->close();
$connection->close();
