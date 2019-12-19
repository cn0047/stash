Redis
-

#### Redis

# init redis
docker run -it --rm --net=xnet -p 6379:6379 --name xredis --hostname xredis redis:latest

# check redis
docker exec -ti xredis redis-cli

#### Redis cluster

# master
docker run -it --rm --net=xnet -p 6378:6378 --name xredis-master --hostname xredis-master \
    -v $PWD/ed/sh/sh.redis/examples/replication.master.redis.conf:/etc/redis.conf \
    redis:latest redis-server /etc/redis.conf

# slave
docker run -it --rm --net=xnet -p 6377:6377 --name xredis-slave --hostname xredis-slave \
    -v $PWD/ed/sh/sh.redis/examples/replication.slave.redis.conf:/etc/redis.conf \
    redis:latest redis-server /etc/redis.conf

# test
docker exec -ti xredis-master redis-cli -h localhost -p 6378
docker exec -ti xredis-slave redis-cli -h localhost -p 6377
# and
docker exec -ti xredis-slave redis-cli -h localhost -p 6377 info | grep -E '(replica|master|slave|status)'
