#!/usr/bin/env bash

# ubuntu
sudo apt-get update

# apache
sudo service apache2 stop

# nginx
sudo apt-get install -y nginx

# git
sudo apt-get install -y git

# nodejs
sudo apt-get install python-software-properties
curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
sudo apt-get install nodejs

# beanstalk
sudo apt-get install -y beanstalkd

# # sqlite
# sudo apt-get install sqlite3 libsqlite3-dev

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

# rabbit
sudo apt-get install -y rabbitmq-server

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
