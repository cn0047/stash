docker
-
<br>17.09.1
<br>1.12.5

## Overview

Containers should be ephemeral (can be stopped and destroyed and a new one built and put in place).
<br>Hence container must be stateless.
<br>Each container should have only one concern.

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
# `-p` maps any required network ports
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

## Dockerfile

* FROM
* MAINTAINER
* RUN
* COPY
* WORKDIR
* ENTRYPOINT
* EXPOSE
* ENV
* VOLUME

## Machine

Use [machine](https://docs.docker.com/machine) to create Docker hosts on your local box,
on your company network, in your data center,
or on cloud providers like AWS or Digital Ocean.

````
docker-machine version

docker-machine ls

docker-machine create --driver virtualbox manager1

docker-machine ip default
docker-machine status default
docker-machine stop default
````

## Compose

[Compose](https://docs.docker.com/compose) is a tool
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

Non-finished swarm:

````
# init rabbit

#
# docker run -it --rm --name php-cli-rabbitmq-c -v $PWD/ed:/gh/ed --link rabbit php-cli

#
# docker-machine create --driver virtualbox manager1
# docker-machine ip manager1
# 192.168.99.100
# docker-machine ssh manager1
docker swarm init --advertise-addr 192.168.99.100
docker node ls
#
# docker service create --replicas 3 --name php-cli-rabbitmq-swarm php-cli-rabbitmq-c \
#     php /gh/ed/php/examples/rabbitmq/tutorials/workQueue/worker.php
docker service ls
# docker service ps php-cli-rabbitmq-swarm

#
# docker-machine create --driver virtualbox worker1
# docker-machine ssh worker1
docker swarm join \
    --token SWMTKN-1-3ie1vxyfmvh1756tv37dyp8datyyfcsfrnkhmzofwk3nsle7ud-cblwlb51iv251evjiudwxs6li \
    192.168.99.100:2377
docker node ls
````

## Network

````
# list these networks
docker network ls

docker network create --driver bridge x_node_mongo
docker network inspect x_node_mongo
````

https://app.pluralsight.com/library/courses/getting-started-kubernetes/table-of-contents
