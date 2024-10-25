# NATS

tag=2.10
docker run -it --rm -p 4222:4222 -p 6222:6222 -p 8222:8222 --hostname localhost --name xnats nats:$tag
