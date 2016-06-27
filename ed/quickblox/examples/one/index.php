<?php

define('USER_LOGIN', '');
define('USER_PASSWORD', '');

define('APPLICATION_ID', '');
define('AUTH_KEY', '');
define('AUTH_SECRET', '');

define('QB_API_ENDPOINT', '');

class QuickBloxBridge
{
    public function getToken()
    {
        $body = [
            'application_id' => APPLICATION_ID,
            'auth_key' => AUTH_KEY,
            'nonce' => time(),
            'timestamp' => time(),
            'user' => ['login' => USER_LOGIN, 'password' => USER_PASSWORD]
        ];
        $built_query = urldecode(http_build_query($body));
        $signature = hash_hmac( 'sha1', $built_query , AUTH_SECRET);
        $body['signature'] = $signature;
        $post_body = http_build_query($body);

        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/session.json');
        curl_setopt($curl, CURLOPT_POST, true);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $post_body);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        return json_decode($response, true)['session']['token'];
    }

    public function createUser()
    {
        $id = time();
        $body = json_encode([
            'user' => [
                'login' => "test_$id",
                'password' => '12345678',
                'external_user_id' => $id
            ]
        ]);
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $post_body = $body;

        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/users.json');
        curl_setopt($curl, CURLOPT_POST, true);
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $post_body);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        return $response;
    }
}

$qbb = new QuickBloxBridge();
// $token = $qbb->getToken();

var_export($qbb->createUser());
