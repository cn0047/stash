<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../../config.php';

use Aws\S3\S3Client;

/**
 * @example  php index.php upload
 */
class Command
{
    private $s3;

    public function __construct($config_aws, $commandName)
    {
        $this->s3 = S3Client::factory([
            'region'      => $config_aws->region,
            'credentials' => (array) $config_aws->credentials,
            'version' => 'latest',
        ]);
        $this->$commandName($config_aws);
    }

    private function upload()
    {
    }

    private function getUrl($config_aws)
    {
        $command = $this->s3->getCommand('GetObject', [
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => 'test/logo.jpg',
        ]);
        $url = $command->createPresignedUrl('+5 minutes');
        echo "$url \n";
    }
}

new Command($config_aws, $argv[1]);


// Upload image to s3.
// Next code works, uncomment when you need it.
// $s3->putObject([
//     'Bucket' => $config_aws->s3->bucket,
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
//     'region' => $config_aws->region,
//     'version' => 'latest',
//     'credentials' => (array)$config_aws->credentials,
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
//    'Key' => '000046059/private/1.jpg',
//    'SourceFile' => '/home/kovpak/Downloads/images.jpg',
//    'ACL' => 'public-read',
// ]);

