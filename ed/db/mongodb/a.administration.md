Administration
-

`ext4` - the best file system for mongo.
`/etc/fstab` - add `noatime` (last access time) to this file = os won't update atime.

`--oplogSize` - must be 5% of disck space and <= 50GB.

````sh
sudo service mongodb start|stop|restart

sudo rm /var/lib/mongodb/mongod.lock
sudo service mongodb restart

# osx
brew services start mongodb

mongo
mongo "mongodb://$h:$p/$db" -u $user -p $password
mongo | tee log.txt
````

````js
// set custom prompt
prompt = () => {
    return `main@primary> `;
}
````

````js
db.version();

DBQuery.shellBatchSize = 10;

// to fix NotPrimaryNoSecondaryOk
rs.secondaryOk()

db.getLastError()
db.getLastErrorObj()

db.adminCommand('listDatabases'); // list DBs
db.getSiblingDB('<db>');
db.getCollectionNames(); // list collections
db.getUsers();
db.getRoles({showBuiltinRoles: true});

db.currentOp()
// kill <mongod process ID>

db.serverStatus()
db.isMaster()
// replica set’s status
rs.status()
// sharding status
sh.status()

// get queries
db.currentOp();
// kill query
db.killOp($opID);

db.stats()
// slow queries log
db.setProfilingLevel(1, 3) // 3 millisecond
// tool with stats
mongostat
// shows time spend
mongotop

// measure working set
db.runCommand({serverStatus: 1, workingSet: 1})

db.collection.stats()
db.collection.dataSize() // collection's size in bytes
db.collection.storageSize()
db.collection.totalSize()
db.collection.getIndexes()
db.collection.totalIndexSize()
db.collection.drop() // delete collection

// Data Type Fidelity
// data_binary
{"$binary" : "<bindata>", "$type" : "<t>"}
// data_date
Date(<date>)
// data_timestamp
Timestamp(<t>, <i>)
// data_regex
/<jRegex>/<jOptions>
// data_oid
ObjectId("<id>")
// data_ref
DBRef("<name>", "<id>")
// type
db.inventory.find({price: {$type: 1}});
db.inventory.find({$where: "Array.isArray(this.tags)"})
[Available types values](http://docs.mongodb.org/manual/reference/operator/query/type/)
````

````sh
# Collection Export
mongoexport --collection collection --out collection.json
mongoexport --db sales --collection contacts --query '{"field": 1}'
# Collection Import
mongoimport --collection collection --file collection.json
````

````js
// Creates a new collection explicitly.
// maximum size of 5 megabytes and a maximum of 5000 documents.
db.createCollection("log", { capped : true, size : 5242880, max : 5000 })

// Check if a Collection is Capped
// Capped collections are fixed-size collections. Work in a way similar to circular buffers.
db.collection.isCapped()
db.cappedCollection.find().sort( { $natural: -1 } )
// Convert a Collection to Capped
db.runCommand({"convertToCapped": "mycoll", size: 100000});

// Expire Documents after a Certain Number of Seconds
db.log.events.ensureIndex( { "createdAt": 1 }, { expireAfterSeconds: 3600 } )
// Expire Documents at a Certain Clock Time
db.app.events.ensureIndex( { "expireAt": 1 }, { expireAfterSeconds: 0 } )

db.runCommand({isMaster: 1})
db.runCommand({buildInfo: 1})
db._adminCommand({buildInfo: 1})
// Shut down the mongod from the mongo shell
db.shutdownServer()
db.shutdownServer({timeoutSecs : 5})
// Force Replica Set Shutdown
db.adminCommand({shutdown : 1, force : true})
db.adminCommand({shutdown : 1, timeoutSecs : 5})

// Store a JavaScript Function on the Server
db.system.js.save({
    _id: "echoFunction",
    value : function(x) { return x; }
})
db.eval("echoFunction( 'test' )")

// Query Authenticated Users
db.system.users.find()
// Create a User Administrator
db.addUser({
    user: "<username>",
    pwd: "<password>",
    roles: [ "userAdminAnyDatabase" ]
})
// Add a User to a Database
db.addUser({
    user: "Alice",
    pwd: "Moon1234",
    roles: [ "readWrite", "dbAdmin" ],
    otherDBRoles: { config: [ "readWrite" ]
})
````

#### Monitoring

Logging:

````js
var level = 5; // verbosity level from 0 to 5
db.setLogLevel(level, 'query'); // 2nd param is optional
````

Profilling:

````js
// To enable profiling:
mongod --profile 1 --slowms 2

// Profiling levels:
// 1 - nothing,
// 2 - slow queries,
// 3 - log all.

// profiling info
db.system.profile.find().pretty();

var level = 2; // 0, 1 or 2
db.setProfilingLevel(level)
````

#### Dump

````sh
mongodump --host mongodb1.example.net --port 3017 --username user --password pass --out /opt/backup/mongodump-2013-10-24
mongorestore --host mongodb1.example.net --port 3017 --username user --password pass /opt/backup/mongodump-2013-10-24/
````
