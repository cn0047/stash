<?php
/*
docker exec -it zmq-hw php /gh/ed/php/examples/zeromq/hw/client.php
*/

echo "Connecting to HW server…\n";
$context = new ZMQContext();
$requester = new ZMQSocket($context, ZMQ::SOCKET_REQ);
$requester->connect('tcp://localhost:31256');
$requester->setSockOpt(ZMQ::SOCKOPT_LINGER, 0);

for ($request_nbr = 0; $request_nbr != 5; $request_nbr++) {
    printf ("Sending request %d…\n", $request_nbr);
    $requester->send('Hello');

    $reply = $requester->recv();
    printf ("Received reply %d: [%s]\n", $request_nbr, $reply);
}
