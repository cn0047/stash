<?php

require_once __DIR__ . '/../config.php';

set_error_handler(function ($code, $description) {
    throw new ErrorException($description, $code);
});

class QuickBloxBridge
{
    private $login;
    private $password;
    private $token;

    public function __construct($login, $password)
    {
        $this->login = $login;
        $this->password = $password;
    }

    public function getToken()
    {
        if ($this->token !== null) {
            return $this->token;
        }
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
        $this->token = json_decode($response, true)['session']['token'];
        return $this->token;
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

    public function createChat($userId)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $body = [
            'type' => 3, // private chat
            'type' => 2, // group chat
            'name' => 'chat with ziipr admin',
            'occupants_ids' => $userId,
        ];
        $post_body = http_build_query($body);
        $data = json_encode($body);
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Dialog.json');
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_POST, true);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        var_export($response);
    }
}

// $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);
// var_export($qbb->getToken());
// var_export($qbb->createChat(203379));
// var_export($qbb->createUser());
// var_export($qbb->getHaveUnreadMessage());

function getRows() {
    $handle = fopen('/home/kovpak/csv.csv', 'rb');
    while (feof($handle) === false) {
        yield fgetcsv($handle);
    }
    fclose($handle);
}
$i = 0;
foreach (getRows() as $row) {
    list(, $email, $userId, $qbUserId, $password) = str_getcsv($row[0], '|');
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
        file_put_contents('/home/kovpak/r.txt', trim($email)."\n", FILE_APPEND);
    }
}
