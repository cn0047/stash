Administration
-

````js
db.version();

DBQuery.shellBatchSize = 10;

db.getLastError()
db.getLastErrorObj()

db.adminCommand('listDatabases');
db.getSiblingDB('<db>');
db.getCollectionNames();
db.getUsers();
db.getRoles({showBuiltinRoles: true});

db.currentOp()
// kill <mongod process ID>
db.serverStatus()
db.stats()
db.isMaster()
// replica set’s status
rs.status()
// sharding status
sh.status()
db.killOp()

// measure working set
db.runCommand({serverStatus: 1, workingSet: 1})

db.collection.stats()
db.collection.dataSize()
db.collection.storageSize()
db.collection.totalSize()
db.collection.totalIndexSize()

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


// Collection Export
mongoexport --collection collection --out collection.json
mongoexport --db sales --collection contacts --query '{"field": 1}'
// Collection Import
mongoimport --collection collection --file collection.json

// Creates a new collection explicitly.
// maximum size of 5 megabytes and a maximum of 5000 documents.
db.createCollection("log", { capped : true, size : 5242880, max : 5000 })

// Check if a Collection is Capped
db.collection.isCapped()
db.cappedCollection.find().sort( { $natural: -1 } )
// Convert a Collection to Capped
db.runCommand({"convertToCapped": "mycoll", size: 100000});

// Expire Documents after a Certain Number of Seconds
db.log.events.ensureIndex( { "createdAt": 1 }, { expireAfterSeconds: 3600 } )
// Expire Documents at a Certain Clock Time
db.app.events.ensureIndex( { "expireAt": 1 }, { expireAfterSeconds: 0 } )

db.runCommand({ isMaster: 1 })
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
