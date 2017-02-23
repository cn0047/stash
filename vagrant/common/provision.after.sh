#!/usr/bin/env bash

# nginx
sudo service nginx restart

# composer
sudo curl -sS https://getcomposer.org/installer | sudo php
sudo mv composer.phar /usr/local/bin/composer

##############################################################################
# By hands:
##############################################################################

# mysql
# sudo apt-get install -y mysql-server
# mysql -uroot -e 'create database test'
# mysql -uroot -e "create user 'user'@'localhost' identified by 'pass'"
# mysql -uroot -e "grant all privileges on test.* to 'user'@'localhost' with grant option"
# mysql -uroot -e "SET global general_log_file='/var/log/mysql/general.log';"
# mysql -uroot -e "SET global general_log = 1;"


# # java
# sudo add-apt-repository -y ppa:webupd8team/java
# sudo apt-get update
# sudo apt-get -y install oracle-java8-installer

# # elasticsearch
# curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.4.4.deb
# sudo dpkg -i elasticsearch-2.4.4.deb
# sudo /etc/init.d/elasticsearch start

# # import elasticsearch megacorp test index
# sh /vagrant/ed/elasticsearch/megacorp.sh

# aws
# aws configure
