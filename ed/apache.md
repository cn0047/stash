Apache
-

````
sudo apt-get install tasksel
sudo tasksel install lamp-server

sudo service apache2 restart

sudo apt-get install php5-mcrypt
sudo apt-get install php5-curl
````

````
sudo a2enmod vhost_alias
sudo a2enmod rewrite

sudo a2dissite default
sudo a2ensite mysite

sudo php5enmod mcrypt
````
