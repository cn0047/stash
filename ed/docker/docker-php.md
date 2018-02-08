docker-php
-

````
# php
docker build -t cn007b/php:7.1 ./docker/7.1
docker push cn007b/php:7.1
# test
docker run -it --rm cn007b/php:7.1 php -v

# php-composer
docker build -t cn007b/php:7.1-composer ./docker/7.1-composer
docker push cn007b/php:7.1-composer
# test
docker run -it --rm cn007b/php:7.1-composer php -v
docker run -it --rm cn007b/php:7.1-composer composer

# php-protobuf
docker build -t cn007b/php:7.1-protobuf ./docker/7.1-protobuf
docker push cn007b/php:7.1-protobuf
# test
docker run -it --rm cn007b/php:7.1-protobuf php -v
docker run -it --rm cn007b/php:7.1-protobuf composer
docker run -it --rm cn007b/php:7.1-protobuf protoc --version

# php-protobuf-3
docker build -t cn007b/php:7.1-protobuf-3 ./docker/7.1-protobuf-3
docker push cn007b/php:7.1-protobuf-3
# test
docker run -it --rm cn007b/php:7.1-protobuf-3 php -v
docker run -it --rm cn007b/php:7.1-protobuf-3 composer
docker run -it --rm cn007b/php:7.1-protobuf-3 protoc --version

# php-fpm
docker build -t cn007b/php:7.1-fpm ./docker/7.1-fpm
docker push cn007b/php:7.1-fpm
# test
docker run -it --rm cn007b/php:7.1-fpm php -v
docker run -it --rm cn007b/php:7.1-fpm composer
docker run -it --rm cn007b/php:7.1-fpm protoc --version
docker run -it --rm cn007b/php:7.1-fpm service --status-all | grep fpm

# php-nginx
docker build -t cn007b/php:7.1-nginx ./docker/7.1-nginx
docker push cn007b/php:7.1-nginx
# test
docker run -it --rm cn007b/php:7.1-nginx php -v
docker run -it --rm cn007b/php:7.1-nginx composer
docker run -it --rm cn007b/php:7.1-nginx protoc --version
docker run -it --rm cn007b/php:7.1-nginx service --status-all | grep fpm
docker run -it --rm cn007b/php:7.1-nginx service --status-all | grep nginx

# php latest
docker build -t cn007b/php:latest ./docker/7.1-nginx
docker push cn007b/php:latest
# test latest
docker run -it --rm cn007b/php:latest php -v
docker run -it --rm cn007b/php:latest composer
docker run -it --rm cn007b/php:latest protoc --version
docker run -it --rm cn007b/php:latest service --status-all | grep fpm
docker run -it --rm cn007b/php:latest service --status-all | grep nginx
# test
docker run -it --rm cn007b/php php -v
docker run -it --rm cn007b/php composer
docker run -it --rm cn007b/php protoc --version
docker run -it --rm cn007b/php service --status-all | grep fpm
docker run -it --rm cn007b/php service --status-all | grep nginx

# php-phalcon
# docker build -t cn007b/php:7.1-phalcon ./docker/7.1-phalcon
````
