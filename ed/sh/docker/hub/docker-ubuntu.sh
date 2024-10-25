docker-ubuntu
-

# ubuntu
version=17.10
version=18.04
version=18.04-video
version=20.04
version=20.04-protobuf-3
version=22.10
docker build -t cn007b/ubuntu:$version docker/$version

# check
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash vim'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash tree'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash mc'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash git'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash colordiff'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash curl'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash wget'
# docker run -it --rm cn007b/ubuntu:$version net-tools
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash telnet'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash ftp'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash nmap'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash unzip'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash make'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash pkg-config'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash lsb_release'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash uuid'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash jq'
# docker run -it --rm cn007b/ubuntu:$version libpcre3-dev
#
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash docker'
#
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash aws'
#
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash ab'
#
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash dstat'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash iotop'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash cbm'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash mpstat'
# check protobuf-3
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash protoc'
docker run -it --rm cn007b/ubuntu:$version sh -c 'hash clang-format'

# video
docker run -it --rm cn007b/ubuntu:$version ffmpeg -version \
  2>&1 1>/dev/null && if [[ $? -eq 0 ]]; then echo -e "\033[32mok\033[0m"; else echo -e "\033[31merr\033[0m"; fi

# push
docker push cn007b/ubuntu:$version
# latest
docker tag cn007b/ubuntu:$version cn007b/ubuntu:latest
docker push cn007b/ubuntu:latest
