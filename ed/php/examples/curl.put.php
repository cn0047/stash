<?php

$headers = [
    'X-AUTH-TOKEN: R7fFeZSETGZgm7Irb2zY640j94tXLCrQ1uXcgffIx75HSKI9\/cZIvy8XE9y1IEHbV1\/qV0Je6+dpl4QAbRNAdw==',
    'Content-Type: application/json',
];
$put = [
    'settings' => [
        'unit_system' => 'metric',
        'client_lang' => 'fr',
        'mood_alert' => 1,
        'two_column_summary_layout' => 0,
        'ntf_mask' => 255,
    ],
];
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, 'http://api.ziipr.dev/v1/users/69074/settings');
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'PUT');
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($put));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HEADER, true);
$response = curl_exec($ch);
$responseHttpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
if ($response === false) {
    throw new RuntimeException(curl_error($ch));
}
var_export($response);
