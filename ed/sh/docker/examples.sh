# Examples


sudo ifconfig lo0 alias 10.254.254.254

docker network create --driver bridge xnet

# php
docker pull cn007b/php:7.1-protobuf-3
docker tag cn007b/php:7.1-protobuf-3 xphp
docker pull cn007b/php
docker tag cn007b/php xphp
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
docker run -ti --rm --net=xnet -v $PWD:/gh -w /gh xpy /bin/bash

### Linux

# ubuntu
docker pull cn007b/ubuntu
docker tag cn007b/ubuntu xubuntu
docker run -ti --rm --net=xnet -v $PWD:/gh -w /gh xubuntu /bin/bash

# debian
docker pull cn007b/debian
docker tag cn007b/debian xdebian
docker run -ti --rm --net=xnet -v $PWD:/gh -w /gh xdebian /bin/bash

# alpine
docker pull cn007b/alpine
docker tag cn007b/alpine xalpine
docker run -ti --rm --net=xnet -v $PWD:/gh -w /gh xalpine sh
