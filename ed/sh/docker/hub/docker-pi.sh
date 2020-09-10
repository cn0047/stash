docker-pi
-

d=ed/sh/docker/examples.Dockerfile

# ai
tag=ai
tag=ai.tf
tag=ai.tf.1.13
tag=ai.tf.2.2
tag=ai.cuda.0
tag=ai.cuda.0b
tag=ai.cuda.1
f=$d/$tag.Dockerfile
# sh
tag=ping
tag=pinger
tag=ffmpeg
f=$d/sh.$tag.Dockerfile

docker build -t cn007b/pi:$tag -f $f $d
# check
docker run -it --rm -v $PWD:/gh -w /gh cn007b/pi:$tag /bin/bash
# push
docker push cn007b/pi:$tag

# latest
docker tag cn007b/alpine cn007b/pi:latest
docker push cn007b/pi:latest
