# Neo4j


docker run -it --rm -p 7474:7474 -p 7687:7687 --name xneo4j \
  -v $PWD/.data/.docker/neo4j:/data \
  neo4j

# usr=neo4j
# pwd=test
open http://0.0.0.0:7474/

docker exec -it xneo4j sh -c 'cypher-shell -u neo4j -p test'
