nginx
-

Nginx relying on an asynchronous, events-driven architecture.

Nginx spawns worker processes, each of which can handle thousands of connections.
The worker processes accomplish this by implementing a fast looping mechanism
that continuously checks for and processes events.

Nginx was created to be both a web server and a proxy server.

````sh
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

SSI (Server Side Includes) - the most frequent use of SSI
is to include the contents of one or more files into a web page.

````html
<!--#include virtual="../quote.txt" -->
````

osx:
````sh
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

http:

````sh
# for: Request URI Too Large
large_client_header_buffers 4 16k;

# for big response headers:
proxy_buffer_size   64k;
proxy_buffers   4 64k;
proxy_busy_buffers_size   64k;

worker_processes auto;
multi_accept on;

client_body_timeout 10;
client_header_timeout 10;
reset_timedout_connection on;

keepalive_timeout 70;
keepalive_requests 100;
send_timeout 2;

use epoll; # on linux

sendfile on; # for big static video files
tcp_nopush on;
tcp_nodelay on;
directio 10m;
limit_rate 196K;

accept_mutex on;

open_file_cache max=200000 inactive=20s;
open_file_cache_valid 30s;
open_file_cache_min_uses 2;
open_file_cache_errors on;

client_body_buffer_size 10K;
client_header_buffer_size 1k;
client_max_body_size 8m;
````

server:

````sh
gzip on;
gzip_comp_level 5; # 1 low, 9 high

# cache
etag on;
ssi on;

expires max;
expires 7d;
expires modified 3d;
expires off;

limit_req zone=flood;

error_page   404   /404.html;
error_page   403   /403.html;
error_page   405   =200 $uri;

proxy_connect_timeout 600;
proxy_send_timeout    600;
proxy_read_timeout    600;
send_timeout          600;
````

````sh
server {
    root /var/www/birthplace/app/web;
}

location / {
    access_log /var/log/nginx/access.log;
}

location /admin.php {
    deny all;
}
location ~ \.(js|css|png|jpg|gif|swf|ico|pdf|mov|fla|zip|rar)$ {
    try_files $uri =404;
}

# Redirect uri like gXh6UAA727XX
if ($request_uri ~* "gXh6UAA727XX") {
    return 301 https://another.site.com/sc.js?gXh6UAA727XX;
}

# Redirect no-www to www
if ($host = 'zii.dev') {
    rewrite  ^/(.*)$  $scheme://www.zii.com/$1  permanent;
}
# 1 more example:
rewrite ^(.+)$ /ru/$1 permanent;

# redirect
return 301 $scheme://somesite.com$request_uri;

server {
    root /var/www/somesite.com;
    location / {
        try_files $uri $uri/ /index.html;
    }
}
````

````sh
events {
    worker_connections 2048;
}
http {
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

http2:

````sh
server {
    listen 443 ssl http2;
}
````

For socket.io:

````sh
upstream io_nodes {
  # Load Balancing methods:
  # Round-robin               - default;
  # least_conn                - request sends to node with least count of requests
  # hash $scheme$request_uri; - 
  ip_hash;
  server 127.0.0.1:6001 weight=10;
  server 127.0.0.1:6002 weight=5;
  server 127.0.0.1:6003 max_fails=3 fail_timeout=30s;
  server 127.0.0.1:6004;
  server 192.0.0.1 backup;
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

PHP:

````sh
server {
    server_name localhost;
    error_log /dev/stdout debug; # warn, error crit, alert, emerg
    access_log /dev/stdout combined;
    root /app/;

    # for security
    limit_conn conn_limit_per_ip 10;
    limit_req zone=req_limit_per_ip burst=10 nodelay;

    # for: Request Entity Too Large
    client_max_body_size 32m; # default 1Mb

    location ~ \.php$ {
        try_files $uri =404;
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/run/php/php7.1-fpm.sock;
        fastcgi_index index.php;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param PATH_INFO $fastcgi_path_info;

        # for upstream timeout
        fastcgi_read_timeout 150;

        # for big response headers:
        fastcgi_buffer_size 32k;
        fastcgi_buffers 4 32k;
    }
}

server {
    server_name _;
    root /var/www/site;
    location / {
        try_files $uri $uri/ /index.php;
    }
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_pass unix:/tmp/phpcgi.socket;
    }
}
````

PHP Pool (reverse proxy):

````sh
upstream backend {
        server 10.10.0.5 fail_timeout=360s max_fails=2;
        server 10.10.0.6 fail_timeout=360s max_fails=2;
        server 10.10.0.7 fail_timeout=360s max_fails=2;
}

server {
  location ~* \.(php)$ {
        fastcgi_pass backend;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }
}
````

#### HTTPS

````sh
cd docker/nginx/https/

openssl genrsa -des3 -out ca.orig.key 4096
# password: aldfnvljs034ilsdkjf
openssl req -new -x509 -days 365 -key ca.orig.key -out ca.crt
echo 01 > ca.srl

openssl genrsa -des3 -out ia.key 4096
# password: WUIYlsdjlkio2309
openssl req -new -key ia.key -out ia.csr
# password: OIsldjflkwjedkL
openssl x509 -req -days 730 -in ia.csr -CA ca.crt -CAkey ca.orig.key -out ia.crt
echo 01 > ia.srl

MYDOMAIN="localhost"
openssl genrsa -des3 -out "$MYDOMAIN".orig.key 2048
# password: salkdfjlkjIOIEDFds
openssl rsa -in "$MYDOMAIN".orig.key -out "$MYDOMAIN".key
openssl req -new -key "$MYDOMAIN".key -out "$MYDOMAIN".csr
# password: TRYRTUYGjksdnlfksdle
openssl x509 -req -days 365 -in "$MYDOMAIN".csr -CA ia.crt -CAkey ia.key -out "$MYDOMAIN".crt

docker run -ti --rm --name nginx-html \
    -v $PWD/docker/nginx/https/html.conf:/etc/nginx/conf.d/default.conf \
    -v $PWD/docker/nginx/https/localhost.crt:/ssl/localhost.crt \
    -v $PWD/docker/nginx/https/localhost.key:/ssl/localhost.key \
    -v $PWD:/gh \
    -p 3443:443 nginx:latest

docker run -it --rm -v $PWD:/app -w /app -p 3443:443 \
  -v $PWD/docker/nginx/https/php-fpm.conf:/etc/nginx/conf.d/default.conf \
  -v $PWD/docker/nginx/https/localhost.crt:/ssl/localhost.crt \
  -v $PWD/docker/nginx/https/localhost.key:/ssl/localhost.key \
  -v $PWD:/gh \
  cn007b/php /bin/bash -c '
    service php7.1-fpm start;
    service nginx start;
    tail -f /dev/stdout
  '
````

````sh
server {
    listen 443;
    server_name localhost;

    ssl on;
    ssl_certificate /ssl/localhost.crt;
    ssl_certificate_key /ssl/localhost.key;
    # ssl_session_cache   shared:SSL:100m;
    # ssl_session_timeout 1h;

    location / {
        root /gh/ed/html/examples;
        index table.html;
    }
}
````
