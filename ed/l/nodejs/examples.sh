# node.js

docker run -it --rm node:latest node -v

# from Dockerfile
docker build -t xnodejs ./.docker/nodejs
docker run -it --rm -p 8000:3000 xnodejs
# to test
curl 0.0.0.0:8000

# based on Alpine Linux
docker run -it --rm node:alpine node -v
docker run -it --rm -v $PWD:/gh -w /gh node:latest node /gh/x.js

# simple mysql test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mysql --link mysql-master node:latest node index.js

# simple elasticsearch test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/elasticsearch --link es node:latest node index.js

# simple mongo test
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest npm install
docker run -it --rm -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo --net=xnet node:latest node index.js
# simple mongo test with bridge
docker network create --driver bridge x_node_mongo
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node index.js
#
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo node:latest node mongo.universityhw3-3.js
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 node npm i
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh/ed/nodejs/examples/mongo.university/hw3-4 \
    node:latest node overviewOrTags.js
