docker-alpine
-

# alpine
version=3.12
version=3.15
version=3.16
version=3.18
docker build -t cn007b/alpine:$version docker/$version

# # check
docker run -it --rm cn007b/alpine:$version sh -c 'hash vim'
docker run -it --rm cn007b/alpine:$version sh -c 'hash tree'
docker run -it --rm cn007b/alpine:$version sh -c 'hash mc'
docker run -it --rm cn007b/alpine:$version sh -c 'hash git'
docker run -it --rm cn007b/alpine:$version sh -c 'hash hg'
docker run -it --rm cn007b/alpine:$version sh -c 'hash colordiff'
docker run -it --rm cn007b/alpine:$version sh -c 'hash curl'
docker run -it --rm cn007b/alpine:$version sh -c 'hash wget'
# docker run -it --rm cn007b/alpine:$version net-tools
docker run -it --rm cn007b/alpine:$version sh -c 'hash nmap'
docker run -it --rm cn007b/alpine:$version sh -c 'hash telnet'
# docker run -it --rm cn007b/alpine:$version sh -c 'hash ftp'
docker run -it --rm cn007b/alpine:$version sh -c 'hash lftp'
docker run -it --rm cn007b/alpine:$version sh -c 'hash zip'
docker run -it --rm cn007b/alpine:$version sh -c 'hash unzip'
docker run -it --rm cn007b/alpine:$version sh -c 'hash rsync'
docker run -it --rm cn007b/alpine:$version sh -c 'hash make'
# docker run -it --rm cn007b/alpine:$version sh -c 'hash lsb_release'
# docker run -it --rm cn007b/alpine:$version sh -c 'hash uuid'
docker run -it --rm cn007b/alpine:$version sh -c 'hash jq'
#
docker run -it --rm cn007b/alpine:$version sh -c 'hash aws'
#
docker run -it --rm cn007b/alpine:$version sh -c 'hash ab'
#
# docker run -it --rm cn007b/alpine:$version sh -c 'hash dstat'
docker run -it --rm cn007b/alpine:$version sh -c 'hash mpstat'
docker run -it --rm cn007b/alpine:$version sh -c 'hash iotop'
# docker run -it --rm cn007b/alpine:$version sh -c 'hash cbm'

docker run -it --rm cn007b/alpine:$version sh

# # push
docker push cn007b/alpine:$version
# latest
docker tag cn007b/alpine:$version cn007b/alpine:latest
docker push cn007b/alpine:latest
