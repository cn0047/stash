FROM php:7.1-cli

RUN apt-get update \
    && apt-get install -y libmcrypt-dev libicu-dev mysql-client \
    && docker-php-ext-install pdo pdo_mysql

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring mcrypt \
    && pecl install uuid \
    && docker-php-ext-enable uuid

CMD ["php"]
