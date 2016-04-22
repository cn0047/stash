#!/usr/bin/env bash

# CENTOS

yum install -y vim
yum install -y htop

# java
yum install -y java-1.7.0-openjdk
yum install -y java-1.7.0-openjdk-devel

# elasticsearch
rpm --import https://packages.elastic.co/GPG-KEY-elasticsearch
mkdir -p /etc/yum.repos.d/
cp /vagrant/elasticsearch.repo /etc/yum.repos.d/elasticsearch.repo
yum install -y elasticsearch
chkconfig --add elasticsearch
# /bin/systemctl daemon-reload
# /bin/systemctl enable elasticsearch.service
service elasticsearch restart
