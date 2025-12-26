# Redis

tag=8.4
docker pull redis:$tag


# run redis
docker run -it --rm --net=xnet -p 6379:6379 --name xredis --hostname xredis redis:$tag
docker run -it --rm --net=xnet -p 6379:6379 --name xredis --hostname xredis \
    -v $PWD/.data/.docker/redis:/data redis:$tag

# check redis
docker exec -it xredis redis-cli
docker exec -it xredis redis-cli -h localhost -p 6379



#### Redis cluster (replication)

# master
docker run -it --rm --net=xnet -p 6378:6378 --name xredis-master-1 --hostname xredis-master-1 \
    -v $PWD/ed/sh/sh.redis/examples/replication.master-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

# slave
docker run -it --rm --net=xnet -p 6377:6377 --name xredis-slave-1 --hostname xredis-slave-1 \
    -v $PWD/ed/sh/sh.redis/examples/replication.slave-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

# test
docker exec -it xredis-master-1 redis-cli -h localhost -p 6378
docker exec -it xredis-slave-1 redis-cli -h localhost -p 6377
# and
docker exec -it xredis-slave-1 redis-cli -h localhost -p 6377 info | grep -E '(replica|master|slave|status)'



#### Redis cluster (sharding)

docker run -it --rm --net=xnet -p 6361:6361 --name xredis-shard-1 --hostname xredis-shard-1 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6362:6362 --name xredis-shard-2 --hostname xredis-shard-2 \
    -v $PWD/ed/d/redis/examples/shard-2.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6363:6363 --name xredis-shard-3 --hostname xredis-shard-3 \
    -v $PWD/ed/d/redis/examples/shard-3.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6364:6364 --name xredis-shard-4 --hostname xredis-shard-4 \
    -v $PWD/ed/d/redis/examples/shard-4.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6365:6365 --name xredis-shard-5 --hostname xredis-shard-5 \
    -v $PWD/ed/d/redis/examples/shard-5.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6366:6366 --name xredis-shard-6 --hostname xredis-shard-6 \
    -v $PWD/ed/d/redis/examples/shard-6.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker exec -it xredis-shard-1 redis-cli -h localhost -p 6361 --cluster create --cluster-replicas 1 \
  xredis-shard-1:6361 \
  xredis-shard-2:6362 \
  xredis-shard-3:6363 \
  xredis-shard-4:6364 \
  xredis-shard-5:6365 \
  xredis-shard-6:6366

# test
docker exec -it xredis-shard-1 redis-cli -c -h localhost -p 6361
docker exec -it xredis-shard-2 redis-cli -c -h localhost -p 6362
docker exec -it xredis-shard-3 redis-cli -c -h localhost -p 6363



#### Redis cluster (sharding) 2

docker run -it --rm --net=xnet -p 6361:6361 --name xredis-shard-1 --hostname xredis-shard-1 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6362:6361 --name xredis-shard-2 --hostname xredis-shard-2 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6363:6361 --name xredis-shard-3 --hostname xredis-shard-3 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6364:6361 --name xredis-shard-4 --hostname xredis-shard-4 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6365:6361 --name xredis-shard-5 --hostname xredis-shard-5 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker run -it --rm --net=xnet -p 6366:6361 --name xredis-shard-6 --hostname xredis-shard-6 \
    -v $PWD/ed/d/redis/examples/shard-1.redis.conf:/etc/redis.conf \
    redis:$tag redis-server /etc/redis.conf

docker exec -it xredis-shard-1 redis-cli -h localhost -p 6361 --cluster create --cluster-replicas 1 \
  xredis-shard-1:6361 \
  xredis-shard-2:6361 \
  xredis-shard-3:6361 \
  xredis-shard-4:6361 \
  xredis-shard-5:6361 \
  xredis-shard-6:6361

# test
docker exec -it xredis-shard-1 redis-cli -c -h localhost -p 6361
docker exec -it xredis-shard-2 redis-cli -c -h localhost -p 6361
docker exec -it xredis-shard-3 redis-cli -c -h localhost -p 6361
