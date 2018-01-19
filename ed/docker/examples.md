Examples
-

### RUN

````
docker network create --driver bridge xnet

# bash
docker build -t xubuntu ./docker/ubuntu
docker run -ti --rm xubuntu /bin/bash

# composer
docker build -t xcomposer ./docker/composer
docker run -ti --rm -v $PWD:/app xcomposer install
````

#### MONGO

````
docker run -it --rm -p 27017:27017 --hostname localhost --name xmongo --net=x_node_mongo \
    -v /Users/k/Downloads/:/tmp/d \
    -v $PWD/docker/.data/mongodb:/data/db mongo:latest

# dump
docker exec -it xmongo mongorestore /tmp/d/creating_documents/dump

# test
docker exec -it xmongo mongo
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

#### MYSQL

````
docker run -it --rm --net=xnet -p 3307:3306 --name xmysql --hostname xmysql \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test mysql:latest

docker exec -ti xmysql mysql -P3307 -udbu -pdbp -Dtest

# root
docker run -it --rm --net=xnet -p 3307:3306 --name xmysql --hostname xmysql \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root mysql:latest
````

#### MYSQL cluster

````
# init master node
docker run -it --rm -p 3307:3306 --name mysql-master --hostname mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-master.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test mysql:latest

# replication user on master
docker exec mysql-master mysql -uroot -proot -e "CREATE USER 'repl'@'%' IDENTIFIED BY 'slavepass'"
docker exec mysql-master mysql -uroot -proot -e "GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%'"

# init slave 1 node
docker run -it --rm -p 3308:3306 --name mysql-slave-1 --link mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-slave-1.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu2 -e MYSQL_PASSWORD=dbp2 -e MYSQL_DATABASE=test mysql:latest

# start slave 1
docker exec mysql-slave-1 mysql -uroot -proot -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='repl', MASTER_PASSWORD='slavepass'"
docker exec mysql-slave-1 mysql -uroot -proot -e "START SLAVE"
docker exec mysql-slave-1 mysql -uroot -proot -e "SHOW SLAVE STATUS \G"

# test
docker exec -ti mysql-master mysql -P3307 -udbu -pdbp -Dtest
````

#### POSTGRESQL

````
docker run -it --rm --name xpostgres --hostname xpostgres --net=xnet \
    -v $PWD/docker/.data/postgresql/xpostgres:/var/lib/postgresql/data \
    -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# import dump
# docker exec -ti cpqsql /bin/bash -c "psql -d postgres://dbu:dbp@cpqsql/test < /app/dump.sql"

# test
docker exec -ti xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti -e PGPASSWORD=dbp xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@localhost/test
````

#### POSTGRESQL cluster

````
docker run -it --rm -p 5432:5432 --name postgres-master --hostname postgres-master \
    -v $PWD/docker/.data/postgresql:/var/lib/postgresql/data \
    -v $PWD/docker/postgresql/master.conf:/var/lib/postgresql/data/postgresql.conf \
    -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# test
docker exec -ti postgres-master psql -h localhost -p 5432 -U dbu -d test
````

#### REDIS

````
# init redis
docker run -it --rm -p 6379:6379 --hostname xredis --name xredis redis:latest

# check redis
docker exec -ti redis redis-cli
````

#### RabbitMQ

````
# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# check rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged
````

#### NODEJS

````
docker run -it --rm node:latest node -v

# from Dockerfile
docker build -t xnodejs ./docker/nodejs
docker run -it --rm -p 8080:3000 xnodejs
# to test
curl 0.0.0.0:8080

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
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --link xmongo node:latest node index.js
# simple mongo test with bridge
docker network create --driver bridge x_node_mongo
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --net=x_node_mongo node:latest node index.js
````

#### NGINX

````
# html
docker run -ti --rm --name nginx-html \
    -v $PWD/docker/nginx/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest

# test
curl http://localhost:8080/bootstrap.popover.html

# https html
docker run -ti --rm --name nginx-html \
    -v $PWD/docker/nginx/https/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD/docker/nginx/https/localhost.crt:/ssl/localhost.crt \
    -v $PWD/docker/nginx/https/localhost.key:/ssl/localhost.key \
    -v $PWD:/gh \
    -p 3443:443 nginx:latest

# test
curl https://localhost:3443/bootstrap.popover.html

# php (all scripts from `ed/php/examples`)
docker run -ti --rm --name nginx-and-php --link php-fpm \
    -v $PWD/docker/nginx/php-fpm.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD:/gh \
    -p 8080:80 nginx:latest
# # php with xnet
# docker run -ti --rm --name nginx-and-php --net=xnet \
#     -v $PWD/docker/nginx/php-fpm.conf:/etc/nginx/conf.d/default.conf \
#     -v $PWD:/gh \
#     -p 8080:80 nginx:latest

# test
curl localhost:8080/healthCheck.php
````

#### PHP

````
# php-nginx
docker build -t nphp ./docker/php-nginx
docker run -it --rm -p 8080:80 -v $PWD:/gh nphp php -v

# composer
docker run -it --rm -v $PWD:/app -w /app nphp composer --help

# built-in web server
# docker run -it --rm -p 8080:80 --link mysql-master -v $PWD:/gh nphp \
#     php -S 0.0.0.0:80 /gh/ed/php/examples/whatever/healthCheck.php
docker run -it --rm --net=xnet -p 8080:80 -v $PWD:/gh nphp \
    php -S 0.0.0.0:80 /gh/ed/php/examples/whatever/healthCheck.php
# nginx
docker run -it --rm -p 8080:80 -v $PWD:/gh nphp \
    /bin/bash -c 'service php7.1-fpm start; service nginx start; tail -f /dev/stdout'
# test
curl http://localhost:8080/healthCheck.php?XDEBUG_SESSION_START=PHPSTORM

# php-cli
docker build -t php-cli ./docker/php-cli
docker run -it --rm -v $PWD:/gh php-cli php -v
docker run -it --rm -v $PWD:/gh php-cli php /gh/x.php

# php-fpm
docker build -t php-fpm ./docker/php-fpm
docker run -it --rm -p 9000:9000 --hostname localhost --name php-fpm -v $PWD:/gh php-fpm
docker run -it --rm -p 9000:9000 --hostname localhost --name php-fpm -v $PWD:/gh --net=xnet php-fpm

# mysql
# docker run -it --rm --link mysql-master -v $PWD:/gh php-cli php /gh/ed/php/examples/mysqlAndPdo/pdo.simplestExample.php
docker run -it --rm --net=xnet -v $PWD:/gh php-cli php /gh/ed/php/examples/mysqlAndPdo/pdo.simplestExample.php

# postgres
docker run -it --rm -v $PWD:/gh --net=xnet \
    php-cli php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.simplestExample.php

# mongo
docker run -it --rm -v $PWD:/gh --link xmongo php-cli php /gh/ed/php/examples/whatever/mongo.simplest.php

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

# phalcon
docker build -t php-cli-phalcon ./docker/php-cli-phalcon
docker run -it --rm -v $PWD:/gh php-cli-phalcon php -v

# phpspec
docker run -ti --rm -v $PWD/ed/php.phpspec/examples/one:/app -w /app \
    nphp php vendor/phpspec/phpspec/bin/phpspec desc App/Boo
# run
docker run -ti --rm -v $PWD/ed/php.phpspec/examples/one:/app -w /app \
    nphp php vendor/phpspec/phpspec/bin/phpspec run

# codeception
docker run -ti --rm -v $PWD/ed/php.codeception/examples/one:/app -w /app \
    nphp php vendor/bin/codecept bootstrap
# run
docker run -ti --rm -v $PWD/ed/php.codeception/examples/one:/app -w /app \
    nphp php vendor/bin/codecept run

# kahlan
docker run -ti --rm -v $PWD/ed:/gh/ed -w /gh/ed/php.kahlan/examples/one \
    -e PHP_IDE_CONFIG="serverName=docker" \
    nphp php vendor/bin/kahlan
````

#### PHP Symfony

````
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

# ng
docker run -ti --rm -v $PWD/ed/php.symfony/examples/ng:/app -w /app nphp composer install
docker run -ti --rm -v $PWD/ed/php.symfony/examples/ng:/app -w /app \
    nphp php vendor/phpspec/phpspec/bin/phpspec run
````

#### PHP Yii

````
mkdir ed/php.yii/examples/testdrive/protected/runtime
mkdir ed/php.yii/examples/testdrive/assets

docker run -it --rm -v $PWD/ed/php.yii/examples/testdrive:/app -w /app nphp composer install

# main db
docker exec -ti xmysql mysql -P3307 -uroot -proot -e 'create database testdrive'
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "create user 'user'@'%' identified by 'pass'"
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "grant all privileges on testdrive.* to 'user'@'%' with grant option"
# test db
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "
    create database testdrive_unit_test;
    grant all privileges on testdrive_unit_test.* to 'user'@'%' with grant option;
"
# check
docker exec -ti xmysql mysql -P3307 -uuser -ppass -Dtestdrive
docker exec -ti xmysql mysql -P3307 -uuser -ppass -Dtestdrive_unit_test

# migration
docker run -it --rm --net=xnet -v $PWD/ed/php.yii/examples/testdrive:/app -w /app \
    nphp php protected/yiic.php migrate

# phpunit
docker run -it --rm --net=xnet \
    -v $PWD/ed/php.yii/examples/testdrive:/app -w /app/protected/tests \
    nphp php ../../vendor/bin/phpunit ./

docker run -it --rm --net=xnet -p 8080:8080 -v $PWD/ed/php.yii/examples/testdrive:/app nphp \
    php -S 0.0.0.0:8080 -t /app
````
