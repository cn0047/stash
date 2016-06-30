<?php

require_once __DIR__.'/vendor/autoload.php';
require_once __DIR__.'/../../config.php';

use Aws\Ec2\Ec2Client;

$ec2 = Ec2Client::factory([
    'region' => $config_aws->region,
    'credentials' => (array) $config_aws->credentials,
    'version' => 'latest',
]);
var_dump($ec2->DescribeInstances());
