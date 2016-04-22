#!/usr/bin/env bash

# UBUNTU

apt-get update

# php
add-apt-repository ppa:ondrej/php
apt-get update
apt-get install -y php7.0
apt-get install -y php7.0-mysql
apt-get install -y php7.0-fpm
