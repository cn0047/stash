<?php

require_once __DIR__ . '/../config.php';

set_error_handler(function ($code, $description) {
    throw new ErrorException($description, $code);
});

class QuickBloxBridge
{
    private $login;
    private $password;

    public function __construct($login, $password)
    {
        $this->login = $login;
        $this->password = $password;
    }

    public function getToken()
    {
        $body = [
            'application_id' => APPLICATION_ID,
            'auth_key' => AUTH_KEY,
            'nonce' => time(),
            'timestamp' => time(),
            'user' => ['login' => $this->login, 'password' => $this->password]
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

    public function getChats()
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Dialog.json');
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        return json_decode($response, true);
    }

    public function getHaveUnreadMessage()
    {
        foreach ($this->getChats()['items'] as $key => $chat) {
            if ($chat['unread_messages_count'] > 0) {
                return true;
            }
        }
        return false;
    }
}

// $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);
// var_export($qbb->getToken());
// var_export($qbb->createUser());
// var_export($qbb->getHaveUnreadMessage());

$csv = array_map('str_getcsv', file('/home/kovpak/csv.csv'));
$i = 0;
foreach ($csv as list($email, $userId, $qbUserId, $password)) {
    $i++;
    $error = 0;
    $haveUnreadMessage = false;
    try {
        $qbb = new QuickBloxBridge(sprintf('user_%s', trim($userId)), trim($password));
        $haveUnreadMessage = $qbb->getHaveUnreadMessage();
    } catch (Exception $e) {
        $error = 1;
    }
    if ($haveUnreadMessage) {
        printf('| %5d | %50s | %s | %s', $i, trim($email), trim($userId), PHP_EOL);
    }
}
