#!/usr/bin/env bash

. vagrant/common/provision.before.sh

###############################################################################
# GENERAL
###############################################################################

# php 5.6
sudo add-apt-repository ppa:ondrej/php
sudo apt-get update
sudo apt-get install -y php5.6 php5.6-cli php5.6-fpm php5.6-mongo php5.6-common php5.6-dev
sudo apt-get install -y php5.6-mcrypt php5.6-mbstring
sudo apt-get install -y php5.6-mysql php5.6-pdo
sudo apt-get install -y php5.6-dom php5.6-xml php5.6-json
sudo apt-get install -y php5.6-zip  php5.6-curl php5.6-gd php5.6-imap

# nginx
sudo cp /vagrant/vagrant/ubuntu-14.04/nginx.conf /etc/nginx/sites-available/default

###############################################################################
# SYMFONY
###############################################################################

#bulletinBoard
cd /vagrant/ed/symfony/examples/bulletinBoard/ && sudo php app/console cache:clear --env=prod --no-debug
cd /vagrant/ed/symfony/examples/bulletinBoard/ && sudo php app/console assetic:dump --env=prod --no-debug
composer update -o







. vagrant/common/provision.after.sh
