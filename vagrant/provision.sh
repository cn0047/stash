#!/usr/bin/env bash

# ubuntu
apt-get update

# php
add-apt-repository ppa:ondrej/php
apt-get update
apt-get install -y php7.0 php7.0-fpm php7.0-cli php7.0-opcache php7.0-common php7.0-phpdbg php7.0-dev
apt-get install -y php7.0-mcrypt php7.0-mbstring
apt-get install -y php7.0-mysql php7.0-pdo
apt-get install -y php7.0-dom php7.0-xml php7.0-json
apt-get install -y php7.0-zip php7.0-curl php7.0-gd php7.0-imap
sudo apt-get install php-mongodb

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
apt-get install -y nginx
cp /vagrant/vagrant/nginx.conf /etc/nginx/sites-available/default
service nginx restart

# mysql
apt-get install -y mysql-server

# mongodb
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list
sudo apt-get update
sudo apt-get install -y mongodb-org
sudo cp /vagrant/vagrant/mongodb.service /etc/systemd/system/
sudo systemctl start mongodb
sudo systemctl enable mongodb



#
cd /vagrant/ed/laravel/examples/one \
&& composer install \
&& php artisan cache:clear \
&& php artisan config:cache \
&& php artisan migrate \
# chmod 777 -R /vagrant/ed/laravel/examples/one/storage/
# chmod 777 -R /vagrant/ed/laravel/examples/one/bootstrap/cache/
#
# mysql -uroot -e 'create database laravelOne'
# mysql -uroot -e "CREATE USER 'laravelUser'@'localhost' IDENTIFIED BY 'laravelPassword'"
# mysql -uroot -e "GRANT ALL PRIVILEGES ON laravelOne.* TO 'laravelUser'@'localhost' WITH GRANT OPTION;"
# mysql -ularavelUser -plaravelPassword -DlaravelOne
#
mysql -uroot -e 'create database homestead'
mysql -uroot -e "CREATE USER 'homestead'@'localhost' IDENTIFIED BY 'secret'"
mysql -uroot -e "GRANT ALL PRIVILEGES ON homestead.* TO 'homestead'@'localhost' WITH GRANT OPTION;"
mysql -uhomestead -psecret -Dhomestead
cd /vagrant/ed/laravel/examples/one && php artisan migrate

