<?php

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

    public function getChatBetweenAdminAndTargetUser($targetUserId)
    {
        $chatName = "chat between ziipr admin and $targetUserId";
        $chats = $this->getChats($chatName)['items'];
        if (!isset($chats[0])) {
            throw new \DomainException('Chat between admin and target user not found.');
        }
        return $chats[0];
    }

    public function getChats($name = '')
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $get = [];
        if ($name !== '') {
            $get = [
                'name' => $name,
                'sort_desc' => 'created_at',
            ];
        }
        $query = http_build_query($get);
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Dialog.json?' . $query);
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException($response);
        }
        return $payload;
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

    public function createChatBetweenAdminAndTargetUser($targetUserId)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $body = [
            'type' => 3, // private chat
            'type' => 2, // group chat
            'name' => 'chat with ziipr admin',
            'name' => "chat between ziipr admin and $targetUserId",
            'occupants_ids' => $targetUserId,
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
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException(__FILE__ . __LINE__);
        }
        return $payload;
    }

    public function sendChatMessageFromAdmin($targetUserId, $message)
    {
        try {
            $chatDialogId = $this->getChatBetweenAdminAndTargetUser($targetUserId)['_id'];
            $oldDialog = true;
        } catch (\DomainException $e) {
            $chatDialogId = $this->createChatBetweenAdminAndTargetUser($targetUserId)['_id'];
            $oldDialog = false;
        }
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $data = json_encode([
            'chat_dialog_id' => $chatDialogId,
            'recipient_id' => $targetUserId,
            'message' => $message,
        ]);
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Message.json');
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_POST, true);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException($response);
        }
        $payload['oldDialog'] = $oldDialog;
        return $payload;
    }
}
