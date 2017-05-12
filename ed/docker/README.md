docker
-
1.12.5

## Overview

````
# build image from Dokerfile in dir ./examples/mydockerbuild
# `-t` it is tag
# `.` it is current directory 
docker build -t docker-whale .

# run
docker run docker-whale

docker pull ubuntu

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

docker-machine ip

# information about all the containers
docker ps -a

# shows the standard output of a container
docker logs

# stop the running container
docker stop

docker start nostalgic_morse

# remove the web application container
docker rm nostalgic_morse

# delete
docker rmi -f docker-whale

exit
````

````
# show containers & images
docker images

docker exec -it happy_babbage
docker exec -it happy_babbage bash
cd /opt/docker/etc/supervisor.d
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
docker-compose run php-cli php /gh/x.php
docker-compose run mysql /bin/bash

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
