<?php

require 'vendor/autoload.php';

$client = Elasticsearch\ClientBuilder::create()->build();

// Search Jackie Chan
$params = [
    'index' => 'megacorp',
    'type' => 'employee',
    'body' => [
        'query' => [
            'match' => [
                'first_name' => 'Jackie'
            ]
        ]
    ]
];
$response = $client->search($params);
// var_export($response);

// Sorting by distance from London.
$params = [
    'index' => 'megacorp',
    'type' => 'employee',
    'fields' => ['city'],
    'body' => [
        'sort' => [
            '_geo_distance' => [
                'location' => ['lat' => '51.5072', 'lon' => '0.1275'],
                'order' => 'asc',
                'unit' => 'km',
                'distance_type' => 'plane',
            ]
        ]
    ],
];
$response = $client->search($params);
// var_export($response);

// Search by range.
$params = [
    'index' => 'megacorp',
    'type' => 'employee',
    'body' => [
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
        ]
    ]
];
$response = $client->search($params);
print_r($response);
