View
-

View - virtual table that does not store data.
When query view - DB executes the underlying query each time to retrieve data.
Do not provide any performance benefits,
performance depends on efficiency of underlying query.

Materialized View - physical copy of query result at the time it was created or refreshed.
Data in MV not automatically updated when the underlying data changes.
Can significantly improve query performance
for complex or resource-intensive queries, because MV store precomputed results.
