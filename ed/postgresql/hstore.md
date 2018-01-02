HSTORE
-

HSTORE - data type for storing sets of key/value pairs
within a single PostgreSQL value. 

````
psql -h localhost -d test -c 'create extension hstore;'
````

````
# will return x
'a=>x, b=>y'::hstore -> 'a'

#will return {z,x}
'a=>x, b=>y, c=>z'::hstore -> ARRAY['c','a']

# will return true
'a=>1'::hstore ? 'a'
'a=>1,b=>2'::hstore ?& ARRAY['a','b']
'a=>1,b=>2'::hstore ?| ARRAY['b','c']
'a=>b, b=>1, c=>NULL'::hstore @> 'b=>1'

# will return false
'a=>1,b=>2'::hstore ?| ARRAY['c'];
````

````
# ADD a key, or UPDATE an existing key with a new value:
UPDATE tab SET h = h || ('c' => '3');

# DELETE key:
'a=>1, b=>2, c=>3'::hstore - 'b'::text

# DELETE keys:
'a=>1, b=>2, c=>3'::hstore - ARRAY['a','b']
'a=>1, b=>2, c=>3'::hstore - 'a=>4, b=>2'::hstore
# or
UPDATE tab SET h = delete(h, 'k1');
````

HSTORE has GiST and GIN index support.

````
CREATE INDEX hidx ON testhstore USING GIST (h);
CREATE INDEX hidx ON testhstore USING GIN (h);
````

HSTORE also supports btree or hash indexes for the = operator.
This allows HSTORE columns to be declared UNIQUE, or to be used in GROUP BY, ORDER BY or DISTINCT expressions. 

````
CREATE INDEX hidx ON testhstore USING BTREE (h);
CREATE INDEX hidx ON testhstore USING HASH (h);
````
