# ElasticSearch

#### ES cluster v7

img=docker.elastic.co/elasticsearch/elasticsearch:7.5.1
img=elasticsearch:7.1.1
docker run -it --rm -p 9200:9200 -p 9300:9300 --name es \
  -v $PWD/.data/.docker/elasticsearch:/usr/share/elasticsearch/data \
  --memory=1024m \
  -e http.host=0.0.0.0 \
  -e http.publish_host=127.0.0.1 \
  -e network.publish_host=127.0.0.1 \
  -e transport.host=127.0.0.1 \
  -e ES_JAVA_OPTS="-Xms512m -Xmx512m" \
  -e discovery.type=single-node \
  -e discovery.zen.minimum_master_nodes=1 \
  -e http.port=9200 \
  -e xpack.security.enabled=false \
  -e xpack.monitoring.enabled=false \
  -e xpack.graph.enabled=false \
  -e xpack.watcher.enabled=false \
  $img

#### ES cluster v5

# init master 1 node
docker run -it --rm -p 9200:9200 --name es-master-1 \
  -e "bootstrap.memory_lock=true" -e "ES_JAVA_OPTS=-Xms256m -Xmx256m" \
  -e "http.host=_eth0_" -e "cluster.name=ec" \
  -e "node.master=true" -e "node.data=false" elasticsearch:5.4

# init data 1 node - ⚠️ not finished
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1  \
  -e "bootstrap.memory_lock=true" -e "ES_JAVA_OPTS=-Xms256m -Xmx256m" \
  -e "http.host=_eth0_" -e "cluster.name=ec" \
  -e "node.master=false" -e "node.data=true" \
  -e "discovery.zen.ping.unicast.hosts=es-master-1" elasticsearch:5.4

#### ES cluster v2

docker run -it --rm -p 9200:9200 --name es elasticsearch:2.2

# init master 1 node
docker run -it --rm -p 9200:9200 --name es-master-1 elasticsearch:2.2 \
  elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=true -Des.node.data=false

# init data 1 node
docker run -it --rm -p 9201:9200 --name es-data-1 --link es-master-1 elasticsearch:2.2 \
  elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=false -Des.node.data=true \
  -Des.discovery.zen.ping.unicast.hosts=es-master-1

# init data 2 node
docker run -it --rm -p 9202:9200 --name es-data-2 --link es-master-1 elasticsearch:2.2 \
  elasticsearch -Des.network.host=_eth0_ -Des.cluster.name=ec -Des.node.master=false -Des.node.data=true \
  -Des.discovery.zen.ping.unicast.hosts=es-master-1
