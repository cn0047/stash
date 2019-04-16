docker-debian
-

# debian
docker build -t cn007b/debian:jessie docker/jessie
# check
docker run -ti --rm cn007b/debian:jessie vim --version
docker run -ti --rm cn007b/debian:jessie git --version
docker run -ti --rm cn007b/debian:jessie colordiff -v
docker run -ti --rm cn007b/debian:jessie curl -V
docker run -ti --rm cn007b/debian:jessie telnet --help
docker run -ti --rm cn007b/debian:jessie unzip -v
docker run -ti --rm cn007b/debian:jessie make -v
docker run -ti --rm cn007b/debian:jessie ab -V
docker run -ti --rm cn007b/debian:jessie lsb_release -c -s
# push
docker push cn007b/debian:jessie

# latest
docker build -t cn007b/debian:latest docker/jessie
docker push cn007b/debian:latest
