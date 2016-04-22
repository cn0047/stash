#!/usr/bin/env bash

# CENTOS

yum install -y vim
yum install -y htop

# php 5.5
# rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
# rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
# yum install -y php55w-cli php55w-fpm php55w-opcache php55w-xml

# php 5.6
# rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
# rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
# yum install -y php56w-fpm php56w-opcache

# php 7
rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
yum install -y php70w
yum install -y php70w-fpm
yum install -y php70w-opcache
yum install -y php70w-xml php70w-soap php70w-xmlrpc
yum install -y php70w-mbstring php70w-json php70w-gd php70w-mcrypt

service php-fpm restart

# nginx
yum install -y nginx
cp /vagrant/nginx.default.conf /etc/nginx/conf.d/default.conf
service nginx restart
