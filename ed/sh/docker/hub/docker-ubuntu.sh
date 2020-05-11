docker-ubuntu
-

# ubuntu
version=17.10
version=18.04
version=18.04-video
version=20.04
docker build -t cn007b/ubuntu:$version docker/$version

# check
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash vim'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash tree'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash mc'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash git'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash colordiff'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash curl'
# docker run -ti --rm cn007b/ubuntu:$version net-tools
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash telnet'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash ftp'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash nmap'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash unzip'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash make'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash lsb_release'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash uuid'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash jq'
# docker run -ti --rm cn007b/ubuntu:$version libpcre3-dev
#
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash docker'
#
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash aws'
#
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash ab'
#
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash dstat'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash iotop'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash cbm'
docker run -ti --rm cn007b/ubuntu:$version sh -c 'hash mpstat'

# video
docker run -ti --rm cn007b/ubuntu:$version ffmpeg -version \
  2>&1 1>/dev/null && if [[ $? -eq 0 ]]; then echo -e "\033[32mok\033[0m"; else echo -e "\033[31merr\033[0m"; fi

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
