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

var_dump(new \Elastica\Document());
die;

// // Bulk indexing.
// for($i = 0; $i < 10; $i++) {
//     // $params['body'][] = [
//     //     'index' => [
//     //         '_id' => 'my_id_' . $i,
//     //         'type' => 'my_type',
//     //     ],
//     //     'my_index',
//     //     'type' => 'my_type',
//     // ];
//     // $params['body'][] = [
//     //     'my_field' => 'my_value',
//     //     'second_field' => 'some more values',
//     // ];

//     // $params['body'][] = [
//     //     'index' => 'my_index',
//     //     'type' => 'my_type',
//     //     'id' => 'my_id_' . $i,
//     //     'body' => ['testOk' => 'true'],
//     //     'timestamp' => time(),
//     // ];
//     // $params['body'][] = [
//     //     'my_field' => 'my_value',
//     //     'second_field' => 'some more values'
//     // ];
// }
// $responses = $client->bulk($params);
// var_export($response);

// Get document.
$params = [
    'index' => 'my_index',
    'type' => 'my_type',
    'id' => $id,
];
$response = $client->get($params);

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
//
$params = [
    'index' => 'my_index',
    'type' => 'my_type',
    // 'timestamp' => [
    //     'query' => [
    //         'wildcard' => '*',
    //     ]
    // ]
    'body' => [
        'query' => [
            'bool' => [
                'should' => [
                    'wildcard' => ['timestamp' => '*'],
                ],
            ],
        ],
    ],
];
$response = $client->search($params);


//
var_export($response);




// curl "localhost:9200/my_index/_search?search_type=count" -d '{
//     "aggs": {
//         "count_by_type": {
//             "terms": {
//                 "field": "_type"
//             }
//         }
//     }
// }'