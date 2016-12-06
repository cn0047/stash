#!/usr/bin/env bash

###############################################################################
# GENERAL
###############################################################################

# php 5.5
sudo apt-get install software-properties-common
sudo add-apt-repository ppa:ondrej/php5-5.6
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install php5 php5.6-fpm

# php 7
# add-apt-repository ppa:ondrej/php
# apt-get update
# apt-get install -y php7.0 php7.0-fpm php7.0-cli php7.0-opcache php7.0-common php7.0-phpdbg php7.0-dev
# apt-get install -y php7.0-mcrypt php7.0-mbstring
# apt-get install -y php7.0-mysql php7.0-pdo
# apt-get install -y php7.0-dom php7.0-xml php7.0-json
# apt-get install -y php7.0-zip php7.0-curl php7.0-gd php7.0-imap
# apt-get install -y php-mongodb
# php7.0-bcmath
# php7.0-bz2
# php7.0-calendar
# php7.0-cgi
# php7.0-cli
# php7.0-ctype
# php7.0-dba
# php7.0-enchant
# php7.0-exif
# php7.0-fileinfo
# php7.0-fpm
# php7.0-ftp
# php7.0-gettext
# php7.0-gmp
# php7.0-iconv
# php7.0-interbase
# php7.0-intl
# php7.0-ldap
# php7.0-mcrypt
# php7.0-mysql
# php7.0-mysqli
# php7.0-mysqlnd
# php7.0-odbc
# php7.0-pdo-dblib
# php7.0-pdo-firebird
# php7.0-pdo-mysql
# php7.0-pdo-odbc
# php7.0-pdo-pgsql
# php7.0-pdo-sqlite
# php7.0-pgsql
# php7.0-phar
# php7.0-posix
# php7.0-pspell
# php7.0-readline
# php7.0-recode
# php7.0-shmop
# php7.0-simplexml
# php7.0-snmp
# php7.0-soap
# php7.0-sockets
# php7.0-sqlite3
# php7.0-sybase
# php7.0-sysvmsg
# php7.0-sysvsem
# php7.0-sysvshm
# php7.0-tidy
# php7.0-tokenizer
# php7.0-wddx
# php7.0-xmlreader
# php7.0-xmlrpc
# php7.0-xmlwriter
# php7.0-xsl

# composer
curl -sS https://getcomposer.org/installer | php
mv composer.phar /usr/local/bin/composer

# apache
# service apache2 stop

# nginx
cp /vagrant/vagrant/nginx.conf /etc/nginx/sites-available/default

# java
sudo add-apt-repository -y ppa:webupd8team/java
sudo apt-get update
sudo apt-get -y install oracle-java8-installer

# logstash
echo 'deb http://packages.elastic.co/logstash/2.2/debian stable main' | sudo tee /etc/apt/sources.list.d/logstash-2.2.x.list
sudo apt-get update
sudo apt-get install logstash



###############################################################################
# LARAVEL
###############################################################################

# one
cd /vagrant/ed/laravel/examples/one \
    && composer install \
    && php artisan cache:clear \
    && php artisan config:cache \
    && php artisan migrate \
# chmod 777 -R /vagrant/ed/laravel/examples/one/storage/
# chmod 777 -R /vagrant/ed/laravel/examples/one/bootstrap/cache/
mysql -uroot -e 'create database homestead'
mysql -uroot -e "CREATE USER 'homestead'@'localhost' IDENTIFIED BY 'secret'"
mysql -uroot -e "GRANT ALL PRIVILEGES ON homestead.* TO 'homestead'@'localhost' WITH GRANT OPTION;"
mysql -uhomestead -psecret -Dhomestead
cd /vagrant/ed/laravel/examples/one && php artisan migrate

###############################################################################
# Phalcon
###############################################################################

# phalcon
curl -s https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh | bash
apt-get install php5-phalcon
# apt-get install php7.0-phalcon

# dirtondirt
mysql -uroot -e 'create database dirtondirt'
mysql -uroot -e "CREATE USER 'dod'@'localhost' IDENTIFIED BY '111'"
mysql -uroot -e "GRANT ALL PRIVILEGES ON dirtondirt.* TO 'dod'@'localhost' WITH GRANT OPTION;"
