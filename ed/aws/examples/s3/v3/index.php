<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../../config.php';

use Aws\S3\S3Client;

/**
 * @example  php index.php upload
 */
class Command
{
    public $config;
    public $s3;

    public function __construct($configAws)
    {
        $this->config = $configAws;
        $this->s3 = S3Client::factory([
            'region' => $configAws->region,
            'credentials' => (array)$configAws->credentials,
            'version' => 'latest',
        ]);
    }

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
     * @example php index.php fileExists '00000x/public/photo1.jpg'
     */
    public function fileExists($key)
    {
        try {
            $o = $this->s3->getObject([
                'Bucket' => $this->config->s3->bucket,
                'Key' => $key,
            ]);
            $r = $o;
        } catch (\Aws\S3\Exception\S3Exception $e) {
            $r = false;
        }
        var_export($r);
    }
}

$phpSelf = array_shift($argv);
$action = array_shift($argv);
(new Command($config_aws))->$action(...$argv);


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
//     'Bucket' => 'w3.ziipr.bucket',
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
//     ['Bucket' => 'w3.ziipr.bucket', 'Prefix' => '000046059']
// );
// foreach ($objects as $object) {
//     echo $object['Key'] . "\n";
// }
// $s3->putObject([
//    'Bucket' => 'w3.ziipr.bucket',
//    'Key' => '000046059/public/1.jpg',
//    'SourceFile' => '/home/kovpak/Downloads/images.jpg',
//    'ACL' => 'public-read',
// ]);

