#!/usr/bin/env bash

apt-get update

# php
add-apt-repository ppa:ondrej/php5-5.6
apt-get update
apt-get install -y python-software-properties
apt-get update
apt-get install -y php5
apt-get install -y php5-fpm

# apache
service apache2 stop

# nginx
apt-get install -y nginx
# enable php at /etc/nginx/sites-available/default
