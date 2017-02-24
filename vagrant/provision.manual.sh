#!/usr/bin/env bash

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

# logstash
# echo 'deb http://packages.elastic.co/logstash/2.2/debian stable main' | sudo tee /etc/apt/sources.list.d/logstash-2.2.x.list
# sudo apt-get update
# sudo apt-get install logstash

# redis
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

# aws
# aws configure
