docker-pi
-

d=ed/sh/docker/examples.Dockerfile

tag=ping
tag=pinger
f=$d/sh.$tag.Dockerfile

docker build -t cn007b/pi:$tag -f $f $d
# check
docker run -it --rm cn007b/pi:$tag

# push
docker push cn007b/pi:$tag
