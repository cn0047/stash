<?php
/**
 * PHP benchmarking ElasticSearch vendors.
 *
 * README:
 *
 * This is a simple example which provides different options
 * how to receive data from ElasticSearch into PHP.
 * In this test, we have simple `type` inside ElasticSearch `index` which contains 1K records,
 * @see https://github.com/cn007b/my/blob/master/ed/php/examples/elasticsearch/comparison
 * @see https://github.com/cn007b/my/blob/master/ed/php/examples/elasticsearch/comparison/import.sh
 * I've used PHP from official docker container - php:7.1-cli;
 * Composer from - composer:php7;
 * ElasticSearch from - elasticsearch:5.3.0.
 * Also, I've used all default setting for PHP and ElasticSearch.
 * Hence you can easily re-test my benchmark and prove or disproof it.
 *
 * CONTRIBUTING:
 *
 * In case I made mistake - please let me know through [twitter](https://twitter.com/cn0047) or [email](cn007b@gmail.com).
 * In case you wanna add additional information or additional vendor - please make pull-request.
 */

/*
INIT:
  
For this benchmarkin you need docker and curl.

Run in terminal next commands:

# Download code:
mkdir /tmp/php-es/
curl -o /tmp/php-es/index.php "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/index.php"
curl -o /tmp/php-es/composer.json "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/composer.json"
curl -o /tmp/php-es/import.sh "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/import.sh"
   
# Install composer packages:
docker run -ti --rm -v /tmp/php-es/:/app composer/composer:php7 install

# Run ElasticSearch:
docker run -d -p 9200:9200 --hostname localhost --name es elasticsearch:5.3.0
# IMPORTANT: please wait few secconds till request `curl localhost:9200` return valid response.

# Import data into ElasticSearch:
sh /tmp/php-es/import.sh
# As result you must receive message like this - Imported:  {"count":1000,"_shards":{"total":5,"successful":5,"failed":0}}

RUN:

# Run benchmark:
docker run -it --rm -v /tmp/php-es/:/app --link es php:7.1-cli php /app/index.php

ENJOY!!!
*/

declare(strict_types = 1);

/**
 * * @see https://github.com/cn007b/my/blob/master/ed/php/examples/elasticsearch/comparison/composer.json
 */
require __DIR__ . '/vendor/autoload.php';

// It is super simple ElasticSearch query.
$queryParameters = [
    'query' => [
        'bool' => [
            'filter' => [
                'bool' => [
                    'must' => [
                        ['range' => ['age' => ['gt' => 30, 'lt' => 40]]],
                    ],
                ],
            ],
        ],
    ],
    'size' => 1000,
];
$jsonQueryParameters = json_encode($queryParameters);

// It is super simple lambda function
// which used for verifying result obtained from ElasticSearch.
$getUsers = function (array $iterator) {
    $result = [];
    foreach ($iterator as $user) {
        $data = $user['_source'];
        $result[] = $data['first_name'] . ' ' . $data['last_name'];
    }
    return $result;
};

/**
 * Elastica client.
 */
$t1 = microtime(true);
$client = new Elastica\Client();
$search = new Elastica\Search($client);
$search->addIndex('somecorp');
$search->addType('employee');
$query = new Elastica\Query();
$query->setRawQuery($queryParameters);
$search->setQuery($query);
$resultSet = $search->search();
$users1 = [];
foreach ($resultSet as $user) {
    $data = $user->getSource();
    $users1[] = $data['first_name'] . ' ' . $data['last_name'];
}
$d1 = microtime(true) - $t1;

/**
 * Official Elasticsearch client.
 */
$t2 = microtime(true);
$client = Elasticsearch\ClientBuilder::create()->build();
$params = ['index' => 'somecorp', 'type' => 'employee', 'body' => $queryParameters];
$resultSet = $client->search($params);
$users2 = $getUsers($resultSet['hits']['hits']);
$d2 = microtime(true) - $t2;

/**
 * Simple curl.
 */
$t3 = microtime(true);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, 'localhost:9200/somecorp/employee/_search');
curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'GET');
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
curl_setopt($ch, CURLOPT_POSTFIELDS, $jsonQueryParameters);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);
$resultSet = json_decode($response, true);
$users3 = $getUsers($resultSet['hits']['hits']);
$d3 = microtime(true) - $t3;

/**
 * Bash.
 */
$t4 = microtime(true);
$response = `curl -s -XGET localhost:9200/somecorp/employee/_search -d '$jsonQueryParameters'`;
$resultSet = json_decode($response, true);
$users4 = $getUsers($resultSet['hits']['hits']);
$d4 = microtime(true) - $t4;

printf('Elastica      took: %f, result hash: %s %s', $d1, md5(serialize($users1)), PHP_EOL);
printf('Elasticsearch took: %f, result hash: %s %s', $d2, md5(serialize($users2)), PHP_EOL);
printf('Curl          took: %f, result hash: %s %s', $d3, md5(serialize($users3)), PHP_EOL);
printf('Bash          took: %f, result hash: %s %s', $d4, md5(serialize($users4)), PHP_EOL);

/**
 * Result.
 *
 * For me on ubuntu 16.04 with PHP 7.1.3 and ES 5.3.0 result was:
 *
 * Elastica      took: 0.115600, result hash: 0b472524a817277d06b3fac3f136de76
 * Elasticsearch took: 0.243299, result hash: 0b472524a817277d06b3fac3f136de76
 * Curl          took: 0.006087, result hash: 0b472524a817277d06b3fac3f136de76
 * Bash          took: 0.019164, result hash: 0b472524a817277d06b3fac3f136de76
 *
 * Elastica      took: 0.117375, result hash: 0b472524a817277d06b3fac3f136de76
 * Elasticsearch took: 0.174706, result hash: 0b472524a817277d06b3fac3f136de76
 * Curl          took: 0.005072, result hash: 0b472524a817277d06b3fac3f136de76
 * Bash          took: 0.016369, result hash: 0b472524a817277d06b3fac3f136de76
 *
 * Elastica      took: 0.122851, result hash: 0b472524a817277d06b3fac3f136de76
 * Elasticsearch took: 0.192609, result hash: 0b472524a817277d06b3fac3f136de76
 * Curl          took: 0.005957, result hash: 0b472524a817277d06b3fac3f136de76
 * Bash          took: 0.019113, result hash: 0b472524a817277d06b3fac3f136de76
 *
 * Elastica      took: 0.108815, result hash: 0b472524a817277d06b3fac3f136de76
 * Elasticsearch took: 0.187460, result hash: 0b472524a817277d06b3fac3f136de76
 * Curl          took: 0.005355, result hash: 0b472524a817277d06b3fac3f136de76
 * Bash          took: 0.021437, result hash: 0b472524a817277d06b3fac3f136de76
 *
 * Elastica      took: 0.104871, result hash: 0b472524a817277d06b3fac3f136de76
 * Elasticsearch took: 0.146614, result hash: 0b472524a817277d06b3fac3f136de76
 * Curl          took: 0.005068, result hash: 0b472524a817277d06b3fac3f136de76
 * Bash          took: 0.019623, result hash: 0b472524a817277d06b3fac3f136de76
 */
