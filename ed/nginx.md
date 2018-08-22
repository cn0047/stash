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
worker_processes auto;
use epoll; # on linux
multi_accept on;
sendfile on;
tcp_nodelay on;
tcp_nopush on;
open_file_cache max=200000 inactive=20s;
open_file_cache_valid 30s;
open_file_cache_min_uses 2;
open_file_cache_errors on;
keepalive_timeout 70;
keepalive_requests 100;
gzip on;
reset_timedout_connection on;
client_body_timeout 10;
send_timeout 2;
client_max_body_size 1m;
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

# redirect
return 301 $scheme://somesite.com$request_uri;

server {
    root /var/www/somesite.com;
    location / {
        try_files $uri $uri/ /index.html;
    }
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

http2:

````
server {
    listen 443 ssl http2;
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

PHP:

````
server {
    server_name localhost;
    error_log /dev/stdout;
    access_log /dev/stdout;
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

http {
  # also for: Request URI Too Large
  large_client_header_buffers 4 16k;
}

# also for big response headers:
http {
  proxy_buffer_size   64k;
  proxy_buffers   4 64k;
  proxy_busy_buffers_size   64k;
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

PHP Pool:

````
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

````
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

````
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
