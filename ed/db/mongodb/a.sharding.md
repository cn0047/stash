Sharding
-

[FAQ: Sharding with MongoDB](http://docs.mongodb.org/manual/faq/sharding/)

`mongos` - router to shards.
For PROD must be 3 config servers.

Ways to shard:
* range based
* hash based (document must contain shard key)

If the shard key is not included in a find operation
and there are 4 shards - mongos has to send the query
to all 4 of the shards.

For PROD you have to have at least 3 config servers.

It's impossible to use unique key with sharding,
because there is no way to check uniqueness between all shards.

AutoSharding - app doesn't need to specify:
* where documents go when writing
* which shard to read from

````js
// in mongos
// db.chunks.find().pretty();

db.test_collection.getShardDistribution();

// update without _id will fail
db.test_collection.update({country: "UA"}, {$set: {id: 1}});
````
