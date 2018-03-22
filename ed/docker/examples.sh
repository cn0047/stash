Examples
-

````
sudo ifconfig lo0 alias 10.254.254.254

docker network create --driver bridge xnet
````

### Ubuntu

````
docker tag cn007b/ubuntu:17.10 xubuntu
docker run -ti --rm xubuntu /bin/bash
````

#### Memcached

````
docker run -it --rm --net=xnet --name xmemcached memcached

# build
docker build -t kmemcached ./docker/memcached

# run
docker run -it --rm --net=xnet -p 11211:11211 --hostname xmemcached --name xmemcached kmemcached

# check
docker exec -it xmemcached telnet 0.0.0.0 11211
````

#### MONGO

````
docker run -it --rm --net=xnet -p 27017:27017 --hostname xmongo --name xmongo \
    -v /Users/k/Downloads/:/tmp/d \
    -v $PWD/docker/.data/mongodb:/data/db mongo:latest

# dump
docker exec -it xmongo mongorestore /tmp/d/creating_documents/dump
docker exec -it xmongo mongoimport --drop -d crunchbase -c companies /tmp/d/mongo/findAndCursorsInNodeJSDriver/companies.json

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
    -v $PWD/docker/.data/mysql:/var/lib/mysql -v /tmp:/tmp \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test mysql:latest

# general_log
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "set global general_log_file='/tmp/mysql.log';"
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "set global general_log = 1;"
docker exec -ti xmysql tail -f /tmp/mysql.log

docker exec -ti xmysql mysql -P3307 -uroot -proot
docker exec -ti xmysql mysql -P3307 -udbu -pdbp -Dtest
````

#### POSTGRESQL

````
docker run -it --rm --name xpostgres --hostname xpostgres --net=xnet \
    -v $PWD/docker/.data/postgresql/xpostgres:/var/lib/postgresql/data \
    -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# check
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test -c 'select count(*) from test'

# dump
docker exec -ti xpostgres pg_dump -d postgres://dbu:dbp@xpostgres/test -t test --schema-only

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
docker exec -ti xredis redis-cli
````

#### RabbitMQ

````
# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# check rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged
````

#### GO (GOLANG)

````
docker build -t xgo ./docker/go
docker run -it --rm -v $PWD:/gh -w /gh xgo go

docker run -it --rm -v $PWD:/gh -w /gh golang:latest go
docker run -it --rm -v $PWD:/gh -w /gh golang:latest go run /gh/ed/go/examples/hw.go

docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh' golang:latest sh -c 'echo $GOPATH'

# db postgresql
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/lib/pq'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go run src/postgresql/simplest.go'

# db mongo
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get gopkg.in/mgo.v2'
# or
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get ./...'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go run src/mongodb/simple.go'

````

````
# Simple Web Server

# web.one
docker run -it --rm -p 8000:8000 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && go run src/firstapp/main.go'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && cd src/firstapp && go test -cover'
# check
curl http://localhost:8000/
curl http://localhost:8000/health-check

# Livereload
#
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/codegangsta/gin'
docker run -it --rm --name go-one -p 8000:8000 -p 8001:8001 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8001 --appPort 8000 --path src/firstapp/ run main.go'

# 2
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && go install templateapp'
docker run -it --rm -p 8000:8000 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && ./bin/templateapp'
# http://localhost:8000/home

# web.two
docker run -it --rm -p 8000:8000 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.two/' \
    golang:latest sh -c 'cd $GOPATH && go run src/webapp/main.go'
# http://localhost:8000/home

# web.three.tiny
docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three.tiny/' \
    golang:latest sh -c 'cd $GOPATH && go run src/app/main.go'

# web.three ⭐️ ⭐️ ⭐️
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' golang:latest sh -c '
        cd $GOPATH \
        && go get gopkg.in/mgo.v2 \
        && go get github.com/codegangsta/gin \
        && go get -u github.com/derekparker/delve/cmd/dlv
    '
# run
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
    golang:latest sh -c 'cd $GOPATH && go run src/app/main.go'
# livereload
docker run -it --rm --net=xnet -p 8080:8080 -p 8081:8081 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
#     golang:latest sh -c 'cd $GOPATH && ./bin/dlv debug src/app/main.go'

# test
curl -i 'http://localhost:8080'
curl -i 'http://localhost:8080/home'
curl -i 'http://localhost:8080/cars'
curl -i -XGET 'http://localhost:8080/cars'
curl -i -XPUT 'http://localhost:8080/cars'
curl -i -XDELETE 'http://localhost:8080/cars'
curl -i -XPOST 'http://localhost:8080/cars'
curl -i -XPOST 'http://localhost:8080/cars' -H 'Content-Type: application/json' \
   -d '{"vendor": "BMW", "name": "X5"}'
# test lr
curl -i -XGET 'http://localhost:8081/cars'
curl -i -XPUT 'http://localhost:8081/cars'
curl -i -XDELETE 'http://localhost:8081/cars/1'
curl -i -XPOST 'http://localhost:8081/cars' -H 'Content-Type: application/json' \
   -d '{"vendor": "BMW", "name": "M6"}'

# # web.HTTPS
# docker run -it --rm -p 8000:8000 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.https/' \
#     golang:latest sh -c 'cd $GOPATH && go install webapp'
# docker run -it --rm -p 8000:8000 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.https/' \
#     golang:latest sh -c 'cd $GOPATH && ./bin/webapp'

# # zeromq
# docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/zeromq/' \
#     xgo sh -c 'cd $GOPATH && go get github.com/pebbe/zmq4'
# docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/zeromq/' \
#     golang:latest sh -c 'cd $GOPATH && go run src/hw/server.go'
````

````
# api-gateway

docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
        go get github.com/codegangsta/gin;
        go get -u golang.org/x/lint/golint;
    '

docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
        ./bin/golint src/app/...
    '

docker run -it --rm -p 8080:8080 -p 8081:8081 \
    -v $PWD:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c './bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'

curl -i 'http://localhost:8081/github/users/cn007b'
````

#### GO Echo

````
# one

# init
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' golang:latest sh -c '
    cd $GOPATH \
    && go get -u github.com/labstack/echo/... \
    && go get -u github.com/codegangsta/gin
'
# run
docker run -it --rm -p 8080:8080 -p 8081:8081 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' golang:latest sh -c '
    cd $GOPATH && cd src/app && go test -cover
'
# check
curl -i -XGET 'http://localhost:8081'
curl -i -XGET 'http://localhost:8081/products'
curl -i -XGET 'http://localhost:8081/products/iphone'
````

#### GO Gin

````
docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/gin-gonic/gin'

docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
    golang:latest sh -c 'cd $GOPATH && go run src/one/main.go'

# curl localhost:8080/v1/file-info/id/7
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
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --net=xnet node:latest node index.js
# simple mongo test with bridge
docker network create --driver bridge x_node_mongo
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node index.js
#
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node mongo.universityhw3-3.js
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 node npm i
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 \
    node:latest node overviewOrTags.js
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
docker pull cn007b/php:7.1-protobuf-3
docker tag cn007b/php:7.1-protobuf-3 xphp
````

````
# php-nginx local
docker pull cn007b/php
docker tag cn007b/php nphp
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
    nphp php /gh/ed/php/examples/mysqlAndPdo/pdo.postgresql.simplestExample.php

# mongo
docker run -it --rm -v $PWD:/gh --link xmongo php-cli php /gh/ed/php/examples/whatever/mongo.simplest.php

# memcached
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.1.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.get.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.add.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.set.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.increment.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.replace.php
docker run -it --rm --net=xnet -v $PWD:/gh nphp php /gh/ed/php/examples/whatever/memcache.delete.php
docker run -it --rm --net=xnet -v $PWD:/gh -e PHP_IDE_CONFIG='serverName=docker' \
    nphp php /gh/ed/php/examples/whatever/memcache.TaskProgress.php

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
# docker build -t php-cli-phalcon ./docker/php-cli-phalcon
# docker run -it --rm -v $PWD:/gh php-cli-phalcon php -v

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
    -e PHP_IDE_CONFIG='serverName=docker' \
    nphp php vendor/bin/kahlan
````

#### PHP Symfony

````
# symfony
docker run -ti --rm -v $PWD/ed/php.symfony/examples/News/symfony/news:/app -w /app xphp composer install
docker run -it --rm --net=xnet --hostname 0.0.0.0 -p 8181:8181 -v $PWD:/gh \
    xphp php /gh/ed/php.symfony/examples/News/symfony/news/bin/console server:run 0.0.0.0:8181

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
docker run -it --rm --net=xnet -e PHP_IDE_CONFIG='serverName=docker' \
    -v $PWD/ed/php.yii/examples/testdrive:/app -w /app/protected/tests \
    nphp php ../../vendor/bin/phpunit ./

docker run -it --rm --net=xnet -p 8080:8080 -v $PWD/ed/php.yii/examples/testdrive:/app nphp \
    php -S 0.0.0.0:8080 -t /app
````
