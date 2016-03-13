<?php

require __DIR__ . '/vendor/autoload.php';

$client = new Elastica\Client();

$search = new Elastica\Search($client);
$search->addIndex('megacorp');
$search->addType('employee');

$query = new Elastica\Query();
$query->setSort([
    '_geo_distance' => [
        'location' => ['lat' => '51.5072', 'lon' => '0.1275'],
        'order' => 'asc',
        'unit' => 'km',
        'distance_type' => 'plane',
    ]
]);
$search->setQuery($query);
$resultSet = $search->search();
print_r($resultSet);
