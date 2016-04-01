#!/usr/bin/env bash

apt-get update

# php
add-apt-repository ppa:ondrej/php5-5.6
apt-get update
apt-get install -y python-software-properties
apt-get update
apt-get install -y php5

# git
apt-get install -y git

# repo
rm -r /var/www/html
git clone https://github.com/cn007b/beeGame.git /var/www/html
