# FROM ubuntu:latest
FROM ubuntu:17.10
# FROM ubuntu:16.04

ENV TERM xterm

RUN apt-get update

RUN apt-get install -y vim
RUN apt-get install -y git
RUN apt-get install -y colordiff

RUN apt-get install -y curl
RUN apt-get install -y net-tools
RUN apt-get install -y telnet

RUN apt-get install -y unzip

RUN apt-get install -y apache2-utils

# ?
RUN apt-get install -y libpcre3-dev

# ZeroMQ
RUN apt-get install -y libzmq5 libzmq5-dev 

# # nodejs v 8
# RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
# RUN apt-get install -y nodejs
# # nvm
# RUN curl https://raw.githubusercontent.com/creationix/nvm/master/install.sh | bash
