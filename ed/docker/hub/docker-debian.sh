docker-debian
-

version=jessie
version=buster
version=stretch

# debian
docker build -t cn007b/debian:$version docker/$version
# check
docker run -ti --rm cn007b/debian:$version vim --version
docker run -ti --rm cn007b/debian:$version git --version
docker run -ti --rm cn007b/debian:$version colordiff -v
docker run -ti --rm cn007b/debian:$version curl -V
docker run -ti --rm cn007b/debian:$version telnet --help
docker run -ti --rm cn007b/debian:$version unzip -v
docker run -ti --rm cn007b/debian:$version make -v
docker run -ti --rm cn007b/debian:$version ab -V
docker run -ti --rm cn007b/debian:$version lsb_release -c -s
# push
docker push cn007b/debian:$version

# latest
docker build -t cn007b/debian:latest docker/$version
docker push cn007b/debian:latest
