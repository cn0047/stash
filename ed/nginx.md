nginx
-

Nginx relying on an asynchronous, events-driven architecture.

Nginx spawns worker processes, each of which can handle thousands of connections.
The worker processes accomplish this by implementing a fast looping mechanism
that continuously checks for and processes events.

Nginx was created to be both a web server and a proxy server.

````
sudo nginx -t        # check config file
sudo nginx -s reload # reload only configs
sudo nginx -s stop
sudo nginx -s quit
sudo nginx -s reopen
sudo nginx -c /usr/local/etc/nginx/ni.nginx.conf

sudo s/etc/init.d/nginx reload
sudo service nginx restart

brew services start nginx
````

osx:
````
# config:
/usr/local/etc/nginx/nginx.conf

# document root:
/usr/local/Cellar/nginx/1.8.0/html
/usr/local/Cellar/nginx/1.12.1/html

mkdir -p /usr/local/etc/nginx/conf.d
mkdir -p /usr/local/etc/nginx/sites-enabled
mkdir -p /usr/local/etc/nginx/sites-available
touch    /usr/local/etc/nginx/sites-available/default
ln -sfv  /usr/local/etc/nginx/sites-available/default /usr/local/etc/nginx/sites-enabled/default

# logs:
mkdir -p /usr/local/etc/nginx/logs
touch /usr/local/etc/nginx/logs/access.log
touch /usr/local/etc/nginx/logs/error.log
# or
tail -f /usr/local/var/log/nginx/access.log
tail -f /usr/local/var/log/nginx/error.log
````

linux:
````
config:
/etc/nginx/sites-enabled/app.conf
````

#### [Congif](http://nginx.org/en/docs/ngx_core_module.html)

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

````
events {
    worker_connections 2048;
}
http {
    gzip on;
    server_tokens off; # Delete X-Powered-By
    add_header X-Frame-Options Deny;
    server {
        listen 80;
        server_name z.dev;
        location ~ \. (xml|ini) {
            deny all;
        }
        location / {
            proxy_pass http://192.168.56.101:3000;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host $host;
            proxy_cache_bypass $http_upgrade;
        }
    }
}
````

For socket.io:

````
upstream io_nodes {
  ip_hash;
  server 127.0.0.1:6001;
  server 127.0.0.1:6002;
  server 127.0.0.1:6003;
  server 127.0.0.1:6004;
}

server {
  listen 3000;
  server_name io.yourhost.com;
  location / {
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $host;
    proxy_http_version 1.1;
    proxy_pass http://io_nodes;
  }
}
````
