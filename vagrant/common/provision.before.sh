#!/usr/bin/env bash

# ubuntu
sudo apt-get update

# composer
sudo curl -sS https://getcomposer.org/installer | sudo php
sudo mv composer.phar /usr/local/bin/composer

# apache
sudo service apache2 stop

# git
sudo apt-get install -y git
