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
# # sudo apt-get install -y mysql-server
# # mysql -uroot -e 'create database test'
# # mysql -uroot -e "SET global general_log_file='/var/log/mysql/general.log';"
# # mysql -uroot -e "SET global general_log = 1;"

# # sqlite
# sudo apt-get install sqlite3 libsqlite3-dev

# # java
# sudo add-apt-repository -y ppa:webupd8team/java
# sudo apt-get update
# sudo apt-get -y install oracle-java8-installer

# # elasticsearch
# sudo apt-get install -y openjdk-7-jre
# curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.3.3.deb
# sudo dpkg -i elasticsearch-2.3.3.deb
# sudo /etc/init.d/elasticsearch start

# # import elasticsearch megacorp test index
# sh /vagrant/ed/elasticsearch/megacorp.sh

# # mongodb
# sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
# sudo echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list
# sudo apt-get update
# sudo apt-get install -y mongodb
# sudo apt-get install -y mongodb-org
# sudo cp /vagrant/vagrant/common/mongodb.service /etc/systemd/system/
# sudo systemctl start mongodb
# sudo systemctl enable mongodb

# # logstash
# # echo 'deb http://packages.elastic.co/logstash/2.2/debian stable main' | sudo tee /etc/apt/sources.list.d/logstash-2.2.x.list
# # sudo apt-get update
# # sudo apt-get install logstash

# redis (MANUALLY)
# sudo apt-get install -y build-essential
# sudo apt-get install -y tcl8.5
# wget http://download.redis.io/releases/redis-stable.tar.gz
# tar xzf redis-stable.tar.gz
# cd redis-stable
# make
# make test
# sudo make install
# cd utils
# sudo ./install_server.sh
# sudo service redis_6379 start
