mongo
-

sudo service mongodb start|stop|restart

mongo
````js
// current database
db
show dbs
use mydb

// insert a collection
j = {name : "mongo"}
db.testData.insert(j)
show collections
db.testData.find()
db.testData.findOne()
db.testData.find().limit(3)
db.testData.find({ name : "mongo" })
// iterate over the cursor with a loop
var c = db.testData.find()
while (c.hasNext()) printjson(c.next())
// print 4th item in list
printjson(c[4])
// search all documents with age exactly 18
db.users.find({age: 18})
// greater than condition
db.users.find({age:{$gt: 18}}).sort({age; 1})
//
db.records.find( { "user_id": { $lt: 42} }, { "name": 1, "email": 1} )
// switch on/off fields
db.users.find({age: {$gt: 18}}, {name: 1, address: 1, _id: 0}).limit(5)

// server will close CURSOR after 10 minutes, to avoid it do:
var myCursor = db.inventory.find().addOption(DBQuery.Option.noTimeout);
// data returns in batches
var myCursor = db.inventory.find();
var myFirstDocument = myCursor.hasNext() ? myCursor.next() : null;
// shows how many documents remain in batch
myCursor.objsLeftInBatch();
// info about all cursors
db.runCommand( { cursorInfo: 1 } )

// INDEX
db.inventory.find( { type: 'aston' } );
db.inventory.ensureIndex( { type: 1 } )
// index uses
db.inventory.find( { type: "food", item:/^c/ }, { item: 1, _id: 0 } )
// index not uses, bacause query returns _id field
db.inventory.find( { type: "food", item:/^c/ }, { item: 1 } )

// EXPLAIN
db.collection.find().explain()

// INSERT
// MongoDB always adds the _id field
db.users.insert({name: "sue", age: 26, status: "A"})
// UPDATE
// by default mongo update single row. use {multi: true} to update all maches documents
db.users.update({ age: { $gt: 18 } }, { $set: { status: "A" } }, { multi: true })
// DELETE
db.users.remove({ status: "D" })

````
