nginx
-

````
nginx -t        # check config file
nginx -s reload # reload only configs
````

/etc/nginx/sites-enabled/app.conf
````
location / {
    access_log /var/log/nginx/access.log;
}
````
