Index
-

If Cardinality to low - index will not uses!

Max key length is 767 bytes.

If you use latin1 then the largest column you can index is varchar(767),
<br>but if you use utf8 then the limit is varchar(255).
<br>There is also a separate 3072 byte limit per index.
<br>The 767 byte limit is per column.
<br>So you can include multiple columns (each 767 bytes or smaller) up to 3072 total bytes per index.
<br>Using **innodb_large_prefix** allows you to include columns up to 3072 bytes long in InnoDB indexes.

Types:

KEY or INDEX: refers to a normal non-unique index.
Non-distinct values for the index are allowed

UNIQUE: refers to an index where all rows of the index must be unique.

PRIMARY: like a UNIQUE index but may be only one on a table,
should not be used on any columns which allow NULL values, can be auto_increment.

FULLTEXT: are different from all of the above, it's only used for a "full text search" feature.

####Detection of bad index

````sql
SELECT
    t.TABLE_NAME,
    SUM(t.ROWS_READ) AS raw_readed,
    SUM(i.ROWS_READ) AS key_readed,
    ROUND((SUM(i.ROWS_READ)/SUM(t.ROWS_READ))*100, 2) AS index_coverage
FROM information_schema.TABLE_STATISTICS t
LEFT join information_schema.INDEX_STATISTICS i ON t.TABLE_SCHEMA = i.TABLE_SCHEMA AND t.TABLE_NAME = i.TABLE_NAME
WHERE t.TABLE_SCHEMA = 'dbName'
GROUP BY t.TABLE_NAME
HAVING raw_readed > 10000
ORDER BY raw_readed DESC
;
````
