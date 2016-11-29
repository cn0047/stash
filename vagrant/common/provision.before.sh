#!/usr/bin/env bash

# ubuntu
sudo apt-get update

# composer
sudo curl -sS https://getcomposer.org/installer | sudo php
sudo mv composer.phar /usr/local/bin/composer

# apache
sudo service apache2 stop

# git
sudo apt-get install -y git

# elasticsearch
sudo apt-get install -y openjdk-7-jre
curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.3.3.deb
sudo dpkg -i elasticsearch-2.3.3.deb
sudo /etc/init.d/elasticsearch start

# import elasticsearch megacorp test index
sh /vagrant/ed/elasticsearch/megacorp.sh
