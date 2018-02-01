FROM phalconphp/php:7

# FROM phalconphp/php:7
# FROM phalconphp/php-fpm:7

# phalcon
RUN curl -s "https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh" | bash \
    && apt-get install -y php7.0-phalcon

CMD ["php"]
