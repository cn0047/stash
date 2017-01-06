<?php

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

class Command
{
    public function one()
    {
        $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);
        var_export($qbb->getToken());
        var_export($qbb->createChat(203379));
        var_export($qbb->createUser());
        var_export($qbb->getHaveUnreadMessage());
    }

    /**
     * @example php index.php sendChatMessageFromAdmin 282040 'my test'
     */
    public function sendChatMessageFromAdmin($controller, $action, $targetQBUserId, $message)
    {
        $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);
        $result = $qbb->sendChatMessageFromAdmin($targetQBUserId, $message);
        var_export($result);
    }

    /**
     * CSV format:
     * | qbUserId |
     * 
     * @example php index.php sendChatMessageFromAdminToUsersFromCsvFile '/home/kovpak/csv.csv'
     */
    public function sendChatMessageFromAdminToUsersFromCsvFile(
        $controller,
        $action,
        $csvFile,
        $message = 'La semaine est presque finie, bientôt les fêtes🎉 🎉 ! Bon courage 😀'
    ) {
        $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);
        foreach ($this->getCsvRow($csvFile) as $row) {
            $targetQBUserId = preg_replace('( |\|)', '', $row)[0];
            try {
                $result = $qbb->sendChatMessageFromAdmin($targetQBUserId, $message);
                file_put_contents('/home/kovpak/quickblox.chat-message-raw-response.log', var_export(json_encode($result), 1)."\n", FILE_APPEND);
            } catch (\RuntimeException $e) {
                $result = $e->getMessage();
                file_put_contents('/home/kovpak/quickblox.chat-message-raw-response.log', $result."\n", FILE_APPEND);
            }
            echo $targetQBUserId . PHP_EOL;
        }
    }

    /**
     * @example php index.php getChats user_82273 NnMPG8AxFMJpBuCbUQby
     */
    public function getChats($controller, $action, $userId, $qbUserPassword)
    {
        $qbb = new QuickBloxBridge($userId, $qbUserPassword);
        var_export($qbb->getChats());
    }

    /**
     * @example php index.php getAdminChatMessages
     */
    public function getAdminChatMessages()
    {
        $qbb = new QuickBloxBridge('user_80716', 'vDxIEM5I7iaqy4p4e78e');
        $chatId = $qbb->getChatBetweenAdminAndTargetUser(280068)['_id'];
        var_export($qbb->getChatMessages($chatId));
    }

    /**
     * @example php index.php markAdminMessageAsReadForQbUser user_46059 '3elkFCQcPS243QDiTxem' 202636
     */
    public function markAdminMessageAsReadForQbUser($controller, $action, $login, $password, $qbUserId)
    {
        Command::markAdminMessageAsRead($login, $password, $qbUserId);
    }

    private static function markAdminMessageAsRead($login, $password, $qbUserId)
    {
        $msg1 = 'La semaine est presque finie, bientôt les fêtes🎉 🎉 ! Bon courage 😀';
        $msg2 = "Noël est presque là 🎄Meilleurs voeux de la part de toute l'équipe de Ziipr !";
        $qbb = new QuickBloxBridge($login, $password);
        $chatId = $qbb->getChatBetweenAdminAndTargetUser($qbUserId)['_id'];
        $messages = $qbb->getChatMessages($chatId)['items'];
        foreach ($messages as $message) {
            var_export($message);
            if (($message['message'] === $msg1 || $message['message'] === $msg2) && $message['read'] === 0) {
                // $r = $qbb->markMessageAsRead($chatId, $message['_id']);
                // var_dump($r);
                // die;
            }
        }
    }

    private function getCsvRow($csvFile = '/home/kovpak/csv.csv')
    {
        $handle = fopen($csvFile, 'rb');
        while (feof($handle) === false) {
            yield fgetcsv($handle);
        }
        fclose($handle);
    }

    /**
     * CSV format:
     * | id | email | userId | qbUserId | password |
     */
    public function exportUsersWithUnreadMessages()
    {
        $i = 0;
        foreach ($this->getCsvRow() as $row) {
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
    }

    /**
     * CSV format:
     * | qbUserId |
     */
    private function fromvCsvRowToStrWithTargetQBUserId($row)
    {
        return json_encode(preg_replace('( |\|)', '', $row)[0]);
    }

    /**
     * CSV format:
     * | quickBloxId | userId | quickBloxUserId | password |
     */
    private function fromvCsvRowToStrWithUserJson($row)
    {
        return json_encode(array_combine(
            ['_id', 'userId', 'qbUserId', 'password'],
            array_map('trim', array_filter(str_getcsv($row[0], '|')))
        ));
    }

    public function importCsvIntoRabbitMQ()
    {
        $cb = 'fromvCsvRowToStrWithUserJson';

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $durable = true;
        $channel->queue_declare('durable_task_queue', false, $durable, false, false);

        foreach ($this->getCsvRow('/vagrant/csv.csv1') as $row) {
            $body = $this->$cb($row);
            $msg = new AMQPMessage($body, array('delivery_mode' => 2) /* make message persistent */);
            $channel->basic_publish($msg, '', 'durable_task_queue');
            echo " [v] User added to queue: ", $body, "\n";
        }

        $channel->close();
        $connection->close();
    }

    public function sendChatMessageFromAdminToUsersFromRabbitMQ()
    {
        $qbb = new QuickBloxBridge(USER_LOGIN, USER_PASSWORD);

        echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $durable = true;
        $channel->queue_declare('durable_task_queue', false, $durable, false, false);

        $callback = function ($msg) use ($qbb) {
            $message = "Noël est presque là 🎄Meilleurs voeux de la part de toute l'équipe de Ziipr !";
            $targetQBUserId = $msg->body;
            try {
                $result = $qbb->sendChatMessageFromAdmin($targetQBUserId, $message);
                file_put_contents('/home/kovpak/quickblox.chat-message-raw-response.log', var_export(json_encode($result), 1)."\n", FILE_APPEND);
            } catch (\RuntimeException $e) {
                $result = $e->getMessage();
                file_put_contents('/home/kovpak/quickblox.chat-message-raw-response.log', $result."\n", FILE_APPEND);
            }
            $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
            echo " [x] Done user: $targetQBUserId \n";
            usleep(100);
        };

        $channel->basic_qos(null, 1, null);
        $noAck = false;
        $channel->basic_consume('durable_task_queue', '', false, $noAck, false, false, $callback);
        while(count($channel->callbacks)) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }

    public function markAdminMessageAsReadFromRabbitMQ()
    {
        echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $durable = true;
        $channel->queue_declare('durable_task_queue', false, $durable, false, false);

        $callback = function ($msg) {
            $body = $msg->body;
            echo $body . PHP_EOL;
            $data = json_decode($body, true);
            try {
                Command::markAdminMessageAsRead('user_' . $data['userId'], $data['password'], $data['qbUserId']);
            } catch (ChatBetweenAdminAndTargetUserNotFound $e) {
                echo 'ChatBetweenAdminAndTargetUserNotFound ';
            }

            $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
            echo "[✅] Done:  \n";
            usleep(50);
        };

        $channel->basic_qos(null, 1, null);
        $noAck = false;
        $channel->basic_consume('durable_task_queue', '', false, $noAck, false, false, $callback);
        while(count($channel->callbacks)) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }
}
