Explain
-

[demo db](http://bit.ly/pagilia-dl)

ANALYZE
BUFFERS

In output:
* `cost` - time expended before output phase begin,
   after .. number - is estimated total cost.
* `rows` - estimated rows.
* `width` - what we got as result divided to number of rows.

<br>`Seq Scan` - not very good;
<br>`Gather ... Parallel Seq Scan` - parallel query (must be not bad);
<br>`Filter` - when field in where clause has NOT index;
<br>`Index Cond` - when field in where clause has index;
<br>`Index San using name_of_key` - good;
<br>`Heap Fetches` - very good;

Stats about index usage:
````sql
SELECT relpages AS 'Disk page read', reltuples AS 'Rows scanned'
FROM pg_class WHERE relname = 'filmTable';
````
