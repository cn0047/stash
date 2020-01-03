<?php

$user = "mail@gmail.com";
$pass = "pwd";

require __DIR__.'../../../../vendor/autoload.php';

$server = new \Fetch\Server('imap.gmail.com', 993);
$server->setAuthentication($user, $pass);

$messages = $server->getMessages();
/** @var $message \Fetch\Message */
foreach ($messages as $message) {
    $b = $message->getMessageBody(true);
    echo "Subject: {$message->getSubject()}\nBody: $b\n";
}
