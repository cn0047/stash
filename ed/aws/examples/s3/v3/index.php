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
     * @example php index.php fileExists '000010215/big.mp4'
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

    /**
     * Downloads picture by public url and save it in temporary file.
     *
     * @param string $url Public url to picture.
     * @param string $file Target file name.
     *
     * @throws \RuntimeException In case when picture download failed.
     *
     * @example php index.php download 'https://s3-eu-west-1.amazonaws.com/w3.stage.ziipr.bucket/000010215/big.mp4' '/tmp/ziipr.v.mp4'
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
     * @example php index.php download2 'https://s3-eu-west-1.amazonaws.com/bucket/000010215/big.mp4' '/tmp/ziipr.v.mp4'
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

