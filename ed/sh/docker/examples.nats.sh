# NATS

# server
tag=2.10
docker run -it --rm -p 4222:4222 -p 6222:6222 -p 8222:8222 --hostname localhost --net=xnet --name xnats nats:$tag

# dashboard
open http://localhost:8222
