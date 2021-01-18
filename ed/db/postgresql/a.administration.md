Administration
-

````bash
# init on osx
brew services start postgresql
createdb cws
psql -d cws -L /tmp/p.log
CREATE USER usr WITH PASSWORD 'pass';
GRANT ALL PRIVILEGES ON database dbName TO usr;

psql -d postgres://dbu:dbp@localhost/test
psql -h localhost -U dbu -d test -c 'select 204'

pg_dump -h localhost -p 5432 -U dbu -d test > /var/lib/postgresql/data/dump.sql
psql -h localhost -p 5432 -U dbu -d td < /var/lib/postgresql/data/dump.sql

# remote connection
ssh -i $k -N -L 5431:remoteDBHostName:5432 ubuntu@$h
psql -p 5431 -d postgres://dbu:dbp@localhost/db

pg_ctl status

# general log
# SELECT current_setting('log_statement');
# SET log_statement='all';
# ALTER DATABASE dbname SET log_statement='all';

pg_trgm # Trigram (Trigraph) module
````

#### REPL:

````sql
\pset pager on
\pset pager always
\pset pager off
-- \setenv PAGER less

\set PROMPT2 '    '

# show databases
\l
\d
# use db
\c unittests

# show tables
\dt *.*
\dt *test*
# more info
\d+ viewName

# `SHOW CREATE TABLE`
pg_dump -t tableName --schema-only

# store output into file
\o /tmp/x.sql
\q
````

````sql
VACUUM VERBOSE tableName; -- helpful info

COPY tableName TO 'filePath' CSV (DELIMER ',');
COPY tableName FROM 'filePath' DELIMER ',';
````

#### Configuration:

`/etc/postgresql/8.3/main/postgresql.conf`
`/usr/local/var/postgres/postgresql.conf`

````sh
listen_addresses = '*'
max_connections
shared_buffers # cache, must be 15-25% from OS memory
effective_cache_size # memory for disk cache, must be 50-75% from OS memory

# for intensive writes
checkpoint_segments
wal_buffers
synchronous_commit

work_mem
maintainance_work_mem

log_destination = 'syslog'
redirect_stderr = off
silent_mode = on
syslog_facility = 'LOCAL0'
syslog_ident = 'postgres'

log_min_duration_statement = 0
log_duration = on
log_statement = 'none'
````

Dev:
````sh
# logs
log_statement = 'all'
logging_collector = on
log_min_duration_statement = 0
log_connections = on
# target
log_destination = 'csvlog'
log_directory = '/tmp'
log_filename = 'psql.log'
# verbosity
client_min_messages = notice
log_min_messages = info
log_min_error_statement = info
# output
debug_pretty_print = on
debug_print_parse = off
debug_print_plan = off
debug_print_rewritten = off
````

````sql
SET constraint_exclusion = off;
````
