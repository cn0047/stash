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

`Index San using name_of_key` - good,
`Index Cond` - when field in where clause has index,
`Filter` - when field in where clause has NOT index.
`Heap Fetches` - very good.

Stats about index usage:

````
SELECT relpages AS 'Disk page read', reltuples AS 'Rows scanned'
FROM pg_class WHERE relname = 'filmTable';
````
