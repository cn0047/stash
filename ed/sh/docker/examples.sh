# Examples


sudo ifconfig lo0 alias 10.254.254.254

docker network create --driver bridge xnet

# php
tag=latest
tag=7.1-protobuf-3
tag=7.1
docker pull cn007b/php:$tag
docker tag cn007b/php:$tag xphp
# docker tag cn007b/php nphp
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -p 8080:80 xphp /bin/bash

# go
docker pull cn007b/go:latest
docker tag cn007b/go xgo
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xgo /bin/bash
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh' xgo /bin/bash

# python
docker pull cn007b/python:latest
docker tag cn007b/python xpy
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xpy /bin/bash



## Linux

# ubuntu
docker pull cn007b/ubuntu
docker tag cn007b/ubuntu xubuntu
docker tag cn007b/ubuntu:20.04-protobuf-3 xubuntu
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xubuntu /bin/bash

# debian
docker pull cn007b/debian
docker tag cn007b/debian xdebian
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xdebian /bin/bash

# alpine
docker pull cn007b/alpine
docker tag cn007b/alpine xalpine
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xalpine sh



## PI

# ai
docker run -it --rm -p 6006:6006 cn007b/pi:ai         /bin/bash
docker run -it --rm -p 6006:6006 cn007b/pi:ai.tf      /bin/bash
docker run -it --rm -p 6006:6006 cn007b/pi:ai.tf.1.13 /bin/bash
docker run -it --rm -p 6006:6006 cn007b/pi:ai.tf.2.2  /bin/bash

# sh
docker run -it --rm cn007b/pi:ping
docker run -it --rm cn007b/pi:pinger
docker run -it --rm -v $PWD:/gh -w /gh -v /tmp:/tmp cn007b/pi:ffmpeg  /bin/bash
