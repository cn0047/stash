Memcache
-

When memcache overflows, it will expire oldest keys and flush them.

````
sudo service memcached restart
````

In php `$memcached->getStats()` will return info about: usage, items, total items, get hits and misses etc.
