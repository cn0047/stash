#!/usr/bin/env bash

# UBUNTU

apt-get update

# elasticsearch
sudo apt-get install openjdk-7-jre
curl -L -O https://download.elastic.co/elasticsearch/elasticsearch/elasticsearch-2.3.3.deb
sudo dpkg -i elasticsearch-2.3.3.deb
sudo /etc/init.d/elasticsearch start