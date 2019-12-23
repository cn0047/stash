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

Memcache - in-memory key-value store, which can store string up to 1MB.

Avoid Memcache hot keys.
Hot keys are a common anti-pattern that can cause Memcache capacity to be exceeded.

For Dedicated Memcache, it's recommend that the peak access rate on a single key
should be 1-2 orders of magnitude less than the per-GB rating.
For example, the rating for 1 KB sized items is 10,000 operations per second per GB of Dedicated Memcache.
Therefore, the load on a single key should not be higher
than 100 - 1,000 operations per second for items that are 1 KB in size.

When memcache overflows, it will expire oldest keys and flush them (eviction).

In php `$memcached->getStats()` will return info about: usage, items, total items, get hits and misses etc.
