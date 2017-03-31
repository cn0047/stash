#!/usr/bin/env bash

# Ubuntu 16.04

###############################################################################
# GENERAL
###############################################################################

# ubuntu
sudo apt-get update

# htop
sudo apt-get install -y htop

# git
sudo apt-get install -y git

# apache
sudo service apache2 stop

# nginx
sudo apt-get install -y nginx
sudo cp /vagrant/vagrant/nginx.conf /etc/nginx/sites-available/default
sudo service nginx restart

# mysql
sudo debconf-set-selections <<< 'mysql-server-5.7 mysql-server/root_password password root'
sudo debconf-set-selections <<< 'mysql-server-5.7 mysql-server/root_password_again password root'
sudo apt-get -y install mysql-server-5.7
sudo apt-get install -y mysql-server
# sudo apt-get install -y mysql-client
sudo mysql -uroot -proot -e 'create database test'
sudo mysql -uroot -proot -e "create user 'user'@'%' identified by 'pass'"
sudo mysql -uroot -proot -e "grant all privileges on test.* to 'user'@'%' with grant option"
sudo mysql -uroot -proot -e "set global general_log_file='/var/log/mysql/general.log';"
sudo mysql -uroot -proot -e "set global general_log = 1;"
# IMPORTANT
sudo sed -i "s/127.0.0.1/0.0.0.0/" `sudo grep bind-address -ril /etc/mysql/`
sudo service mysql restart

# sqlite
# sudo apt-get install sqlite3 libsqlite3-dev

# mongodb
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
sudo echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list
sudo apt-get update
sudo apt-get install -y mongodb
sudo apt-get install -y mongodb-org
sudo cp /vagrant/vagrant/mongodb.service /etc/systemd/system/
sudo systemctl start mongodb
sudo systemctl enable mongodb

# redis
sudo apt-get install -y redis-server

# java
sudo add-apt-repository -y ppa:webupd8team/java
sudo apt-get update
echo debconf shared/accepted-oracle-license-v1-1 select true | sudo debconf-set-selections
echo debconf shared/accepted-oracle-license-v1-1 seen true | sudo debconf-set-selections
sudo apt-get -y install oracle-java8-installer

# elasticsearch
# 2.4.4
sudo curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.4.4.deb
sudo dpkg -i elasticsearch-2.4.4.deb
sudo service elasticsearch start
# 5.3.0
sudo wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
sudo apt-get install apt-transport-https
sudo echo "deb https://artifacts.elastic.co/packages/5.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-5.x.list
sudo apt-get update && sudo apt-get install elasticsearch
sudo service elasticsearch start

# rabbit
sudo apt-get install -y rabbitmq-server

# beanstalk
# sudo apt-get install -y beanstalkd

# aws cli
sudo apt-get -y install python-pip
sudo pip install awscli --ignore-installed six
# init aws commands
sudo chmod +x /vagrant/ed/bash/examples/describeAwsInstances.sh
sudo ln -s /vagrant/ed/bash/examples/describeAwsInstances.sh /usr/bin/dai
sudo chmod +x /vagrant/ed/bash/examples/executeCommandOnAws.sh
sudo ln -s /vagrant/ed/bash/examples/executeCommandOnAws.sh /usr/bin/ecoa
sudo chmod +x /vagrant/ed/bash/examples/sshToAws.sh
sudo ln -s /vagrant/ed/bash/examples/sshToAws.sh /usr/bin/sta

# nodejs
sudo apt-get install python-software-properties
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo apt-get install -y npm
# npm packages
# sudo npm install -g create-react-app react react-dom
# sudo npm install webpack -g webpack webpack-dev-server

# php 5.6
# sudo add-apt-repository ppa:ondrej/php
# sudo apt-get update
# sudo apt-get install -y php5.6 php5.6-cli php5.6-fpm php5.6-mongo php5.6-common php5.6-dev
# sudo apt-get install -y php5.6-mcrypt php5.6-mbstring
# sudo apt-get install -y php5.6-pdo php5.6-mysql php5.6-mysqli php5.6-sqlite
# sudo apt-get install -y php5.6-dom php5.6-xml php5.6-json
# sudo apt-get install -y php5.6-zip  php5.6-curl php5.6-gd php5.6-imap

# php 7.0
sudo add-apt-repository -y ppa:ondrej/php
sudo apt-get update
sudo apt-get install -y php7.0 php7.0-fpm php7.0-cli php7.0-opcache php7.0-common php7.0-phpdbg php7.0-dev
sudo apt-get install -y php7.0-mcrypt php7.0-mbstring php7.0-bcmath
sudo apt-get install -y php7.0-mysql php7.0-pdo php-mongodb php7.0-sqlite php7.0-sqlite3 php7.0-redis
sudo apt-get install -y php7.0-dom php7.0-xml php7.0-json
sudo apt-get install -y php7.0-zip php7.0-curl php7.0-gd php7.0-imap
sudo apt-get install -y php7.0-imagick

# php 7.1
sudo add-apt-repository -y ppa:ondrej/php
sudo apt-get update
sudo apt-get install -y php7.1 php7.1-fpm php7.1-cli php7.1-opcache php7.1-common php7.1-phpdbg php7.1-dev
sudo apt-get install -y php7.1-mcrypt php7.1-mbstring php7.1-bcmath
sudo apt-get install -y php7.1-mysql php7.1-pdo php-mongodb php7.1-sqlite php7.1-sqlite3 php7.1-redis
sudo apt-get install -y php7.1-dom php7.1-xml php7.1-json
sudo apt-get install -y php7.1-zip php7.1-curl php7.1-gd php7.1-imap
sudo apt-get install -y php7.1-imagick

# bz2 calendar cgi cli ctype dba enchant exif fileinfo fpm ftp gettext gmp iconv
# interbase intl ldap mcrypt mysql mysqli mysqlnd odbc pdo-dblib pdo-firebird
# pdo-mysql pdo-odbc pdo-pgsql pgsql phar posix pspell readline recode
# shmop simplexml snmp soap sockets sqlite3 sybase sysvmsg sysvsem sysvshm tidy
# tokenizer wddx xmlreader xmlrpc xmlwriter xsl

# imagemagick
sudo apt-get install -y imagemagick

# xdebug
sudo apt-get install -y php-xdebug

# composer
sudo curl -sS https://getcomposer.org/installer | sudo php
sudo mv composer.phar /usr/local/bin/composer

# heroku
sudo apt-get install -y software-properties-common
sudo add-apt-repository "deb https://cli-assets.heroku.com/branches/stable/apt ./"
curl -L https://cli-assets.heroku.com/apt/release.key | sudo apt-key add -
sudo apt-get update
sudo apt-get install -y heroku

###############################################################################
# LARAVEL
###############################################################################

# one
# cd /vagrant/ed/laravel/examples/one \
#     && composer install \
#     && php artisan cache:clear \
#     && php artisan config:cache \
#     && php artisan migrate \
# # chmod 777 -R /vagrant/ed/laravel/examples/one/storage/
# # chmod 777 -R /vagrant/ed/laravel/examples/one/bootstrap/cache/
# mysql -uroot -e 'create database homestead'
# mysql -uroot -e "CREATE USER 'homestead'@'localhost' IDENTIFIED BY 'secret'"
# mysql -uroot -e "GRANT ALL PRIVILEGES ON homestead.* TO 'homestead'@'localhost' WITH GRANT OPTION;"
# mysql -uhomestead -psecret -Dhomestead
# cd /vagrant/ed/laravel/examples/one && php artisan migrate

###############################################################################
# Phalcon
###############################################################################

# phalcon
# sudo curl -s https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh | sudo bash
# sudo apt-get install php7.0-phalcon
# # phalcon dev-tools
# cd /vagrant/ed/phalcon/examples/one && composer install
# sudo ln -s /vagrant/ed/phalcon/examples/one/vendor/bin/phalcon.php /usr/bin/phalcon
# sudo chmod ugo+x /usr/bin/phalcon

###############################################################################
# Symfony
###############################################################################

#bulletinBoard
# cd /vagrant/ed/symfony/examples/bulletinBoard/ && sudo php app/console cache:clear --env=prod --no-debug
# cd /vagrant/ed/symfony/examples/bulletinBoard/ && sudo php app/console assetic:dump --env=prod --no-debug
# composer update -o

# v3
cd /vagrant/ed/symfony/examples/v3.2/ \
    && rm -rf var/cache/* \
    && rm -rf var/logs/* \
    && sudo php bin/console cache:clear --env=prod --no-debug \
    && sudo php bin/console cache:warmup --env=prod --no-debug \

###############################################################################
# ElasticSearch
###############################################################################

# import elasticsearch megacorp test index
sh /vagrant/ed/elasticsearch/megacorp.sh
