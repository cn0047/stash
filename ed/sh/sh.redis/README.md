Redis
-
<br>5.0.6
<br>3.0.6

[php client](https://github.com/phpredis/phpredis)
[docs](https://redis.io/documentation)
[commands](https://redis.io/commands)

#### redis-cli

````sh
redis-cli # connect to redis
redis-benchmark # shell tool for benchmarking
````

````sh
set mykey 'Hello'
set foo 'bar'

# EX seconds      - expire time
# PX milliseconds - expire time in milliseconds
# NX              - set if does not exist
# XX              - set only if it exists
set mk 'hello' nx px 30000

# set if key does not exist
setnx mykey2 "World"

expire foo 10 # ttl 10 seconds

get mykey
mget mykey foo

flushdb # delete data from current database

keys *

del mykey
````

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

````sh
rpush mylist A
rpush mylist B
rpush mylist End
lrange mylist 0 -1 # print all values in a list
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
