<?php

require_once __DIR__ . '/vendor/autoload.php';
require_once __DIR__ . '/../../config.php';

use Aws\S3\S3Client;
use Aws\S3\Exception\NoSuchKeyException;

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

    private function getUrl($config_aws, $key = null)
    {
        if ($key === null) {
            // p
            $key = '0000id/type/file.jpg';
            // s
            $key = 'test/logo.jpg';
            $key = '00000id/type2/img.jpg';
        }
        $command = $this->s3->getCommand('GetObject', [
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => $key,
        ]);
        $r = $command->createPresignedUrl('+5 minutes');
        return $r;
    }

    private function uploadTxt($config_aws)
    {
        $r = $this->s3->putObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => 'test/data.txt.' . time(),
            'Body'   => 'Hello! ' . time(),
        ]);
        return $r;
    }

    private function uploadImg($config_aws, $key = 'bu/test/BOND.jpg', $sourceFile = '/home/kovpak/Downloads/b.jpg')
    {
        $r = $this->s3->putObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => $key,
            'SourceFile' => $sourceFile,
        ]);
        return $r;
    }

    private function uploadVideo($config_aws, $key = 'test/test.video.mp4', $sourceFile = '/home/kovpak/Downloads/test.video.mp4')
    {
        $r = $this->s3->putObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key'    => $key,
            'SourceFile' => $sourceFile,
        ]);
        return $r;
    }

    private function optimize($config_aws)
    {
        $url = $this->getUrl($config_aws, 'test/BOND.origin.jpg');
        $this->download($url);
        `convert /tmp/tmpS3ImageFile.jpg -resize 50% /tmp/tmpS3ImageFile.result.jpg`;
        $this->uploadImg($config_aws, 'test/BOND.result.jpg', '/tmp/tmpS3ImageFile.result.jpg');
    }

    private function download($uri, $targetFile = '/tmp/tmpS3ImageFile.jpg')
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $uri);
        curl_setopt($ch, CURLOPT_HEADER, false);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_BINARYTRANSFER, true);
        $response = curl_exec($ch);
        if ($response === false) {
            throw new RuntimeException(curl_error($ch));
        }
        if (file_exists($targetFile)) {
            unlink($targetFile);
        }
        $fp = fopen($targetFile, 'x');
        fwrite($fp, $response);
        fclose($fp);
    }

    private function copy($config_aws, $source, $target)
    {
        $r = $this->s3->copyObject([
            'Bucket' => $config_aws->s3->bucket,
            'CopySource' => urlencode($config_aws->s3->bucket . '/' . $source),
            'Key' => $target,
        ]);
        return $r;
    }

    private function createCopy($config_aws, $key = 'test/BOND.jpg')
    {
        return $this->copy($config_aws, $key, $key.'.copy');
    }

    private function fileExists($config_aws, $key = 'test/BOND.jpg')
    {
        try {
            $r = $this->s3->getObject([
                'Bucket' => $config_aws->s3->bucket,
                'Key' => $key,
            ]);
            return $r;
        } catch (NoSuchKeyException $e) {
            return false;
        }
    }

    private function deleteFile($config_aws, $key = '')
    {
        $r = $this->s3->deleteObject([
            'Bucket' => $config_aws->s3->bucket,
            'Key' => $key,
        ]);
        return $r;
    }

    private function getObjectsList($config_aws, $prefix = '')
    {
        $r = $this->s3->getIterator(
            'ListObjects',
            ['Bucket' => $config_aws->s3->bucket, 'Prefix' => $prefix]
        );
        return $r;
    }

    public function deleteCopies($config_aws)
    {
        foreach ($this->getObjectsList($config_aws) as $el) {
            if (preg_match('/.*\.copy$/', $el['Key'])) {
                var_export($this->deleteFile($config_aws, $el['Key']));
            }
        }
    }

    public function findAbandonedThumbnails($config_aws)
    {
        foreach ($this->getObjectsList($config_aws, '') as $el) {
            if (preg_match('/_thumbnail.jpg$/', $el['Key'])) {
                $key = str_replace('_thumbnail.jpg', '.jpg', $el['Key']);
                if (!$this->fileExists($config_aws, $key)) {
                    $this->copy($config_aws, $el['Key'], "trash/$key");
                    $this->deleteFile($config_aws, $el['Key']);
                    var_dump($el['Key']);
                }
            }
        }
    }
}

new Command($config_aws, $argv[1]);
