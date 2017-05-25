docker
-
1.12.5

## Overview

````
docker pull ubuntu

# build image from Dokerfile in dir ./examples/mydockerbuild
# `-t` it is tag
# `.` it is current directory 
docker build -t docker-whale .

# run an interactive container 
# `-t` terminal
# `-i` interactive connection
docker run -t -i ubuntu:latest /bin/bash

# run a web application
# `-d` runs the container as daemon
# ` -P` maps any required network ports
# `-p`
docker run -d -p 8081:80 timber/ziipr
docker run -d -p 192.168.0.32:1111:1111 timber/ziipr

# exec
docker exec -it happy_babbage
docker exec -it happy_babbage bash
cd /opt/docker/etc/supervisor.d

docker-machine ip

# information about all the containers
docker ps -a

# shows the standard output of a container
docker logs

# show containers & images
docker images

# stop the running container
docker stop

docker start nostalgic_morse

# remove the web application container
docker rm nostalgic_morse

# delete
docker rmi -f docker-whale

exit
````

#### RUN

````
docker run -ti --rm -v $PWD:/app composer install

docker run -it --rm node:latest node -v
docker run -it --rm --name log -p 3000:3000 -v $PWD:/usr/src/app -w /usr/src/app node:latest node src/index.js

# ES cluster
docker run -it --rm -p 9200:9200 --name es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=true -Des.node.data=false
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1 elasticsearch:2.2 \
    elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=false -Des.node.data=true \
    -Des.discovery.zen.ping.unicast.hosts=es-master-1
````

## Machine

Use [machine/](https://docs.docker.com/machine) to create Docker hosts on your local box,
on your company network, in your data center,
or on cloud providers like AWS or Digital Ocean.

## Compose

[Compose](https://docs.docker.com/compose)) is a tool
for defining and running multi-container Docker applications.

````
# build the project and detached
docker-compose up -d

# shutdown/clean up
docker-compose down 
docker-compose down --volumes

# builds, (re)creates, starts, and attaches to containers for a service.
docker-compose up

# runs a one-time command against a service
docker-compose run

docker-compose exec php-cli php /gh/x.php
docker-compose exec mysql /bin/bash

# restart stopped container
docker-compose start
````

## Swarm

Docker [Swarm](https://docs.docker.com/swarm) is native clustering for Docker.
It turns a pool of Docker hosts into a single, virtual Docker host.

## Network

````
# list these networks
docker network ls

docker network inspect
````

https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/

https://docs.docker.com/compose/compose-file/
