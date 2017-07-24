Benchmarking ElasticSearch vendors for PHP
-

## Introduction

This is a simple example which provides different options
how to receive data from ElasticSearch into PHP.

In this test, we have simple `type` inside ElasticSearch `index` which contains 1K records,
look [here](https://github.com/cn007b/my/blob/master/ed/php/examples/elasticsearch/comparison/import.sh).

I've used:
PHP from official docker container - `php:7.1-cli`;
Composer from - `composer:php7`;
ElasticSearch from - `elasticsearch:5.3.0`.

Also, I've used all default setting for PHP and ElasticSearch.

## Prepare infrastructure

For this benchmarkin you need `docker` and `curl`.
Run in terminal next commands:

````
mkdir /tmp/php-es/
````

````
# Download code:
curl -o /tmp/php-es/index.php "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/index.php"
curl -o /tmp/php-es/composer.json "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/composer.json"
curl -o /tmp/php-es/import.sh "https://raw.githubusercontent.com/cn007b/my/master/ed/php/examples/elasticsearch/comparison/import.sh"
````

````
# Install composer packages:
docker run -ti --rm -v /tmp/php-es/:/app composer/composer:php7 install
````

````
# Run ElasticSearch:
docker run -d -p 9200:9200 --hostname localhost --name es elasticsearch:5.3.0
# IMPORTANT: please wait few secconds till request `curl localhost:9200` return valid response.

# Import data into ElasticSearch:
sh /tmp/php-es/import.sh
# As result you must receive message like this - Imported:  {"count":1000,"_shards":{"total":5,"successful":5,"failed":0}}
````

## Run benchmark

````
docker run -it --rm -v /tmp/php-es/:/app --link es php:7.1-cli php /app/index.php
````

ENJOY!!!

## Result

For me on ubuntu 16.04 with PHP 7.1.3 and ES 5.3.0 result was:

````
Elastica      took: 0.115600
Elasticsearch took: 0.243299
Curl          took: 0.006087
Bash          took: 0.019164
````

## CONTRIBUTING

Ussing `docker` you can easily re-test my benchmark and prove or disproof it.

In case I made mistake - please let me know through [twitter](https://twitter.com/cn007b) or [email](cn007b@gmail.com).
In case you wanna add additional information or additional vendor - please make pull-request.
