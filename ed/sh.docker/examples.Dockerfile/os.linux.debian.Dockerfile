FROM debian:stretch

MAINTAINER Vladimir Kovpak <cn007b@gmail.com>

RUN apt-get update --fix-missing
RUN apt-get install -y curl make lsb-release git

# go
ENV GOROOT /usr/local/go/
RUN curl -O https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz \
    && tar zxf go1.10.linux-amd64.tar.gz \
    && rm go1.10.linux-amd64.tar.gz \
    && mv go /usr/local \
    && ln -s /usr/local/go/bin/go /usr/bin/

# gcloud
RUN export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" && \
    echo "deb http://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" \
    | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - && \
    apt-get update -y && apt-get install google-cloud-sdk -y

# goapp
RUN apt-get install -y google-cloud-sdk-app-engine-go \
  && ln -s /usr/lib/google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp /usr/local/bin

# datastore emulator
RUN apt-get install -y openjdk-8-jdk
RUN apt-get install -y google-cloud-sdk-datastore-emulator
