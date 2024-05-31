Administration
-

````sh
# redis with custom config:
redis-server /usr/local/etc/redis/redis.conf

redis-cli
redis-cli -h localhost -p 6379

# backup
redis-cli --rdb /tmp/dump.rdb
````

````sh
# info about cluster
# + info abou hits/misses
info

config get databases
````
