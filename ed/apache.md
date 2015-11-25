Apache
-

````
sudo apt-get install tasksel
sudo tasksel install lamp-server

sudo service apache2 restart
````

````
sudo a2enmod vhost_alias
sudo a2enmod rewrite

sudo a2dissite default
sudo a2ensite mysite

sudo php5enmod mcrypt
````

````
rm -rf /var/www
ln -fs /vagrant /var/www
````
