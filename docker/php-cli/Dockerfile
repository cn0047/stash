# FROM php:5.6-cli
# FROM php:7.1-cli
FROM php:7.1

RUN apt-get update \
    && apt-get install -y git libmcrypt-dev libicu-dev mysql-client libpq-dev unzip \
    && docker-php-ext-install pdo pdo_mysql pgsql pdo_pgsql sockets

# uuid
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring mcrypt \
    && pecl install uuid \
    && docker-php-ext-enable uuid

# xdebug
RUN pecl install xdebug && docker-php-ext-enable xdebug \
    && echo 'xdebug.remote_enable=1'                        >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo 'xdebug.remote_autostart=1'                     >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo 'xdebug.remote_connect_back=1'                  >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    # && echo 'xdebug.remote_host=localhost'                  >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    # && echo 'xdebug.remote_host=172.17.0.1'                 >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo 'xdebug.remote_port=9001'                       >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    && echo 'xdebug.remote_log=/tmp/xdebug.log'             >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \
    # && echo "xdebug.profiler_enable=off"                    >> /usr/local/etc/php/conf.d/xdebug.ini \
    # && echo "xdebug.profiler_output_nam=xdebug.profiler.%p" >> /usr/local/etc/php/conf.d/xdebug.ini \
    # && echo "xdebug.profiler_output_dir=/app"               >> /usr/local/etc/php/conf.d/xdebug.ini \
    && echo 'xdebug.idekey=PHPSTORM'                        >> /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini

# protobuf
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
    && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3 \
    && mv protoc3/bin/* /usr/local/bin/ \
    && mv protoc3/include/* /usr/local/include/ \
    && chown $(whoami) /usr/local/bin/protoc \
    && chown -R $(whoami) /usr/local/include/google \
    && pecl install protobuf-3.5.1 \
    && echo extension = protobuf.so >> /usr/local/etc/php/php.ini

# composer
RUN curl -sS https://getcomposer.org/installer | php \
  && mv composer.phar /usr/local/bin/composer

CMD ["php"]
