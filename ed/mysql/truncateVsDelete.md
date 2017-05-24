Truncate Vs delete
-

### TRUNCATE

It requires the DROP privilege.
Does not invoke ON DELETE triggers.
It cannot be performed for InnoDB tables with parent-child foreign key relationships.
Truncate operations drop and re-create the table.
Cannot be rolled back.
Any AUTO_INCREMENT value is reset to its start value.

### DELETE

The DELETE statement deletes rows from tbl_name and returns the number of deleted rows.
Need the DELETE privilege on a table to delete rows from it.
You cannot delete from a table and select from the same table in a subquery.
