<?php

require __DIR__ . '/vendor/autoload.php';

$client = new Elastica\Client();

$search = new Elastica\Search($client);
$search->addIndex('megacorp');
$search->addType('employee');

$index = $client->getIndex('index');
$type = $index->getType('type');
$document = $tpype->getDocument(1);

// Sort by geo distance.
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
// print_r($resultSet);

// Search Jackie Chan
$query = new Elastica\Query();
$params = [
    'query' => [
        'match' => [
            'first_name' => 'Jackie'
        ]
    ]
];
$query->setRawQuery($params);
$search->setQuery($query);
$resultSet = $search->search();
// print_r($resultSet);

// query + sort
$query = new Elastica\Query();
$params = [
    'query' => [
        'filtered' => [
            'filter' => [
                'bool' => [
                    'must' => [
                        ['range' => ['age' => ['gt' => 30, 'lt' => 40]]]
                    ]
                ]
            ]
        ]
    ],
    'sort' => [
        '_geo_distance' => [
            'location' => ['lat' => '51.5072', 'lon' => '0.1275'],
            'order' => 'asc',
            'unit' => 'km',
            'distance_type' => 'plane',
        ]
    ]
];
$query->setRawQuery($params);
$search->setQuery($query);
$resultSet = $search->search();
// print_r($resultSet);

// Simple search
$query = new Elastica\Query();
$query->addSort(['age' => 'desc']);
$search->setQuery($query);
$resultSet = $search->search();
//print_r($resultSet);
