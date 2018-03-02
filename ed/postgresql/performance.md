Performance
-

## Performance optimization

Use `BEGIN-COMMIT` for batch insert (to disable autocommit).

`SOME` is a synonym for `ANY`, `IN` is equivalent to `= ANY`.
`IN` may be faster than `ANY` but it depends...

## Stats

````
select * from pg_stat_all_tables where relname like 'MOCK_DATA%';

seq_scan            - number of sequential scans initiated
seq_tup_read        - number of live rows fetched by sequential scans
idx_scan            - number of index scans initiated (over all indexes belonging to the table)
idx_tup_fetch       - number of live rows fetched by index scans
n_tup_ins           - numbers of row insertions
n_tup_upd           - updates
n_tup_del           - and deletions
n_tup_hot_upd       - number of row updates that were HOT (i.e. no separate index update)
n_live_tup          - numbers of live rows
n_dead_tup          - numbers of dead rows
n_mod_since_analyze - _
last_vacuum         - the last time the table was non-FULL vacuumed manually
last_autovacuum     - the last time it was vacuumed by the autovacuum daemon
last_analyze        - the last time it was analyzed manually
last_autoanalyze    - the last time it was analyzed by the autovacuum daemon
vacuum_count        - number of times it has been non-FULL vacuumed manually
autovacuum_count    - number of times it has been vacuumed by the autovacuum daemon
analyze_count       - number of times it has been analyzed manually
autoanalyze_count   - and the number of times it has been analyzed by the autovacuum daemon

  relid  | schemaname |          relname           | seq_scan | seq_tup_read | idx_scan | idx_tup_fetch | n_tup_ins | n_tup_upd | n_tup_del | n_tup_hot_upd | n_live_tup | n_dead_tup | n_mod_since_analyze | last_vacuum | last_autovacuum | last_analyze | last_autoanalyze | vacuum_count | autovacuum_count | analyze_count | autoanalyze_count
---------+------------+----------------------------+----------+--------------+----------+---------------+-----------+-----------+-----------+---------------+------------+------------+---------------------+-------------+-----------------+--------------+------------------+--------------+------------------+---------------+-------------------
 2394534 | public     | device_file                |       16 |         2240 |          |               |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
 2394811 | public     | device_file_sandbox        |       26 |        15808 |        0 |             0 |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
 2395273 | public     | device_file_sandbox_status |        3 |          276 |          |               |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
 2397930 | public     | device_file_rating         |       38 |       154736 |        0 |             0 |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
 2401336 | public     | device_file_rating_file    |        2 |           88 |          |               |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
 2401714 | public     | device_file_sandbox_file   |        2 |           16 |          |               |         0 |         0 |         0 |             0 |          0 |          0 |                   0 |             |                 |              |                  |            0 |                0 |             0 |                 0
````

````
select * from pg_stat_all_indexes where relname like 'MOCK_DATA%';

idx_scan - number of index scans initiated on that index
idx_tup_read - number of index entries returned by index scans
idx_tup_fetch - and number of live table rows fetched by simple index scans using that index

 relid  | indexrelid | schemaname |                relname                |                      indexrelname                      | idx_scan | idx_tup_read | idx_tup_fetch
--------+------------+------------+---------------------------------------+--------------------------------------------------------+----------+--------------+---------------
  16797 |      17507 | public     | device_file                           | device_file_pkey                                       |      240 |          240 |           240
  16797 |      17655 | public     | device_file                           | sha1_uk                                                |        0 |            0 |             0
  16805 |      19201 | public     | device_file_rating                    | file_rating_device                                     |        0 |            0 |             0
  19215 |      19219 | public     | device_file_sandbox_status            | device_file_sandbox_status_pkey                        |        0 |            0 |             0
  19261 |      19268 | public     | device_file_rating_file               | device_file_rating_file_pkey                           |        0 |            0 |             0
  16812 |      17509 | public     | device_file_sandbox                   | device_file_sandbox_pkey                               |       69 |           69 |            69
  16805 |      17523 | public     | device_file_rating                    | device_raiting_file_pkey                               |        0 |            0 |             0
  16812 |      17698 | public     | device_file_sandbox                   | file_sandbox_id_file_index                             |        8 |            8 |             8
  16805 |      17697 | public     | device_file_rating                    | file_rating_id_file_index                              |        0 |            0 |             0
  19440 |      19447 | public     | device_file_sandbox_file              | device_file_sandbox_file_pkey                          |        0 |            0 |             0
  16805 |      19104 | public     | device_file_rating                    | i_device_file_rating__actual_verdict__id_file__id      |        0 |            0 |             0
````

````
-- Effectiveness of the buffer cache.
-- When the number of actual disk reads is much smaller than the number of buffer hits,
-- then the cache is satisfying most read requests without invoking a kernel call.

select * from pg_statio_all_tables where relname like 'MOCK_DATA%';

heap_blks_read  - number of disk blocks read from that table
heap_blks_hit   - number of buffer hits
idx_blks_read   - numbers of disk blocks read
idx_blks_hit    - and buffer hits in all indexes of that table
toast_blks_read - numbers of disk blocks read
toast_blks_hit  - and buffer hits from that table's auxiliary TOAST table (if any)
tidx_blks_read  - and numbers of disk blocks read
tidx_blks_hit   - and buffer hits for the TOAST table's index.

  relid  | schemaname |          relname           | heap_blks_read | heap_blks_hit | idx_blks_read | idx_blks_hit | toast_blks_read | toast_blks_hit | tidx_blks_read | tidx_blks_hit
---------+------------+----------------------------+----------------+---------------+---------------+--------------+-----------------+----------------+----------------+---------------
 2394534 | public     | device_file                |              3 |            45 |               |              |               0 |              0 |              0 |             0
 2394811 | public     | device_file_sandbox        |             14 |           350 |             1 |           22 |               0 |              0 |              0 |             0
 2395273 | public     | device_file_sandbox_status |              2 |             1 |               |              |                 |                |                |
 2397930 | public     | device_file_rating         |             71 |          2627 |             3 |            6 |               0 |              0 |              0 |             0
 2401336 | public     | device_file_rating_file    |              4 |             0 |               |              |              85 |             51 |              3 |            67
 2401714 | public     | device_file_sandbox_file   |              2 |             0 |               |              |              18 |              6 |              3 |            11
````

````
select * from pg_statio_all_indexes where relname like 'MOCK_DATA%';

idx_blks_read -  numbers of disk blocks read
idx_blks_hit -  and buffer hits in that index

  relid  | indexrelid | schemaname |       relname       |                   indexrelname                    | idx_blks_read | idx_blks_hit
---------+------------+------------+---------------------+---------------------------------------------------+---------------+--------------
 2397930 |    6814659 | public     | device_file_rating  | file_rating_device                                |             1 |            2
 2397930 |    6814660 | public     | device_file_rating  | file_rating_id_file_index                         |             1 |            2
 2394811 |    6814661 | public     | device_file_sandbox | file_sandbox_id_file_index                        |             1 |           22
 2397930 |    6814671 | public     | device_file_rating  | i_device_file_rating__actual_verdict__id_file__id |             1 |            2
````
