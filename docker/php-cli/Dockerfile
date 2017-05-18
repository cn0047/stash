FROM php:7.1-cli

RUN apt-get update \
    # && apt-get install -y libmcrypt-dev libicu-dev mysql-client \
    && docker-php-ext-install pdo pdo_mysql

CMD ["php"]
