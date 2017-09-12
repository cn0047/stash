docker
-
1.12.5

## Overview

Containers should be ephemeral (can be stopped and destroyed and a new one built and put in place).
<br>Hence container must be stateless.
<br>Each container should have only one concern.

````
docker pull ubuntu

# build image from Dokerfile in dir ./examples/mydockerbuild
# `-t` it is tag
# `.` it is current directory 
docker build -t docker-whale .

# run an interactive container 
# `-t` terminal
# `-i` interactive connection
docker run -t -i ubuntu:latest /bin/bash

# run a web application
# `-d` runs the container as daemon
# `-p` maps any required network ports
docker run -d -p 8081:80 timber/ziipr
docker run -d -p 192.168.0.32:1111:1111 timber/ziipr

# exec
docker exec -it happy_babbage
docker exec -it happy_babbage bash
cd /opt/docker/etc/supervisor.d

docker-machine ip

# information about all the containers
docker ps -a

# shows the standard output of a container
docker logs

# show containers & images
docker images

# stop the running container
docker stop

docker start nostalgic_morse

# remove the web application container
docker rm nostalgic_morse

# delete
docker rmi -f docker-whale

exit
````

### RUN

````
# bash
docker build -t xubuntu ./docker/ubuntu
docker run -ti --rm -v $PWD/u.sh:/u.sh xubuntu /u.sh

# composer
docker build -t xcomposer ./docker/composer
docker run -ti --rm -v $PWD:/app xcomposer install
````

#### NODEJS

````
docker run -it --rm node:latest node -v
docker run -it --rm -v $PWD:/gh -w /gh node:latest node /gh/x.js

# simple mysql test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql --link mysql-master node:latest node index.js

# simple elasticsearch test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch --link es node:latest node index.js

# simple mongo test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --link xmongo node:latest node index.js
````

#### PHP

````
# php-cli
docker build -t php-cli ./docker/php-cli
docker run -it --rm -v $PWD:/gh php-cli php -v
docker run -it --rm -v $PWD:/gh php-cli php /gh/x.php

# php-fpm
docker build -t php-fpm ./docker/php-fpm
docker run -it --rm -p 9000:9000 --hostname localhost --name php-fpm -v $PWD:/gh php-fpm

# buil-in web server
docker run -it --rm -p 8000:8000 -v $PWD:/gh php-cli php -S 0.0.0.0:8000 /gh/ed/php/examples/isNumeric.php

# mysql
docker run -it --rm -v $PWD:/gh --link mysql-master php-cli php /gh/ed/php/examples/mysqlAndPdo/pdo.simplestExample.php

# RabbitMQ with php
docker run -ti --rm -v $PWD/ed/php/examples/rabbitmq/tutorials:/app xcomposer install
# direct
docker run -it --rm -v $PWD/ed:/gh/ed --link rabbit \
    php-cli php /gh/ed/php/examples/rabbitmq/tutorials/routingLikeStream/emit_log_direct.php
docker run -it --rm -v $PWD/ed:/gh/ed --link rabbit \
    php-cli php /gh/ed/php/examples/rabbitmq/tutorials/routingLikeStream/receive_logs_direct.php
# worker
docker run -it --rm -v $PWD/ed:/gh/ed --link rabbit \
    php-cli php /gh/ed/php/examples/rabbitmq/tutorials/workQueue/worker.php
docker run -it --rm -v $PWD/ed:/gh/ed --link rabbit \
    php-cli php /gh/ed/php/examples/rabbitmq/tutorials/workQueue/new_task.php

# laravel
docker run -ti --rm -v $PWD/ed/php.laravel/examples/one:/app xcomposer install
docker run -it --rm -v $PWD:/gh php-cli php /gh/ed/php.laravel/examples/one/artisan key:generate
docker run -it --rm -v $PWD:/gh --link mysql-master php-cli php /gh/ed/php.laravel/examples/one/artisan migrate
docker run -it --rm --hostname 0.0.0.0 -p 8181:8181 -v $PWD:/gh --link mysql-master \
    php-cli php /gh/ed/php.laravel/examples/one/artisan serve --host=0.0.0.0 --port=8181

# symfony
docker run -ti --rm -v $PWD/ed/php.symfony/examples/News/symfony/news:/app xcomposer install
docker run -it --rm --hostname 0.0.0.0 -p 8181:8181 -v $PWD:/gh --link mysql-master \
    php-cli php /gh/ed/php.symfony/examples/News/symfony/news/bin/console server:run 0.0.0.0:8181

# symfony + nginx
docker run -ti --rm --name nginx-php --link php-fpm \
    -v $PWD/docker/nginx/symfony.news.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh -p 8080:80 nginx:latest
# prepare prod environment
docker run -it --rm -v $PWD:/gh \
    php-cli php /gh/ed/php.symfony/examples/News/symfony/news/bin/symfony_requirements
docker run -ti --rm -v $PWD/ed/php.symfony/examples/News/symfony/news:/app \
    xcomposer install --no-dev --optimize-autoloader
docker run -it --rm -v $PWD:/gh \
    php-cli php /gh/ed/php.symfony/examples/News/symfony/news/bin/console \
    cache:clear --env=prod --no-debug --no-warmup
docker run -it --rm -v $PWD:/gh \
    php-cli php /gh/ed/php.symfony/examples/News/symfony/news/bin/console \
    cache:warmup --env=prod
docker run -it --rm -v $PWD:/gh \
    php-cli php /gh/ed/php.symfony/examples/News/symfony/news/bin/console \
    assets:dump --env=prod --no-debug
# test
curl localhost:8080/
````

#### NGINX

````
# html
docker run -ti --rm --name nginx-html \
    -v $PWD/docker/nginx/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest

# test
curl http://localhost:8081/bootstrap.popover.html

# php
docker run -ti --rm --name nginx-php --link php-fpm \
    -v $PWD/docker/nginx/php-fpm.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest

# test
curl localhost:8080/healthCheck.php

````

#### MONGO

````
docker run -it --rm --hostname localhost --name xmongo -v $PWD/docker/mongodb/db:/data/db -p 27017:27017 mongo:latest

docker exec -it xmongo mongo test --eval 'db.test.insert({code : 200, status: "ok"})'
docker exec -it xmongo mongo test \
    --eval 'db.createUser({user: "dbu", pwd: "dbp", roles: ["readWrite", "dbAdmin"]})'
````

#### ES cluster

````
docker run -it --rm -p 9200:9200 --name es elasticsearch:latest

# init master 1 node
docker run -it --rm -p 9200:9200 --name es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=true -Des.node.data=false

# init data 1 node
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=false -Des.node.data=true \
    -Des.discovery.zen.ping.unicast.hosts=es-master-1
````

````
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
````

#### MYSQL cluster

````
# init master node
docker run -it --rm -p 3307:3306 --name mysql-master --hostname mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-master.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp mysql:latest

# replication user on master
docker exec mysql-master mysql -uroot -proot -e "CREATE USER 'repl'@'%' IDENTIFIED BY 'slavepass'"
docker exec mysql-master mysql -uroot -proot -e "GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%'"

# init slave 1 node
docker run -it --rm -p 3308:3306 --name mysql-slave-1 --link mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-slave-1.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e MYSQL_USER=dbu2 -e MYSQL_PASSWORD=dbp2 mysql:latest

# start slave 1
docker exec mysql-slave-1 mysql -uroot -proot -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='repl', MASTER_PASSWORD='slavepass'"
docker exec mysql-slave-1 mysql -uroot -proot -e "START SLAVE"
docker exec mysql-slave-1 mysql -uroot -proot -e "SHOW SLAVE STATUS \G"

# test
docker exec -ti mysql-master mysql -P3307 -udbu -pdbp -Dtest
````

#### RabbitMQ

````
# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# check rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged
````

#### Non-finished swarm.

````
# init rabbit

#
docker run -it --rm --name php-cli-rabbitmq-c -v $PWD/ed:/gh/ed --link rabbit php-cli

#
docker-machine create --driver virtualbox manager1
docker-machine ip manager1
# 192.168.99.100
docker-machine ssh manager1
docker swarm init --advertise-addr 192.168.99.100
docker node ls
#
docker service create --replicas 3 --name php-cli-rabbitmq-swarm php-cli-rabbitmq-c \
    php /gh/ed/php/examples/rabbitmq/tutorials/workQueue/worker.php
docker service ls
docker service ps php-cli-rabbitmq-swarm

#
docker-machine create --driver virtualbox worker1
docker-machine ssh worker1
docker swarm join \
    --token SWMTKN-1-3ie1vxyfmvh1756tv37dyp8datyyfcsfrnkhmzofwk3nsle7ud-cblwlb51iv251evjiudwxs6li \
    192.168.99.100:2377
docker node ls
````

## Machine

Use [machine](https://docs.docker.com/machine) to create Docker hosts on your local box,
on your company network, in your data center,
or on cloud providers like AWS or Digital Ocean.

````
docker-machine version

docker-machine ls

docker-machine create --driver virtualbox manager1
````

## Compose

[Compose](https://docs.docker.com/compose) is a tool
for defining and running multi-container Docker applications.

````
# build the project and detached
docker-compose up -d

# shutdown/clean up
docker-compose down
docker-compose down --volumes

# builds, (re)creates, starts, and attaches to containers for a service.
docker-compose up

# runs a one-time command against a service
docker-compose run

docker-compose exec php-cli php /gh/x.php
docker-compose exec mysql /bin/bash

# restart stopped container
docker-compose start
````

## Swarm

Docker [Swarm](https://docs.docker.com/swarm) is native clustering for Docker.
It turns a pool of Docker hosts into a single, virtual Docker host.

## Network

````
# list these networks
docker network ls

docker network inspect
````

https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/

https://docs.docker.com/compose/compose-file/
