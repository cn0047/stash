FROM ubuntu:17.10

RUN apt-get update && apt-get install -y git vim curl wget apt-utils net-tools unzip libpcre3-dev

# php
RUN apt-get install -y \
    php7.1 php7.1-fpm php7.1-cli php7.1-opcache php7.1-common php7.1-phpdbg php7.1-dev \
    php7.1-mcrypt php7.1-mbstring php7.1-bcmath php7.1-gd php7.1-zip php7.1-curl php7.1-imagick \
    php7.1-imap php7.1-ldap \
    php7.1-dom php7.1-xml php7.1-json \
    php7.1-mysql php7.1-pdo php-mongodb php7.1-sqlite php7.1-sqlite3 php7.1-pgsql \
    php7.1-redis \
    libzmq5 php-zmq \
    php-pear php-soap php-ssh2

# SPL_Types
RUN \
    # pecl install SPL_Types
    git clone https://github.com/esminis/php_pecl_spl_types.git \
    && cd php_pecl_spl_types \
    && phpize \
    && ./configure \
    &&  make install \
    && cd .. && rm -r php_pecl_spl_types \
    && echo 'extension=spl_types.so' >> /etc/php/7.1/mods-available/spl_types.ini

# xdebug
RUN apt-get install -y php-xdebug \
    && echo 'xdebug.remote_enable=1'                        >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.remote_autostart=1'                     >> /etc/php/7.1/mods-available/xdebug.ini \
    # IMPORTANT: On OSX this config value must be 0.
    && echo 'xdebug.remote_connect_back=0'                  >> /etc/php/7.1/mods-available/xdebug.ini \
    # IMPORTANT: On OSX run next command:
    # sudo ifconfig lo0 alias 10.254.254.254
    && echo 'xdebug.remote_host=10.254.254.254'             >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.remote_port=9002'                       >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.remote_log=/tmp/xdebug.log'             >> /etc/php/7.1/mods-available/xdebug.ini \
    && echo 'xdebug.idekey=PHPSTORM'                        >> /etc/php/7.1/mods-available/xdebug.ini

# protobuf
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
    && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3 \
    && mv protoc3/bin/* /usr/local/bin/ \
    && mv protoc3/include/* /usr/local/include/ \
    && chown $(whoami) /usr/local/bin/protoc \
    && chown -R $(whoami) /usr/local/include/google \
    && pecl install protobuf-3.5.1 \
    && echo 'extension=protobuf.so' >> /etc/php/7.1/mods-available/protobuf.ini

RUN service php7.1-fpm start

EXPOSE 9000

# composer
RUN curl -sS https://getcomposer.org/installer | php \
  && mv composer.phar /usr/local/bin/composer

CMD ['php', 'php-fpm']
