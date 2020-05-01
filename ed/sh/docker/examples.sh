Examples
-

sudo ifconfig lo0 alias 10.254.254.254

docker network create --driver bridge xnet

### Linux

# Ubuntu
docker pull cn007b/ubuntu
docker tag cn007b/ubuntu xubuntu
docker run -ti --rm --net=xnet -v $PWD:/gh -w /gh xubuntu /bin/bash

# Debian
docker tag cn007b/debian xdebian
docker run -ti --rm -v $PWD:/gh xdebian bash

# Alpine
docker run -ti --rm -v $PWD:/gh alpine:3.7 sh

#### Memcached

# run
docker run -it --rm --net=xnet -p 11211:11211 --hostname xmemcached --name xmemcached kmemcached
# check
docker exec -it xmemcached telnet 0.0.0.0 11211

# build
docker build -t kmemcached -f ed/sh/docker/examples.Dockerfile/cache.memcached.Dockerfile .

#### RabbitMQ

# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# check rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged

#### NGINX

# html
docker run -ti --rm --name nginx-html \
    -v $PWD/.docker/nginx/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest

# test
curl http://localhost:8080/bootstrap.popover.html

# https html
docker run -ti --rm --name nginx-html \
    -v $PWD/.docker/nginx/https/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD/.docker/nginx/https/localhost.crt:/ssl/localhost.crt \
    -v $PWD/.docker/nginx/https/localhost.key:/ssl/localhost.key \
    -v $PWD:/gh \
    -p 3443:443 nginx:latest

# test
curl https://localhost:3443/bootstrap.popover.html

# php (all scripts from `ed/php/examples`)
docker run -ti --rm --name nginx-and-php --link php-fpm \
    -v $PWD/.docker/nginx/php-fpm.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest
# # php with xnet
# docker run -ti --rm --name nginx-and-php --net=xnet \
#     -v $PWD/.docker/nginx/php-fpm.conf:/etc/nginx/conf.d/default.conf \
#     -v $PWD:/gh \
#     -p 8080:80 nginx:latest

# test
curl localhost:8080/healthCheck.php
