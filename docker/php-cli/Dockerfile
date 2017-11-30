FROM php:7.1-cli

RUN apt-get update \
    && apt-get install -y git libmcrypt-dev libicu-dev mysql-client libpq-dev \
    && docker-php-ext-install pdo pdo_mysql pgsql pdo_pgsql sockets

# RUN pecl install xdebug && docker-php-ext-enable xdebug \
#     && echo "zend_extension=$(find /usr/local/lib/php/extensions/ -name xdebug.so)" > /usr/local/etc/php/conf.d/xdebug.ini \
#     && echo "xdebug.profiler_enable=on" >> /usr/local/etc/php/conf.d/xdebug.ini \
#     && echo "xdebug.profiler_output_nam=xdebug.profiler.%p" >> /usr/local/etc/php/conf.d/xdebug.ini \
#     && echo "xdebug.profiler_output_dir=/app" >> /usr/local/etc/php/conf.d/xdebug.ini

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring mcrypt \
    && pecl install uuid \
    && docker-php-ext-enable uuid

RUN curl -sS https://getcomposer.org/installer | php \
  && mv composer.phar /usr/local/bin/composer

CMD ["php"]
