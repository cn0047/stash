Memcached
-

````sh
sudo service memcached restart
````

````sh
memcached -n 16 -f 1.05 -L
# -n min chunk
# -f grow factor
# -L init memory
````

````sh
telnet 0.0.0.0 11211

stats items

flush_all
````

When memcache overflows, it will expire oldest keys and flush them (eviction).

In php `$memcached->getStats()` will return info about: usage, items, total items, get hits and misses etc.
