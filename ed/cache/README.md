cache
-

Cache - small amount of faster memory to improve the performance
of recently or frequently accessed data
that is stored temporarily in a rapidly accessible storage.

Cache algorithms:
LFU - Least Frequently Used
LRU - Least Recently Used
MRU - Most Recently Used

## OS Cache

OS Cache may be hardware or software.

Cache memory (aka CPU memory) is high-speed Static Random Access Memory (SRAM)
usually part of the CPU.

CPU Cache:
1) Level 1 Cache - (primary cache) relatively small (64K cache),
   built on the microprocessor chip.
2) Level 2 cache - (secondary cache) embedded on the CPU
   or on a separate chip or coprocessor.
3) Level 3 cache - (specialized memory).
4) Level 4 cache - accessed and shared by the CPU and GPU.

CPU Cache can be instruction cache or data cache.

GPU Cache.

Disk Cache - usually included as part of the hard disk.
A disk cache can also be a specified portion of random access memory (RAM).

WEB cache (browser cache, DNS cache).

## Memcached

* In-memory key-value store.
* Can store string up to 1MB.

## Redis

* In-memory data structure store, used as database.
* Can store string up to 512MB.
* Transactions - yes.
* Durability - yes.
* Server-side scripts - lua.

* *Partitioning methods - sharding.*
* *Replication methods - master-slave replication.*

## Redis vs Memcached

1. In operations with sets Redis exceed Memcached.
2. In operations with gets Redis loses to Memcached.
3. Memcached faster than Redis with multi gets at high volumes (more than 100K entries).
4. Redis and Memcached are similar on multi gets for small volumes (1000 or 100 entries).
