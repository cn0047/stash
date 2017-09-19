<?php

$address  = "localhost";
$port=12345;

//OPEN SERVICE
$server = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
socket_set_option($server, SOL_SOCKET, SO_REUSEADDR, 1);
socket_bind($server, $address, $port);
socket_listen($server);

//HANDSHAKE
$client = socket_accept($server);
$message = socket_read($client, 5000);
$matches = array();
preg_match('#Sec-WebSocket-Key: (.*)\r\n#', $message, $matches);
$new_key = base64_encode(pack('H*', sha1($matches[1] . '258EAFA5-E914-47DA-95CA-C5AB0DC85B11')));
$new_message = "HTTP/1.1 101 Switching Protocols\r\n";
$new_message .= "Upgrade: websocket\r\n";
$new_message .= "Connection: Upgrade\r\n";
$new_message .= "Sec-WebSocket-Version: 13\r\n";
$new_message .= "Sec-WebSocket-Accept: " . $new_key . "\r\n\r\n";
socket_write($client, $new_message, strlen($new_message));

//LOOP TO PASS UPDATES

$filename = "/tmp/ws.txt";
$lastmodif = 0;

while (true) {
    sleep(1);
    $currentmodif = filemtime($filename);
    clearstatcache();

    if ($currentmodif > $lastmodif) {
      $lastmodif = $currentmodif;
      $content = file_get_contents($filename);
      $message = chr(129) . chr(strlen($content)) ."". $content;
      socket_write($client, $message);
    }
}
