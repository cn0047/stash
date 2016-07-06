<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../../config.php';

use Aws\S3\S3Client;

/**
 * @example php index.php getUrl
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

    private function getUrl($config_aws)
    {
        $command = $this->s3->getCommand('GetObject', [
            'Bucket' => $config_aws->s3->bucket,
            // p
            'Key'    => '000046059/public/photo_2016-06-06_11-30-41_thumbnail.jpg',
            // s
            'Key'    => 'test/logo.jpg',
            'Key'    => '000009355/public/photo_2016-07-06_13-32-57.jpg',
        ]);
        printf("S3 image url: %s\n\n", $command->createPresignedUrl('+5 minutes'));
    }

    private function uploadTxt($config_aws)
    {
        $r = $this->s3->putObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => 'test/data.txt.' . time(),
            'Body'   => 'Hello! ' . time(),
        ]);
        var_export($r);
    }

    private function uploadImg($config_aws)
    {
        $r = $this->s3->putObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => 'test/BOND.jpg',
            'SourceFile' => '/home/kovpak/Downloads/b.jpg',
        ]);
        var_export($r);
    }
}

new Command($config_aws, $argv[1]);
