#!/usr/bin/env bash

yum install -y vim
yum install -y htop

# php
rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
rpm -Uvh https://mirror.webtatic.com/yum/el6/latest.rpm
# php 5.5
yum install -y php55w-cli php55w-fpm php55w-opcache php55w-xml
# php 5.6
# yum install -y php56w-fpm php56w-opcache
service php-fpm restart

# nginx
yum install -y nginx
cp /vagrant/nginx.default.conf /etc/nginx/conf.d/default.conf
service nginx restart
