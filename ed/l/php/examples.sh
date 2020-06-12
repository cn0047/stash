# php

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
docker build -t php-cli ./.docker/php-cli
docker run -it --rm -v $PWD:/gh php-cli php -v
docker run -it --rm -v $PWD:/gh php-cli php /gh/x.php

# php-fpm
docker build -t php-fpm ./.docker/php-fpm
docker run -it --rm -p 9000:9000 --hostname localhost --name php-fpm -v $PWD:/gh php-fpm
docker run -it --rm -p 9000:9000 --hostname localhost --name php-fpm -v $PWD:/gh --net=xnet php-fpm

# mysql
# docker run -it --rm --link mysql-master -v $PWD:/gh php-cli php /gh/ed/php/examples/mysqlAndPdo/pdo.simplestExample.php
docker run -it --rm --net=xnet -v $PWD:/gh xphp php /gh/ed/php/examples/mysqlAndPdo/pdo.simplestExample.php

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
# docker build -t php-cli-phalcon ./.docker/php-cli-phalcon
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

#### PHP Symfony

# symfony
docker run -ti --rm -v $PWD/ed/php.symfony/examples/News/symfony/news:/app -w /app xphp composer install
docker run -it --rm --net=xnet --hostname 0.0.0.0 -p 8181:8181 -v $PWD:/gh \
    xphp php /gh/ed/php.symfony/examples/News/symfony/news/bin/console server:run 0.0.0.0:8181

# symfony + nginx
docker run -ti --rm --name nginx-php --link php-fpm \
    -v $PWD/.docker/nginx/symfony.news.conf:/etc/nginx/conf.d/default.conf \
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

# bulletinBoard - won't work
# docker run -ti --rm -v $PWD/ed/php.symfony/examples/bulletinBoard:/app -w /app xphp composer install

#### PHP Yii

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

#### AppEngine

cd /Users/k/web/kovpak/gh/ed/php.appengine/examples/one
composer install

~/.google-cloud-sdk/bin/dev_appserver.py \
    --port=8080 --admin_port=8000 --skip_sdk_update_check=true app.yaml

gcloud config set project thisissimplebot
gcloud app deploy
