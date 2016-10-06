#!/usr/bin/env bash

# UBUNTU

apt-get update


# php
add-apt-repository ppa:ondrej/php
apt-get update
apt-get install -y php7.0 php7.0-mysql php7.0-fpm

# apache
service apache2 stop

# nginx
apt-get install -y nginx
# cp /vagrant/nginx.default.conf /etc/nginx/sites-available/default
mkdir -p /var/www/html
service nginx restart
