# Neo4j

# -e NEO4J_dbms_connector_https_enabled=false \
# -e NEO4J_dbms_ssl_policy_bolt_enabled=false \
# -e NEO4J_dbms_ssl_policy_https_enabled=false \

NEO4JLABS_PLUGINS='["apoc", "streams", "graphql", "graph-algorithms", "n10s"]'

tag=4.1
tag=4.1.1-enterprise
tag=4.3.3
tag=4.3.7
tag=4.4.0
tag=4.4.0-enterprise
docker run -it --rm --net=xnet -v $PWD:/gh -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j_$tag:/data \
  -v $PWD/.data/.docker/neo4j_gs_plugins:/var/lib/neo4j/plugins \
  -e NEO4J_AUTH=neo4j/test \
  \
  -e NEO4J_ACCEPT_LICENSE_AGREEMENT=yes \
  \
  -e NEO4J_apoc_export_file_enabled=true \
  -e NEO4J_apoc_import_file_enabled=true \
  -e NEO4J_apoc_import_file_use__neo4j__config=true \
  -e NEO4JLABS_PLUGINS='["apoc"]' \
  \
  neo4j:$tag \
  \
  /bin/bash # in case of import/dump or some shell commands

# in case of apoc related error
rm $PWD/.data/.docker/neo4j_gs_plugins/apoc.jar

# 3.5
docker run -it --rm --net=xnet -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j:/data \
  neo4j:3.5.17

# 3.5 with apoc
docker run -it --rm --net=xnet -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j:/data \
  -v $PWD/.data/.docker/neo4j_plugins:/plugins \
  \
  -e NEO4J_apoc_export_file_enabled=true \
  -e NEO4J_apoc_import_file_enabled=true \
  -e NEO4J_apoc_import_file_use__neo4j__config=true \
  -e NEO4JLABS_PLUGINS=\[\"apoc\"\] \
  \
  -e NEO4J_dbms_logs_debug_level=DEBUG \
  -e NEO4J_dbms_logs_http_enabled=true \
  -e NEO4J_dbms_logs_query_allocation__logging__enabled=true \
  -e NEO4J_dbms_logs_query_enabled=true \
  -e NEO4J_dbms_logs_query_page__logging__enabled=true \
  -e NEO4J_dbms_logs_query_parameter_logging_enabled=true \
  -e NEO4J_dbms_logs_query_runtime__logging__enabled=true \
  -e NEO4J_dbms_logs_query_threshold=0 \
  -e NEO4J_dbms_logs_query_time__logging__enabled=true \
  -e NEO4J_dbms_logs_user_stdout_enabled=true \
  -e NEO4J_dbms_track_query_allocation=true \
  -e NEO4J_dbms_track_query_cpu_time=true \
  \
  neo4j:3.5.17

# logs
docker exec -it xneo4j sh -c 'tail -f /logs/debug.log'

# test UI
# usr=neo4j
# pwd=test
# pwd=1
# prev_pwd=neo4j
open http://0.0.0.0:7474/

# test REPL
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test'
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p 1'
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test -d mydb'
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test "SHOW DATABASES"'
#
docker exec -it xneo4j /bin/bash
docker exec -it xneo4j sh -c 'neo4j status'
docker exec -it xneo4j sh -c 'neo4j-admin memrec'
