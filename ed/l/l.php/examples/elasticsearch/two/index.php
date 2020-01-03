<?php

require 'vendor/autoload.php';

$client = Elasticsearch\ClientBuilder::create()->build();

// Create the index.
$params = ['index' => 'my_index'];
try {
    $resonse = $client->indices()->create($params);
} catch (Elasticsearch\Common\Exceptions\BadRequest400Exception $e) {
    // index already exists
}

// Indexing document.
$id = 'my_id_' . uniqid();
$params = [
    'index' => 'my_index',
    'type' => 'my_type',
    'id' => $id,
    'body' => ['testOk' => (bool)mt_rand(0, 1)],
    'timestamp' => time(),
];
$response = $client->index($params);
// print_r($response);

// Bulk indexing.
$params = [];
for($i = 0; $i < 5; $i++) {
    $params['body'][] = [
        'index' => [
            '_index' => 'my_index',
            '_type' => 'my_type',
        ]
    ];
    $params['body'][] = [
        'index' => 'my_index',
        'type' => 'my_type',
        'body' => ['testOk' => (bool)mt_rand(0, 1)],
        'timestamp' => time(),
    ];
}
$responses = $client->bulk($params);
// print_r($response);

// Get document by id.
$params = [
    'index' => 'my_index',
    'type' => 'my_type',
    'id' => $id,
];
$response = $client->get($params);
// print_r($response);

// Search.
$params = [
    'index' => 'my_index',
    'type' => 'my_type',
    'body' => [
        'query' => [
            'match' => [
                'testOk' => (bool)mt_rand(0, 1)
            ]
        ]
    ]
];
$response = $client->search($params);
// var_export($response);
