# Neo4j


docker run -it --rm --net=xnet -p 7474:7474 -p 7687:7687 --name xneo4j --hostname xneo4j \
  -v $PWD/.data/.docker/neo4j:/data \
  neo4j:3.5

  # -e NEO4J_dbms_connector_https_enabled=false \
  # -e NEO4J_dbms_ssl_policy_https_enabled=false \
  # -e NEO4J_dbms_ssl_policy_bolt_enabled=false \

# usr=neo4j
# pwd=test # prev_pwd=neo4j
open http://0.0.0.0:7474/

docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test'
