Apache
-

````
sudo apt-get install tasksel
sudo tasksel install lamp-server

sudo service httpd restart
sudo service apache2 restart

/usr/sbin/apache2 -v
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
-RewriteCond %{HTTP_HOST} www\.upforit\.com [NC]
-RewriteCond %{REQUEST_URI} ^/aff.php$
-RewriteCond %{QUERY_STRING} ^dynamicpage=find&filter=sexy&a_bid=48c701ef&a_aid=e2ddc951
-RewriteRule ^(.*) - [F]
````

````
rm -rf /var/www
ln -fs /vagrant /var/www
````
