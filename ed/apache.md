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
DocumentRoot /home/kovpak/web/kovpak/gh/ed/phalcon/examples/one/public/
<Directory /home/kovpak/web/kovpak/gh/>
    Options Indexes FollowSymLinks Includes ExecCGI
    Require all granted
    AllowOverride All
    Order deny,allow
    Allow from all
</Directory>
````

````
rm -rf /var/www
ln -fs /vagrant /var/www
````
