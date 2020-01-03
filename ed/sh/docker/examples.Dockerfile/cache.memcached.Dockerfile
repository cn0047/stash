FROM cn007b/ubuntu:17.10

MAINTAINER Vladimir Kovpak <cn007b@gmail.com>

RUN apt-get update \
    && apt-get install -y memcached \
    && sed -i 's/127.0.0.1/0.0.0.0/' /etc/memcached.conf

ENTRYPOINT service memcached start && tail -f /var/log/memcached.log
