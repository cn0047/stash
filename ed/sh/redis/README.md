Redis
-
<br>5.0.6
<br>3.0.6

[php client](https://github.com/phpredis/phpredis)
[docs](https://redis.io/documentation)
[commands](https://redis.io/commands)
[lua](https://redis.io/commands/eval)
[pipelining](https://redis.io/topics/pipelining)

#### redis-cli

````sh
redis-cli # connect to redis
redis-benchmark # shell tool for benchmarking
````

````sh
select 1 # select DB with index 1
config get databases

info keyspace  #
memory stats   # memory usage details
monitor        # listen all requests received by server in real time
latency doctor #

set mykey 'Hello'
set foo 'bar'

incr x # increment by 1, if key not exists will create new one

# EX seconds      - expire time
# PX milliseconds - expire time in milliseconds
# NX              - set if does not exist
# XX              - set only if it exists
set mk 'hello' nx px 30000

# set if key does not exist
setnx mk "World"

expire mk 10 # ttl 10 seconds
ttl mk

get mk
mget mk mykey foo

flushdb # delete data from current database

keys *

del mykey

debug object key
memory usage key
````

Redis - in-memory data structure store (persistent), used as database,
with server-side scripts (lua).

#### Data types:

* Strings - A String value can be at max 512 Megabytes in length.
You can:
Use Strings as atomic counters;
Append to strings;
Use Strings a random access vectors;

* Lists - simply lists of strings.
The max length of a list is 2^32 (more than 4 billion of elements per list).

* Sets - an unordered non repeating collections of Strings.
The max number of members in a set is 2^32.

* Hashes - maps between string fields and string values.
Every hash can store up to 2^32 field-value pairs.

* Sorted sets - similarly to Sets,
the difference is that every member of a Sorted Set is associated with score,
that is used in order to take the sorted set ordered.

* Bitmaps and HyperLogLogs

#### List (Linked List):

Can be at max 512 Megabytes in length.

````sh
lpush mylist 0     # insert at the head of the list
lpop mylist        # delete from head
rpush mylist A     # insert at the tail of the list
rpush mylist B
rpush mylist C
rpush mylist End
rpop mylist        # delete from tail
ltrim mylist 0 2   # trim from 0 element to 2nd element including (length will be 3)
lrange mylist 0 -1 # print all values in a list
llen mylist        # list length

# circular buffer, with length 3
lpush mylist 1
lpush mylist 2
lpush mylist 3
lpush mylist 4
lrange mylist 0 -1
ltrim mylist 0 2
````

#### Hash:

````sh
hmset user:1 username antirez birthyear 1977 verified 1

hget user:1 username
hget user:1 birthyear

# get all the fields and values in a hash
hgetall user:1

# add new values into hash
hmset user:1 country UA

# remove the specified field
hdel user:1 birthyear

# get all the fields
hkeys user:1

# get all the values
hvals user:1

# is field exists
hexists user:1 birthyear
````

#### Set:

````sh
sadd myset 1 2 3

# get all members
smembers myset

# add new values into set
sadd myset 4 5 6 7

# remove item form set
srem myset 6

# check is value present in set
sismember myset 3
````

#### Associative array:

````sh
sadd users:name Bob Sam Joe
sadd users:lastName Wilson Smith Unknown

# check is user present in array
sismember users:name Bob

# get users names
smembers users:name
````

#### Transaction

`MULTI`, `EXEC`, `DISCARD` and `WATCH` are the foundation of transactions in Redis.
Transaction is atomic.

No roll back.

````sh
set a 100
set b 200

multi
incrby a -50
incrby b 50
exec
````

#### Pipelining

For best performance - use pipelining.
Pipelining - to send multiple commands to the server without waiting for the replies at all,
and finally read the replies in a single step.
