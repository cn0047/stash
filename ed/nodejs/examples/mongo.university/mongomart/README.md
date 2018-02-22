MongoMart
-

````
docker run -it --rm --net=xnet -p 27017:27017 --hostname xmongo --name xmongo \
    -v $PWD/ed/nodejs/examples/mongo.university/mongomart/:/tmp/d \
    -v $PWD/docker/.data/mongodb:/data/db mongo:latest

docker exec -it xmongo mongoimport --drop -d mongomart -c item /tmp/d/mongomart/data/items.json
docker exec -it xmongo mongoimport --drop -d mongomart -c cart /tmp/d/mongomart/data/cart.json

docker exec -it xmongo mongo mongomart
````