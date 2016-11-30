#!/usr/bin/env bash

# ubuntu
sudo apt-get update

# composer
sudo curl -sS https://getcomposer.org/installer | sudo php
sudo mv composer.phar /usr/local/bin/composer

# apache
sudo service apache2 stop

# nginx
sudo apt-get install -y nginx

# git
sudo apt-get install -y git

# mysql
# sudo apt-get install -y mysql-server
# mysql -uroot -e 'create database test'
# mysql -uroot -e "SET global general_log_file='/var/log/mysql/general.log';"
# mysql -uroot -e "SET global general_log = 1;"

# sqlite
sudo apt-get install sqlite3 libsqlite3-dev

# elasticsearch
sudo apt-get install -y openjdk-7-jre
curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.3.3.deb
sudo dpkg -i elasticsearch-2.3.3.deb
sudo /etc/init.d/elasticsearch start

# import elasticsearch megacorp test index
sh /vagrant/ed/elasticsearch/megacorp.sh
