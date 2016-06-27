<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../config.php';

use Aws\S3\S3Client;

$s3 = S3Client::factory([
    'region'      => $config_aws->region,
    'credentials' => (array) $config_aws->credentials,
    'version' => 'latest',
]);
$s3->putObject([
    'Bucket' => $config_aws->s3->bucket,
    'Key'    => 'my-object',
    'Body'   => fopen('/home/kovpak/Downloads/images.jpg', 'r'),
    'ACL'    => 'public-read',
]);


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
// $url = $s3->getObjectUrl('x.bucket', 'dir/photo-33.jpg', '+10 minutes');
// var_export($url);
// $command = $s3->getCommand('GetObject', [
//     'Bucket' => 'x.bucket',
//     'Key'    => '000046059/private/1.jpg',
// ]);
// $r = $command->createPresignedUrl('+5 minutes');