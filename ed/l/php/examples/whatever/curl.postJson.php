<?php

$headers = [
    'Content-Type: application/json',
];
$json = json_encode([
    'code' => 100,
]);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, 'https://realtimelog.herokuapp.com/sddjklskj');
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_POSTFIELDS, $json);
curl_exec($ch);
