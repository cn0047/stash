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
