<?php

$post = [
    'username' => 'user1',
    'password' => 'passuser1',
    'gender'   => 1,
];
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, 'http://www.domain.com');
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($post));
$response = curl_exec($ch);
var_export($response);
