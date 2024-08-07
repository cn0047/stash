Index
-

````sql
CREATE INDEX "IDX_image_created" ON image USING btree (created);

SELECT * FROM pg_indexes WHERE tablename = 'tbl';
````

When a row is added, it does not go to the main index directly.
But instead it is added to a TODO list, which is then processed by `VACUUM`.
When you scan the index, you have to scan the tree AND
sequentially read what is still in the pending list.
If the pending list is long, this will have some impact on performance.

Index hint - PostgreSQL don't support index hint.

### Types:

* B-tree (default).
* Hash (can't be unique key).
* Generalized Inverted Index (GIN) - for array and full text.
* Generalized Search Tree Index (GIST) - geo ~and full text~.
* Space-Partitioned GIST (SP-GIST).
* Block Range Index (BRIN).

A `partial` index - is an index built over a subset of a table;
the subset is defined by a conditional expression (called the predicate of the partial index).

Allowed max 32 columns in multicolumn index.

````sql
ANALYZE VERBOSE tableName
````
