# PostgreSQL


docker run -it --rm --name xpostgres --hostname xpostgres --net=xnet \
    -v $PWD/.docker/.data/postgresql/xpostgres:/var/lib/postgresql/data \
    -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# check
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@xpostgres/test -c 'select count(*) from test'

# dump
docker exec -ti xpostgres pg_dump -d postgres://dbu:dbp@xpostgres/test -t test --schema-only

# import dump
# docker exec -ti cpqsql /bin/bash -c "psql -d postgres://dbu:dbp@cpqsql/test < /app/dump.sql"

# test
docker exec -ti xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti -e PGPASSWORD=dbp xpostgres psql -h localhost -p 5432 -U dbu -d test
docker exec -ti xpostgres psql -d postgres://dbu:dbp@localhost/test

#### PostgreSQL cluster

docker run -it --rm -p 5432:5432 --name postgres-master --hostname postgres-master \
    -v $PWD/.docker/.data/postgresql:/var/lib/postgresql/data \
    -v $PWD/.docker/postgresql/master.conf:/var/lib/postgresql/data/postgresql.conf \
    -e POSTGRES_DB=test -e POSTGRES_USER=dbu -e POSTGRES_PASSWORD=dbp postgres

# test
docker exec -ti postgres-master psql -h localhost -p 5432 -U dbu -d test
