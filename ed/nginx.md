nginx
-

````
nginx -t                 # check config file
nginx -s reload          # reload only configs
/etc/init.d/nginx reload
service nginx restart
````

osx:
````
config:
/usr/local/etc/nginx/nginx.conf
document root:
/usr/local/Cellar/nginx/1.8.0/html
````

linux:
````
config:
/etc/nginx/sites-enabled/app.conf
````

````
server {
    root /var/www/birthplace/app/web;
}
location / {
    access_log /var/log/nginx/access.log;
}
# Redirect uri like gXh6UAA727XX
if ($request_uri ~* "gXh6UAA727XX") {
    return 301 https://another.site.com/sc.js?gXh6UAA727XX;
}
````
