<?php

require_once __DIR__.'/vendor/autoload.php';
require_once __DIR__.'/../../config.php';

use Aws\Ec2\Ec2Client;

$config = [
    'region' => $config_aws->region,
    'credentials' => (array)$config_aws->credentials,
    'version' => 'latest',
];
$ec2 = Ec2Client::factory($config);
$args = [
    'Filters' => [
        ['Name' => 'tag:Name', 'Values' => ['prod-zii-web']],
    ]
];
$data = $ec2->describeInstances($args)->toArray();
$instances = [];
array_walk_recursive($data, function ($value, $key) use (&$instances) {
    if ($key === 'PublicDnsName') {
        $instances[$value] = true;
    }
});
var_export($instances);
