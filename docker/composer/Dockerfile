FROM composer/composer:latest

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev openssl zip unzip \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring \
    && pecl install uuid \
    && docker-php-ext-enable uuid
