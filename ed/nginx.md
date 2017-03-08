nginx
-

````
sudo nginx -t                 # check config file
sudo nginx -s reload          # reload only configs
sudo s/etc/init.d/nginx reload
sudo service nginx restart
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
# Redirect no-www to www
if ($host = 'ziipr.dev') {
    rewrite  ^/(.*)$  $scheme://www.ziipr.com/$1  permanent;
}
````

Nginx

````
server {
    listen 80;
    server_name z.dev;
    location / {
        proxy_pass http://192.168.56.101:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
````
