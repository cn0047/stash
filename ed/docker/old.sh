RUN apt-get update \
    && apt-get install -y git libmcrypt-dev libicu-dev mysql-client libpq-dev unzip \
    && docker-php-ext-install pdo pdo_mysql pgsql pdo_pgsql sockets

# uuid
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y uuid-dev \
    && docker-php-ext-install -j$(nproc) bcmath pdo mbstring mcrypt \
    && pecl install uuid \
    && docker-php-ext-enable uuid
