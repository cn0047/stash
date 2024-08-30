# PostgreSQL

tag=9.6
tag=12.19
tag=latest
docker run -it --rm -p 5432:5432 \
  -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres:$tag

docker run -it --rm --name xpostgres --hostname xpostgres --net=xnet \
  -v $PWD/.data/.docker/postgresql_$tag/xpostgres:/var/lib/postgresql/data \
  -v /tmp/dump:/tmp/dump \
  -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres:$tag

# check
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test -c 'select count(*) from test'

# dump export
docker exec -ti xpostgres pg_dump -d postgres://dbu:dbp@xpostgres/test -t test --schema-only
# dump import
docker exec -ti xpostgres /bin/bash
psql -d postgres://dbu:dbp@xpostgres/test < /tmp/dump/dump.sql

# import dump
# docker exec -ti cpqsql /bin/bash -c "psql -d postgres://dbu:dbp@cpqsql/test < /app/dump.sql"

# test
docker exec -ti xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti -e PGPASSWORD=dbp xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@localhost/test



#### PostgreSQL cluster

docker run -it --rm -p 5432:5432 --name postgres-master --hostname postgres-master \
  -v $PWD/.data/.docker/postgresql:/var/lib/postgresql/data \
  -v $PWD/.docker/postgresql/master.conf:/var/lib/postgresql/data/postgresql.conf \
  -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# test
docker exec -ti postgres-master psql -h localhost -p 5432 -U dbu -d test
