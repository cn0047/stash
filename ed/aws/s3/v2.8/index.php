<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../config.php';

use Aws\S3\S3Client;

/**
 * Get s3 image url.
 */
// S3 client.
$s3 = S3Client::factory([
    'region'      => $config_aws->region,
    'credentials' => (array) $config_aws->credentials,
    'version' => 'latest',
]);
// Get url to s3 picture.
$command = $s3->getCommand('GetObject', [
    'Bucket' => $config_aws->s3->bucket,
    'Key'    => '000046059/public/photo_2016-06-06_11-30-41_thumbnail.jpg',
]);
printf("S3 image url: %s\n\n", $command->createPresignedUrl('+5 minutes'));

$r = $s3->putObject([
   'Bucket' => $config_aws->s3->bucket,
   'Key'    => '000046059/public/phpTest.jpg',
   'SourceFile' => '/home/kovpak/Downloads/ZiiprDefault.jpg',
   'ACL' => 'public-read',
]);
