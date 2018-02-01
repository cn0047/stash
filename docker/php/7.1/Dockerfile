FROM ubuntu:17.10

RUN apt-get update && apt-get install -y git vim curl wget apt-utils telnet net-tools unzip libpcre3-dev

# php
RUN apt-get install -y \
    php7.1 php7.1-cli php7.1-opcache php7.1-common php7.1-phpdbg php7.1-dev \
    php7.1-mcrypt php7.1-mbstring php7.1-bcmath php7.1-gd php7.1-zip php7.1-curl php7.1-imagick \
    php7.1-imap php7.1-ldap \
    php7.1-dom php7.1-xml php7.1-json \
    php7.1-mysql php7.1-pdo php-mongodb php7.1-sqlite php7.1-sqlite3 php7.1-pgsql \
    php7.1-redis \
    php7.1-memcache \
    libzmq5 libzmq5-dev php-zmq \
    php-pear php-soap php-ssh2

# protobuf
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
    && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3 \
    && mv protoc3/bin/* /usr/local/bin/ \
    && mv protoc3/include/* /usr/local/include/ \
    && chown $(whoami) /usr/local/bin/protoc \
    && chown -R $(whoami) /usr/local/include/google \
    && pecl install protobuf-3.5.1 \
    && echo 'extension=protobuf.so' >> /etc/php/7.1/mods-available/protobuf.ini \
    && rm protoc-3.5.1-linux-x86_64.zip

# xdebug
RUN apt-get install -y php-xdebug \
    && echo 'xdebug.remote_enable=1'            >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.remote_autostart=1'         >> /etc/php/7.1/mods-available/xdebug.ini \
    # IMPORTANT: On OSX this config value must be 0.
    && echo 'xdebug.remote_connect_back=0'      >> /etc/php/7.1/mods-available/xdebug.ini \
    # IMPORTANT: On OSX run next command:
    # sudo ifconfig lo0 alias 10.254.254.254
    && echo 'xdebug.remote_host=10.254.254.254' >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.remote_port=9002'           >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.idekey=PHPSTORM'            >> /etc/php/7.1/mods-available/xdebug.ini

# php conf
RUN echo 'display_errors=On'              >> /etc/php/7.1/cli/php.ini \
    && echo 'display_startup_errors=On'   >> /etc/php/7.1/cli/php.ini \
    && echo 'track_errors=On'             >> /etc/php/7.1/cli/php.ini \
    && echo 'log_errors=On'               >> /etc/php/7.1/cli/php.ini \
    && echo 'error_log=/dev/stdout'       >> /etc/php/7.1/cli/php.ini

# for Yii 1 template engine
RUN echo 'short_open_tag=On' >> /etc/php/7.1/cli/php.ini

# For big files upload.
RUN echo 'upload_max_filesize=50M' >> /etc/php/7.1/cli/php.ini
