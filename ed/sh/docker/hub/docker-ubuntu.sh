docker-ubuntu
-

# ubuntu
version=17.10
version=18.04
version=18.04-video
docker build -t cn007b/ubuntu:$version docker/$version

# check
docker run -ti --rm cn007b/ubuntu:$version vim --version
docker run -ti --rm cn007b/ubuntu:$version tree --version
docker run -ti --rm cn007b/ubuntu:$version mc --version
docker run -ti --rm cn007b/ubuntu:$version git --version
docker run -ti --rm cn007b/ubuntu:$version colordiff -v
docker run -ti --rm cn007b/ubuntu:$version curl -V
# docker run -ti --rm cn007b/ubuntu:$version net-tools
docker run -ti --rm cn007b/ubuntu:$version telnet --help
docker run -ti --rm cn007b/ubuntu:$version ftp -help
docker run -ti --rm cn007b/ubuntu:$version nmap -V
docker run -ti --rm cn007b/ubuntu:$version unzip -v
docker run -ti --rm cn007b/ubuntu:$version make -v
docker run -ti --rm cn007b/ubuntu:$version lsb_release -c -s
docker run -ti --rm cn007b/ubuntu:$version uuid
# docker run -ti --rm cn007b/ubuntu:$version libpcre3-dev
#
docker run -ti --rm cn007b/ubuntu:$version docker version
#
docker run -ti --rm cn007b/ubuntu:$version aws --version
#
docker run -ti --rm cn007b/ubuntu:$version ab -V
#
docker run -ti --rm cn007b/ubuntu:$version dstat -h
docker run -ti --rm cn007b/ubuntu:$version cbm --version
docker run -ti --rm cn007b/ubuntu:$version mpstat # sysstat
# docker run -ti --rm cn007b/ubuntu:$version iotop

# video
docker run -ti --rm cn007b/ubuntu:$version ffmpeg -version \
  > /dev/null && if [[ $? -eq 0 ]]; then echo -e "\033[32mok\033[0m"; else echo -e "\033[31merr\033[0m"; fi

# protobuf-3
docker build -t cn007b/ubuntu:$version-protobuf-3 docker/$version-protobuf-3
# check
docker run -ti --rm cn007b/ubuntu:$version-protobuf-3 protoc --version
# push
docker push cn007b/ubuntu:$version-protobuf-3

# push
docker push cn007b/ubuntu:$version
# latest
docker tag cn007b/ubuntu:$version cn007b/ubuntu:latest
docker push cn007b/ubuntu:latest
