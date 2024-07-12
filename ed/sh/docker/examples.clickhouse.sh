# ClickHouse

tag=latest
tag=24.5.3.5
docker run -it --rm --net=xnet -p 9000:9000 -v $PWD:/gh --name xclickhouse clickhouse/clickhouse-server:$tag

# test
docker exec -it xclickhouse clickhouse-client --host localhost --port 9000
