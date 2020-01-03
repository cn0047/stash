<?php
/*
docker run -it --rm -v $PWD:/gh --name=zmq-hw nphp \
    php /gh/ed/php/examples/zeromq/hw/server.php
*/

echo "Starting HW server…\n";
$context = new ZMQContext(1);
$responder = new ZMQSocket($context, ZMQ::SOCKET_REP);
$responder->bind('tcp://*:31256');

while (true) {
    //  Wait for next request from client
    $request = $responder->recv();
    printf ("Received request: [%s]\n", $request);

    sleep (1);

    //  Send reply back to client
    $responder->send('World');
}
