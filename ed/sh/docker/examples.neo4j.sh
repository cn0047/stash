# Neo4j

# -e NEO4J_dbms_connector_https_enabled=false \
# -e NEO4J_dbms_ssl_policy_bolt_enabled=false \
# -e NEO4J_dbms_ssl_policy_https_enabled=false \



# 4.1
docker run -it --rm --net=xnet -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j_41:/data \
  neo4j:4.1

# 4.1 enterprise
docker run -it --rm --net=xnet -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j_41:/data \
  -e NEO4J_AUTH=neo4j/test \
  \
  -e NEO4J_ACCEPT_LICENSE_AGREEMENT=yes \
  \
  -e NEO4J_apoc_export_file_enabled=true \
  -e NEO4J_apoc_import_file_enabled=true \
  -e NEO4J_apoc_import_file_use__neo4j__config=true \
  -e NEO4JLABS_PLUGINS=\[\"apoc\"\] \
  \
  neo4j:4.1.1-enterprise

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



# usr=neo4j
# pwd=test
# pwd=1
# prev_pwd=neo4j
open http://0.0.0.0:7474/

# REPL
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test'
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p 1'
docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test -d mydb'

#
docker exec -it xneo4j /bin/bash
