# docker-ubuntu

# ubuntu
docker build -t cn007b/ubuntu:17.10 docker/17.10
# test
docker run -ti --rm cn007b/ubuntu:17.10 vim --version
docker run -ti --rm cn007b/ubuntu:17.10 git --version
docker run -ti --rm cn007b/ubuntu:17.10 colordiff -v
docker run -ti --rm cn007b/ubuntu:17.10 curl -V
docker run -ti --rm cn007b/ubuntu:17.10 telnet --help
docker run -ti --rm cn007b/ubuntu:17.10 unzip -v
docker run -ti --rm cn007b/ubuntu:17.10 ab -V
# push
docker push cn007b/ubuntu:17.10

# protobuf-2
docker build -t cn007b/ubuntu:17.10-protobuf-2 docker/17.10-protobuf-2
# test
docker run -ti --rm cn007b/ubuntu:17.10-protobuf-2 protoc --version
# push
docker push cn007b/ubuntu:17.10-protobuf-2

# protobuf-3
docker build -t cn007b/ubuntu:17.10-protobuf-3 docker/17.10-protobuf-3
# test
docker run -ti --rm cn007b/ubuntu:17.10-protobuf-3 protoc --version
# push
docker push cn007b/ubuntu:17.10-protobuf-3

# latest
docker build -t cn007b/ubuntu:latest docker/17.10-protobuf-3
docker push cn007b/ubuntu:latest
