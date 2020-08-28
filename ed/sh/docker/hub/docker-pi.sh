docker-pi
-

d=ed/sh/docker/examples.Dockerfile

tag=ping
tag=pinger
f=$d/sh.$tag.Dockerfile
#
tag=ai.tf
tag=ai.tf.1.13
tag=ai.tf.2.2
tag=ai.cuda.0
tag=ai.cuda.0b
tag=ai.cuda.1
f=$d/$tag.Dockerfile

docker build -t cn007b/pi:$tag -f $f $d
# check
docker run -it --rm -v $PWD:/gh -w /gh cn007b/pi:$tag /bin/bash
docker run -it --rm -v $PWD:/gh -w /gh cn007b/pi:$tag sh -c 'pip3 list | grep tensor'
# push
docker push cn007b/pi:$tag
