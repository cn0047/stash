#!/usr/bin/env bash

yum install -y vim
yum install -y htop

# php 7
#rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
#yum install -y php70w php70w-fpm php70w-opcache
#yum install -y php70w-xml php70w-soap php70w-xmlrpc
#yum install -y php70w-mbstring php70w-json php70w-gd php70w-mcrypt
#service php-fpm restart

# php 5.5
rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
yum install -y php55w-cli php55w-fpm php55w-opcache php55w-xml
yum install -y php55w-devel.x86_64
yum install -y php55w-dba.x86_64
yum install -y php55w-mysql.x86_64
yum install -y php55w-pdo.x86_64
yum install -y php55w-pear.noarch
yum install -y php55w-pecl-apcu.x86_64
yum install -y php55w-pecl-apcu-devel.x86_64
yum install -y php55w-mbstring.x86_64

yum install -y php55w-common.x86_64



# nginx
yum install -y nginx
cp /vagrant/nginx.default.conf /etc/nginx/conf.d/default.conf
service nginx restart

#
cp /vagrant/MariaDB.repo /etc/yum.repos.d/MariaDB.repo
yum install -y MariaDB-server MariaDB-client
service mysql restart

# java
yum install -y java-1.7.0-openjdk
yum install -y java-1.7.0-openjdk-devel

# elasticsearch
rpm --import https://packages.elastic.co/GPG-KEY-elasticsearch
mkdir -p /etc/yum.repos.d/
cp /vagrant/elasticsearch.repo /etc/yum.repos.d/elasticsearch.repo
yum install -y elasticsearch
chkconfig --add elasticsearch
service elasticsearch restart

# nodejs