Index
-

### Types:

* B-tree (default)
* Hash
* Generalized Inverted Index (GIN) - for array and full text
* Generalized Search Tree Index (GIST) - geo and  full text

A `partial` index - is an index built over a subset of a table;
the subset is defined by a conditional expression (called the predicate of the partial index).

Allowed max 32 columns in multicolumn index.

`ANALYZE VERBOSE tableName`
