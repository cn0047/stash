# Memcached


# run
docker run -it --rm --net=xnet -p 11211:11211 --hostname xmemcached --name xmemcached kmemcached
# check
docker exec -it xmemcached telnet 0.0.0.0 11211

# build
docker build -t kmemcached -f ed/sh/docker/examples.Dockerfile/cache.memcached.Dockerfile .
