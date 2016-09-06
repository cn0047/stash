#!/usr/bin/env bash

# UBUNTU
apt-get update

# php
add-apt-repository ppa:ondrej/php
apt-get update
apt-get install -y php7.0 php7.0-mysql php7.0-fpm php7.0-xml
service php7.0-fpm restart

# apache
service apache2 stop

# nginx
apt-get install -y nginx
cp /vagrant/vagrant/nginx.default.conf /etc/nginx/sites-available/default
service nginx restart

# symfony
curl -LsS https://symfony.com/installer -o /usr/local/bin/symfony
chmod a+x /usr/local/bin/symfony

# composer
php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
php -r "if (hash_file('SHA384', 'composer-setup.php') === 'e115a8dc7871f15d853148a7fbac7da27d6c0030b848d9b3dc09e2a0388afed865e6a3d6b3c0fad45c48e2b5fc1196ae') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
php composer-setup.php
php -r "unlink('composer-setup.php');"
mv composer.phar /usr/local/bin/composer
