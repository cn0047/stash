# docker-php

# php
docker build -t cn007b/php:7.1 ./docker/7.1
# test
docker run -it --rm cn007b/php:7.1 php -v
docker run -it --rm cn007b/php:7.1-composer composer
# push
docker push cn007b/php:7.1

# php-protobuf-2
docker build -t cn007b/php:7.1-protobuf-2 ./docker/7.1-protobuf-2
# test
docker run -it --rm cn007b/php:7.1-protobuf protoc --version
# push
docker push cn007b/php:7.1-protobuf-2

# php-protobuf-3
docker build -t cn007b/php:7.1-protobuf-3 ./docker/7.1-protobuf-3
# test
docker run -it --rm cn007b/php:7.1-protobuf-3 protoc --version
# push
docker push cn007b/php:7.1-protobuf-3

# php-nginx
docker build -t cn007b/php:7.1-nginx ./docker/7.1-nginx
# test
docker run -it --rm cn007b/php:7.1-nginx service --status-all | grep fpm
docker run -it --rm cn007b/php:7.1-nginx service --status-all | grep nginx
# push
docker push cn007b/php:7.1-nginx

# php latest
docker build -t cn007b/php:latest ./docker/7.1-nginx
# push
docker push cn007b/php:latest

# php-phalcon
# docker build -t cn007b/php:7.1-phalcon ./docker/7.1-phalcon
