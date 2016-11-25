#!/usr/bin/env bash

. vagrant/common/provision.before.sh

###############################################################################
# GENERAL
###############################################################################

# php 5.6
sudo add-apt-repository ppa:ondrej/php
sudo apt-get update
sudo apt-get install -y php5.6 php5.6-fpm php5.6-imap php5.6-mongo

# nginx
sudo cp /vagrant/vagrant/ubuntu-14.04/nginx.conf /etc/nginx/sites-available/default


. vagrant/common/provision.after.sh
