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

BLOB/TEXT cannot be indexed.

Kinds:

Covering index - a covering index refers to the case
when all fields selected in a query are covered by an index,
in that case InnoDB (not MyISAM) will never read the data in the table,
but only use the data in the index, significantly speeding up the select.
Note that in InnoDB the primary key is included in all secondary indexes,
so in a way all secondary indexes are compound indexes.

Compound index.

Clustered index is synonymous with the primary key.
If you do not define a PRIMARY KEY for your table,
MySQL locates the first UNIQUE index where all the key columns are NOT NULL
and InnoDB uses it as the clustered index.
If the table has no PRIMARY KEY or suitable UNIQUE index,
InnoDB internally generates a hidden clustered index on a synthetic column containing row ID values
(is a 6-byte field that increases monotonically as new rows are inserted).

All indexes other than the clustered index are known as secondary indexes. 

If the primary key is long, the secondary indexes use more space,
so it is advantageous to have a short primary key.

Types:

KEY or INDEX: refers to a normal non-unique index.
Non-distinct values for the index are allowed.

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
