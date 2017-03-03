Redis
-

#### redis-cli

````
redis-cli # connect to redis
````
````
set mykey "Hello"
get mykey

flushdb # delete data from current database

keys *
LRANGE key 0 -1 # print all values in a list
````
