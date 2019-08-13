Examples
-

sudo ifconfig lo0 alias 10.254.254.254

docker network create --driver bridge xnet

### Linux

# Ubuntu
docker pull cn007b/ubuntu
docker tag cn007b/ubuntu xubuntu
docker run -ti --rm -v $PWD:/gh -w /gh xubuntu /bin/bash

# Debian
docker tag cn007b/debian xdebian
docker run -ti --rm -v $PWD:/gh xdebian bash

# Alpine
docker run -ti --rm -v $PWD:/gh alpine:3.7 sh

#### Memcached

docker run -it --rm --net=xnet --name xmemcached memcached

# build
docker build -t kmemcached ./.docker/memcached

# run
docker run -it --rm --net=xnet -p 11211:11211 --hostname xmemcached --name xmemcached kmemcached

# check
docker exec -it xmemcached telnet 0.0.0.0 11211

#### ES cluster

docker run -it --rm -p 9200:9200 --name es elasticsearch:latest

# init master 1 node
docker run -it --rm -p 9200:9200 --name es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=true -Des.node.data=false

# init data 1 node
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=false -Des.node.data=true \
    -Des.discovery.zen.ping.unicast.hosts=es-master-1

# init master 1 node
docker run -it --rm -p 9200:9200 --name es-master-1 \
    -e "bootstrap.memory_lock=true" -e "ES_JAVA_OPTS=-Xms256m -Xmx256m" \
    -e "http.host=_eth0_" -e "cluster.name=ec" \
    -e "node.master=true" -e "node.data=false" elasticsearch:5.4

# init data 1 node
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1  \
    -e "bootstrap.memory_lock=true" -e "ES_JAVA_OPTS=-Xms256m -Xmx256m" \
    -e "http.host=_eth0_" -e "cluster.name=ec" \
    -e "node.master=false" -e "node.data=true" -e "discovery.zen.ping.unicast.hosts=es-master-1" elasticsearch:5.4

#### REDIS

# init redis
docker run -it --rm -p 6379:6379 --hostname xredis --name xredis redis:latest

# check redis
docker exec -ti xredis redis-cli

#### RabbitMQ

# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# check rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged

#### NODEJS

docker run -it --rm node:latest node -v

# from Dockerfile
docker build -t xnodejs ./.docker/nodejs
docker run -it --rm -p 8000:3000 xnodejs
# to test
curl 0.0.0.0:8000

# based on Alpine Linux
docker run -it --rm node:alpine node -v
docker run -it --rm -v $PWD:/gh -w /gh node:latest node /gh/x.js

# simple mysql test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql --link mysql-master node:latest node index.js

# simple elasticsearch test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch --link es node:latest node index.js

# simple mongo test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --net=xnet node:latest node index.js
# simple mongo test with bridge
docker network create --driver bridge x_node_mongo
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node index.js
#
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node mongo.universityhw3-3.js
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 node npm i
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 \
    node:latest node overviewOrTags.js

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