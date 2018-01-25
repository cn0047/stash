Memcached
-

````
sudo service memcached restart
````

````
telnet 0.0.0.0 11211

stats items

flush_all
````

When memcache overflows, it will expire oldest keys and flush them.

In php `$memcached->getStats()` will return info about: usage, items, total items, get hits and misses etc.
