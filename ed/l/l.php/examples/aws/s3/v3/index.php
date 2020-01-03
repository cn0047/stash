<?php

require_once __DIR__ . '/vendor/autoload.php';

use Aws\S3\S3Client;
use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

class Bridge
{
    public $config;
    public $s3;

    public function __construct()
    {
        $configAws = include __DIR__ . '/../../config.php';
        $this->config = $configAws;
        $this->s3 = S3Client::factory([
            'region' => $configAws->region,
            'credentials' => (array)$configAws->credentials,
            'version' => 'latest',
        ]);
    }

    public function fileExists($key)
    {
        try {
            $this->s3->getObject([
                'Bucket' => $this->config->s3->bucket,
                'Key' => $key,
            ]);
            return true;
        } catch (\Aws\S3\Exception\S3Exception $e) {
        }
        return false;
    }

    public function copy($fromKey, $toKey)
    {
        return $this->s3->copyObject([
            'Bucket' => $this->config->s3->bucket,
            'Key' => $toKey,
            'CopySource' => urlencode($this->config->s3->bucket . '/' . $fromKey),
        ]);
    }
}

class Command
{
    /**
     * @example  php index.php upload
     */
    public function upload()
    {
    }

    public function getUrl($configAws)
    {
        $command = $this->s3->getCommand('GetObject', [
            'Bucket' => $configAws->s3->bucket,
            'Key'    => 'test/logo.jpg',
        ]);
        $url = $command->createPresignedUrl('+5 minutes');
        echo "$url \n";
    }

    /**
     * @example php index.php fileExists '000010215/big.mp4'
     * @example php index.php fileExists '00000x/public/photo1.jpg'
     */
    public function fileExists($key)
    {
        var_export((new Bridge())->fileExists($key));
    }

    /**
     * @example php index.php filesExists
     */
    public function filesExists()
    {
        $s3 = new Bridge();

        $a = [
        ];

        foreach ($a as $key) {
            $o = $key;
            $t = str_replace('.jpg', '_thumbnail.jpg', $key);
            $eo = $s3->fileExists($o);
            $et = $s3->fileExists($t);
            //
            if ($eo === true && $et === false) {
                $s3->copy($o, $t);
                echo "[➡️] Copied: $o -> $t" . PHP_EOL;
            }
            if ($eo === false && $et === true) {
                $s3->copy($t, $o);
                echo "[➡️] Copied: $t -> $o" . PHP_EOL;
            }

//            echo PHP_EOL;
//            foreach (['', '_thumbnail.jpg', '_w400_high_thumbnail.jpg', '_w400_low_thumbnail.jpg'] as $suffix) {
//                if (empty($suffix)) {
//                    $name = $key;
//                } else {
//                    $name = str_replace('.jpg', $suffix, $key);
//                }
//                $exists = (string)$b->fileExists($name);
//                $sign = $exists === '1' ? '✅' : '❌';
//                echo "[$sign ] Done: $name \n";
//            }
        }
    }

    /**
     * Downloads picture by public url and save it in temporary file.
     *
     * @param string $url Public url to picture.
     * @param string $file Target file name.
     *
     * @throws \RuntimeException In case when picture download failed.
     *
     * @example php index.php download 'https://s3-eu-west-1.amazonaws.com/w3.stage.zii.bucket/000010215/big.mp4' '/tmp/zii.v.mp4'
     */
    public function download($url, $file)
    {
        ini_set('memory_limit', '16M');

        // Download picture.
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_HEADER, false);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_BINARYTRANSFER, true);
        $response = curl_exec($ch);
        if ($response === false) {
            throw new \RuntimeException(curl_error($ch));
        }
        // Save picture to temporary file.
        // With purpose to avoid image collisions, remove picture that might be left.
        if (file_exists($file)) {
            unlink($file);
        }
        $fp = fopen($file, 'bx');
        fwrite($fp, $response);
        fclose($fp);
    }

    /**
     * Downloads picture by public url and save it in temporary file.
     *
     * @param string $url Public url to picture.
     * @param string $file Target file name.
     *
     * @throws \RuntimeException In case when picture download failed.
     *
     * @example php index.php download2 'https://s3-eu-west-1.amazonaws.com/bucket/000010215/big.mp4' '/tmp/zii.v.mp4'
     */
    public function download2($url, $file)
    {
        ini_set('memory_limit', '16M');

        // Yes, it looks tricky - but it really works!!!
        // Previous implementation was with using curl_* functions,
        // but for huge files we obtained FATAL error about allowed memory size
        // in line: $response = curl_exec($ch);
        // With this implementation we don't have such errors.
        `curl -s -o $file $url`;
    }

    /**
     * @param string $key
     *
     * @example php index.php putObjectAcl '000239685/p/pl.jpg'
     */
    public function putObjectAcl($key)
    {
        $r = $this->s3->putObjectAcl([
            'Bucket' => $this->config->s3->bucket,
            'Key' => $key,
            'ACL' => 'public-read'
        ]);
        var_export($r);
    }

    /**
     * @example php index.php importCsvIntoRabbitMQ /vagrant/p.csv
     */
    public function importCsvIntoRabbitMQ($fromFile) {
        (new Helper())->importCsvIntoRabbitMQ($fromFile);
    }

    /**
     * @example php index.php checkS3FromRabbitMQ
     */
    public function checkS3FromRabbitMQ() {
        (new Helper())->checkS3FromRabbitMQ();
    }

    /**
     * @example php index.php deleteAbandonedFilesFromS3
     */
    public function deleteAbandonedFilesFromS3() {
        (new Helper())->deleteAbandonedFilesFromS3();
    }

    /**
     * @example php index.php recreateMissedFilesOnS3
     */
    public function recreateMissedFilesOnS3() {
        (new Helper())->recreateMissedFilesOnS3();
    }
}

class Helper
{
    private function getCsvRow($csvFile)
    {
        $handle = fopen($csvFile, 'rb');
        while (feof($handle) === false) {
            yield fgetcsv($handle);
        }
        fclose($handle);
    }

    public function importCsvIntoRabbitMQ($fromFile) {
        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $channel->queue_declare('aws_s3', false, true, false, false);

        $i = 0;
        foreach ($this->getCsvRow($fromFile) as $row) {
            $data = array_map('trim', array_filter(str_getcsv($row[0], '|')));
            if (empty($data)) {
                break;
            }
            $body = json_encode(array_combine(
                ['id', 'user', 'type', 'status', 'fileName', 'createdAt'],
                $data
            ));
            $msg = new AMQPMessage($body, ['delivery_mode' => 2] /* make message persistent */);
            $channel->basic_publish($msg, '', 'aws_s3');
            $i++;
            echo "\r [✅] $i";
        }
        echo PHP_EOL;

        $channel->close();
        $connection->close();
    }

    public function checkS3FromRabbitMQ()
    {
        echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $channel->queue_declare('aws_s3', false, true, false, false);

        $s3 = new Bridge();
        $redis = new Redis();
        $redis->connect('127.0.0.1', 6379);

        $callback = function ($msg) use ($s3, $redis) {
            $body = $msg->body;
            $data = json_decode($body, true);
            $user = sprintf('%09s', $data['user']);
            $type = $data['type'] === '201203' ? 'private' : 'public';
            foreach (['', '_thumbnail.jpg', '_w400_high_thumbnail.jpg', '_w400_low_thumbnail.jpg'] as $suffix) {
                if (empty($suffix)) {
                    $name = $data['fileName'];
                } else {
                    $name = str_replace('.jpg', $suffix, $data['fileName']);
                }
                $key = "$user/$type/$name";
                $e = $s3->fileExists($key);
                $exists = (string)$e;
                $sign = $exists === '1' ? '✅' : '❌';
                $redis->incr("$exists.$suffix");
                echo "\n [$sign ] Done: $key";
                if ($e === false) {
                    if (in_array($data['fileName'], [
                        'photo_2017-05-22_08-17-09.jpg',
                        'photo_2017-05-22_11-24-29.jpg',
                    ], true)) {
                        echo '🙈';
                        continue;
                    }
                    if (preg_match('/_w400_....?_thumbnail\.jpg$/', $key)) {
                        echo '🙈2️⃣';
                        continue;
                    }
                    $url = 'http://ec2-52-31-39-186.eu-west-1.compute.amazonaws.com/x?f=' . $data['fileName'];
                    $r = `curl -s $url`;
                    $d = json_decode($r, true);
                    if (json_last_error() !== JSON_ERROR_NONE) {
                        var_export($r);
                        die;
                    }
                    if (count($d) === 1) {
                        if ($d[0]['StatusID'] === '201303' || $d[0]['StatusID'] === '201304') {
                            echo '🚮';
                        } elseif ($d[0]['TypeID'] === '201203') {
                            echo '⬇️';
                        } else {
                            die;
                        }
                    } elseif (count($d) === 2) {
                        if (
                            ($d[0]['StatusID'] === '201303' || $d[0]['StatusID'] === '201304')
                            && ($d[1]['StatusID'] === '201303' || $d[1]['StatusID'] === '201304')
                        ) {
                            echo '🚮 2️⃣';
                        } else {
                            die;
                        }
                    } else {
                        die;
                    }
                }
            }
            $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
        };

        $channel->basic_qos(null, 1, null);
        $channel->basic_consume('aws_s3', '', false, false, false, false, $callback);
        while(count($channel->callbacks)) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }

    public function deleteAbandonedFilesFromS3()
    {
        echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $channel->queue_declare('aws_s3', false, true, false, false);

        $s3 = new Bridge();
        $redis = new Redis();
        $redis->connect('127.0.0.1', 6379);

        $callback = function ($msg) use ($s3, $redis) {
            $body = $msg->body;
            $data = json_decode($body, true);
            $user = sprintf('%09s', $data['user']);
            $type = $data['type'] === '201203' ? 'private' : 'public';
            foreach (['', '_thumbnail.jpg', '_w400_high_thumbnail.jpg', '_w400_low_thumbnail.jpg'] as $suffix) {
                if (empty($suffix)) {
                    $name = $data['fileName'];
                } else {
                    $name = str_replace('.jpg', $suffix, $data['fileName']);
                }
                $key = "$user/$type/$name";
                $exists = $s3->fileExists($key);
                $keyBackup = "backup/$user/$type/$name";
                $existsBackup = $s3->fileExists($keyBackup);
                var_export([$key, $exists, $existsBackup]);
            }
            die;
            $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
        };

        $channel->basic_qos(null, 1, null);
        $channel->basic_consume('aws_s3', '', false, false, false, false, $callback);
        while(count($channel->callbacks)) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }

    public function recreateMissedFilesOnS3()
    {
        echo ' [*] Waiting for messages. To exit press CTRL+C', "\n";

        $connection = new AMQPStreamConnection('localhost', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $channel->queue_declare('aws_s3', false, true, false, false);

        $s3 = new Bridge();
        $redis = new Redis();
        $redis->connect('127.0.0.1', 6379);

        $callback = function ($msg) use ($s3, $redis) {
            $body = $msg->body;
            $data = json_decode($body, true);
            $user = sprintf('%09s', $data['user']);
            $type = $data['type'] === '201203' ? 'private' : 'public';
            // All keys.
            $o = "$user/$type/" . $data['fileName'];
            $t = "$user/$type/" . str_replace('.jpg', '_thumbnail.jpg', $data['fileName']);
            $th = "$user/$type/" . str_replace('.jpg', '_w400_high_thumbnail.jpg', $data['fileName']);
            $tl = "$user/$type/" . str_replace('.jpg', '_w400_low_thumbnail.jpg', $data['fileName']);
            $bo = "backup/$o";
            $bt = "backup/$t";
            // Existence of all keys.
            $eo = (int)$s3->fileExists($o);
            $et = (int)$s3->fileExists($t);
            $eth = (int)$s3->fileExists($th);
            $etl = (int)$s3->fileExists($tl);
            $ebo = (int)$s3->fileExists($bo);
            $ebt = (int)$s3->fileExists($bt);
            //
            if ($eo === 0 || $et === 0 || $eth === 0 || $etl === 0) {
                $s = $eo + $et + $eth + $etl + $ebo + $ebt;
                if ($s === 0) {
                    $sign = '❌';
                } else {
                    $sign = '2️⃣';
                    echo 'Keys: ';
                    var_export([
                        $o => $eo,
                        $t => $et,
                        $th => $eth,
                        $tl => $etl,
                        $bo => $ebo,
                        $bt => $ebt,
                        $bth => $ebth,
                        $btl => $ebtl,
                    ]);
                    die;
                }
                echo " [$sign] SUM = $s";
            }
            $msg->delivery_info['channel']->basic_ack($msg->delivery_info['delivery_tag']);
        };

        $channel->basic_qos(null, 1, null);
        $channel->basic_consume('aws_s3', '', false, false, false, false, $callback);
        while(count($channel->callbacks)) {
            $channel->wait();
        }
        $channel->close();
        $connection->close();
    }
}

$phpSelf = array_shift($argv);
$action = array_shift($argv);
(new Command())->$action(...$argv);


// Upload image to s3.
// Next code works, uncomment when you need it.
// $s3->putObject([
//     'Bucket' => $configAws->s3->bucket,
//     'Key'    => 'my-object',
//     'Body'   => fopen('/home/kovpak/Downloads/images.jpg', 'r'),
//     'ACL'    => 'public-read',
// ]);

// Prints list of objects in bucket.
// Next code works, uncomment when you need it.
// $result = $s3->listObjects(array(
//     'Bucket' => 'w3.zii.bucket',
//     'MaxKeys' => 5
// ));
// var_export(($result));

// $aws = new \Aws\Sdk([
//     'region' => $configAws->region,
//     'version' => 'latest',
//     'credentials' => (array)$configAws->credentials,
//     'signature_version' => 'v4',
//     // 'debug'   => true
// ]);
// $s3 = $aws->createS3();
// $objects = $s3->getIterator(
//     'ListObjects',
//     ['Bucket' => 'w3.zii.bucket', 'Prefix' => '000046059']
// );
// foreach ($objects as $object) {
//     echo $object['Key'] . "\n";
// }
// $s3->putObject([
//    'Bucket' => 'w3.zii.bucket',
//    'Key' => '000046059/public/1.jpg',
//    'SourceFile' => '/home/kovpak/Downloads/images.jpg',
//    'ACL' => 'public-read',
// ]);

