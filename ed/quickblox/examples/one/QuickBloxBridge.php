<?php

class TokenException1 extends RuntimeException {};
class TokenException2 extends RuntimeException {};
class ItIs502 extends RuntimeException {};

class QuickBloxBridge
{
    private $login;
    private $password;
    private $token;
    private $tokenCacheFile = '/tmp/qb.admin.token';

    public function __construct($login, $password)
    {
        $this->login = $login;
        $this->password = $password;
        if ($this->isAdmin()) {
            $this->initTokenFromCache();
        }
        printf('%s %s ', $this->login, $this->password);
    }

    private function isAdmin()
    {
        return $this->login === 'l' && $this->password === 'p';
    }

    private function initTokenFromCache()
    {
        if (file_exists($this->tokenCacheFile)) {
            $this->token = file_get_contents($this->tokenCacheFile);
        }
    }

    private function saveTokenToCache($token)
    {
        file_put_contents($this->tokenCacheFile, $token);
    }

    public function getTokenValue()
    {
        return $this->token;
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
        curl_setopt($curl, CURLOPT_HEADER, true);
        $response = curl_exec($curl);
        $responseHttpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
        if ($responseHttpCode === 502) {
            throw new ItIs502();
        }
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            // throw new \TokenException1('ERROR-1');
            // Try remove from response headers info.
            print('💊 ');
            $response = substr($response, strpos($response, '{"'));
            $payload = json_decode($response, true);
            if (json_last_error() !== JSON_ERROR_NONE) {
                var_dump($response);
                echo PHP_EOL, PHP_EOL, "================================", PHP_EOL, PHP_EOL;
                var_dump($responseHttpCode);
                echo PHP_EOL, PHP_EOL, "================================", PHP_EOL, PHP_EOL;
                throw new \RuntimeException('ERROR-1');
            }
        }
        if (!isset($payload['session'])) {
            file_put_contents('/tmp/debug.tmp', var_export([$this->login, $this->password, $payload], 1)."\n", FILE_APPEND); /// tail -f /tmp/debug.tmp
            throw new \TokenException2('ERROR-2');
        }
        $this->token = $payload['session']['token'];
        if ($this->isAdmin()) {
            $this->saveTokenToCache($this->token);
        }
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
            throw new ChatBetweenAdminAndTargetUserNotFound($chatName);
        }
        return $chats[0];
    }

    public function getChats($name = '', $type = 0, $limit = 100, $offset = 0)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $get = [
            'limit' => $limit,
            'skip' => ($offset > 0) ? $offset : 0,
        ];
        if ($type > 0) {
            $get['type'] = $type;
        }
        if ($name !== '') {
            $get['name'] = $name;
            $get['sort_desc'] = 'created_at';
        }
        $query = http_build_query($get);
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Dialog.json?' . $query);
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($curl, CURLOPT_HEADER, true);
        $response = curl_exec($curl);
        $responseHttpCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
        // If token expired - we need re-try request.
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            // var_dump($query);
            // var_dump($responseHttpCode);
            // var_dump($response);
            // throw new \RuntimeException('ERROR-3');
            echo "💉 ";
            $response = substr($response, strpos($response, '{"'));
            $payload = json_decode($response, true);
            if (json_last_error() !== JSON_ERROR_NONE) {
                var_dump($response);
                echo PHP_EOL, PHP_EOL, "================================", PHP_EOL, PHP_EOL;
                var_dump($responseHttpCode);
                echo PHP_EOL, PHP_EOL, "================================", PHP_EOL, PHP_EOL;
                throw new \RuntimeException('ERROR-1');
            }
        }
        return $payload;
    }

    public function deleteChat($id, $forse = false)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $ch = curl_init();
        $query = $forse ? 'forse=1' : '';
        curl_setopt($ch, CURLOPT_URL, QB_API_ENDPOINT . "/chat/Dialog/$id.json&$query");
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'DELETE');
        curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_HEADER, true);
        $response = curl_exec($ch);
        $code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        var_dump($code, $response);
        return $code;
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

    public function getChatMessages($chatDialogId)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $query = http_build_query([
            'chat_dialog_id' => $chatDialogId,
        ]);
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_URL, QB_API_ENDPOINT . '/chat/Message.json?' . $query);
        curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($curl);
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException($response);
        }
        return $payload;
    }

    public function markMessageAsRead($chatDialogId, $messageId)
    {
        $headers = [
            'QB-Token: '. $this->getToken(),
            'Content-Type: application/json',
        ];
        $put = [
            'chat_dialog_id' => $chatDialogId,
            'read' => 1,
        ];
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, QB_API_ENDPOINT . "/chat/Message/$messageId.json");
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'PUT');
        curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
        curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($put));
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        $response = curl_exec($ch);
        return $response;
        if ($response === false) {
            throw new RuntimeException(curl_error($ch));
        }
        $payload = json_decode($response, true);
        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new \RuntimeException($response);
        }
        return $payload;
    }
}
