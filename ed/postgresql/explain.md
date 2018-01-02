Explain
-

[demo db](http://bit.ly/pagilia-dl)

ANALYZE
BUFFERS

In output:

* `cost` - time expended before output phase begin,
   after .. number - is estimated total cost.
* `rows` - estimated rows
* `width` - what we got as result dividet to number of rows

`Seq Scan` - not very good;
`Gather ... Parallel Seq Scan` - parallel query (must be not bad);
`Filter` - when field in where clause has NOT index;
`Index Cond` - when field in where clause has index;
`Index San using name_of_key` - good;
`Heap Fetches` - very good;

Stats about index usage:

````
SELECT relpages AS 'Disk page read', reltuples AS 'Rows scanned'
FROM pg_class WHERE relname = 'filmTable';
````
