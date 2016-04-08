<?php

require_once __DIR__ . '/vendor/autoload.php';

$config_aws = (object)[
        'region' => '',
        'credentials' => [
            'key'    => '',
            'secret' => '',
        ],
        'cognito' => [
            'identity_pool_id' => '',
            'identity_name'    => '',
            'token_duration'   => 7200,
            'roles' => [
                'unauth' => '',
                'auth'   => ''
            ]
        ],
        's3' => [
            'bucket' => 'x.bucket'
        ]
];

$aws = new \Aws\Sdk([
    'region' => $config_aws->region,
    'version' => 'latest',
    'credentials' => (array)$config_aws->credentials,
    'signature_version' => 'v4',
    // 'debug'   => true
]);
$s3 = $aws->createS3();

$objects = $s3->getIterator(
    'ListObjects',
    ['Bucket' => 'w3.ziipr.bucket', 'Prefix' => '000017777']
);
foreach ($objects as $object) {
    echo $object['Key'] . "\n";
}

$s3->putObject([
   'Bucket' => 'w3.ziipr.bucket',
   'Key' => '000017777/private/1.jpg',
   'SourceFile' => '/home/kovpak/Downloads/images.jpg',
   'ACL' => 'public-read',
]);

$url = $s3->getObjectUrl('x.bucket', 'dir/photo-33.jpg', '+10 minutes');
var_export($url);

$command = $s3->getCommand('GetObject', [
    'Bucket' => 'x.bucket',
    'Key'    => '000017777/private/1.jpg',
]);
$r = $command->createPresignedUrl('+5 minutes');
