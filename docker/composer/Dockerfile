FROM composer/composer:latest

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev \
    && docker-php-ext-install -j$(nproc) bcmath \
    && pecl install uuid \
    && docker-php-ext-enable uuid
