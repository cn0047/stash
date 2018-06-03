FROM phalconphp/php:7

MAINTAINER Vladimir Kovpak <cn007b@gmail.com>

# FROM phalconphp/php:7
# FROM phalconphp/php-fpm:7

# phalcon
RUN curl -s "https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh" | bash \
    && apt-get install -y php7.0-phalcon

CMD ["php"]



## php-phalcon
#docker build -t cn007b/php:7.1-phalcon ./docker/7.1-phalcon
#docker run -it --rm cn007b/php:7.1-phalcon php -v
#docker run -it --rm cn007b/php:7.1-phalcon php -r 'var_dump(Phalcon\Version::get());'




#curl -s "https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh" | bash
#apt-get install php7.0-phalcon
#apt-get install php7.1-phalcon

#https://github.com/phalcon/dockerfiles
#
##/usr/lib/php/20160303/xdebug.so
##/usr/lib/php/20160303/phalcon.so
#
#libtool: install: cp ./.libs/phalcon.so /app/cphalcon-3.2.1/build/php7/64bits/modules/phalcon.so
#libtool: install: cp ./.libs/phalcon.lai /app/cphalcon-3.2.1/build/php7/64bits/modules/phalcon.la
#
#
#docker run -it --rm -v $PWD:/app -w /app cn007b/php:7.1-phalcon bash
#docker run -it --rm -v $PWD:/app -w /app cn007b/php:7.1-phalcon php -v
#docker run -it --rm -v $PWD:/app -w /app cn007b/php:7.1-phalcon php -r 'var_dump(Phalcon\Version::get());'
#






#
#
#
#
#
#
#docker run -it --rm cn007b/php:7.1 bash
#
#apt-get install -y software-properties-common
#apt-add-repository -y ppa:phalcon/stable
#apt-get update
## apt-get upgrade
#apt-get install -y php7.0-phalcon
#
## apt-get install -qq php7.0-cli php7.0-common
#apt-cache search php7.0-phalcon
#apt-cache search php7.1-phalcon
#
#curl -s "https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh" |  bash
#
#RUN curl -sSLO https://codeload.github.com/phalcon/cphalcon/tar.gz/v3.2.1 \
#    && tar xvzf v3.2.1  \
#    && cd cphalcon-3.2.1/build \
#    && ./install \
#    && echo 'extension=phalcon.so' > /etc/php/7.1/mods-available/30-phalcon.ini \
#    && cd - && rm v3.2.1 && rm -rf cphalcon-3.2.1
#    cp /cphalcon-3.2.1/build/php7/64bits/modules/phalcon.so /usr/lib/php/20160303
