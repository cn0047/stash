FROM composer/composer:latest

RUN apt-get update

RUN apt-get install -y libmcrypt-dev libicu-dev mysql-client libpq-dev \
    && docker-php-ext-install pdo pdo_mysql pgsql pdo_pgsql

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev openssl zip unzip \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring mcrypt \
    && pecl install uuid \
    && docker-php-ext-enable uuid
