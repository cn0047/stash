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
