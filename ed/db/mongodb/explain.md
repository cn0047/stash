EXPLAIN
-

[docs](https://www.mongodb.com/docs/v4.4/reference/command/explain/)
[docs](https://www.mongodb.com/docs/v4.4/reference/explain-results/)

Verbosity: queryPlanner, executionStats, allPlansExecution.

Stage:
* COLLSCAN - collection scan (no index used).
* IXSCAN - scanning index keys.
* FETCH - retrieving documents.
* SHARD_MERGE - merging results from shards.
* SHARDING_FILTER - filtering out orphan documents from shards.

queryPlanner.winningPlan.memLimit - memory that query planner is allowed to use when optimizing query.

queryPlanner.winningPlan.stage:SORT - blocking sort in memmory (bad).
queryPlanner.winningPlan.stage:PROJECTION_COVERED - no need to read documents, fields values taken from index.
