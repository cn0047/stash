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
// 'food' or 'snacks'
db.inventory.find({ type: { $in: [ 'food', 'snacks' ] } })
// food and a less than ($lt) price
db.inventory.find( { type: 'food', price: { $lt: 9.95 } } )
// OR
db.inventory.find({ $or: [{ qty: { $gt: 100 } }, { price: { $lt: 9.95 } } ] })
// 'food' and either the qty has a value greater than ($gt) 100 or price is less than ($lt) 9.95
db.inventory.find( { type: 'food', $or: [ { qty: { $gt: 100 } }, { price: { $lt: 9.95 } } ] } )
// subdocument
db.inventory.find({producer: {company: 'ABC123', address: '123 Street'} })
db.inventory.find( { 'producer.company': 'ABC123' } )
// tags is an array
db.inventory.find( { tags: [ 'fruit', 'food', 'citrus' ] } )
// value of the tags field is an array whose first element equals 'fruit'
db.inventory.find( { 'tags.0' : 'fruit' } )
db.inventory.find( { 'memos.0.by': 'shipping' } )
// field memo equal to 'on time' and the field by equal to 'shipping'
db.inventory.find({'memos.memo': 'on time', 'memos.by': 'shipping'} )
// memos is an array that has memo equal to 'on time' and the field by equal to 'shipping'
db.inventory.find( {memos: {$elemMatch: {memo : 'on time', by: 'shipping'} } } )

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
db.inventory.find({ type: 'aston' });
db.inventory.ensureIndex( { type: 1 })
// index uses
db.inventory.find({ type: "food", item:/^c/ }, { item: 1, _id: 0 })
// index not uses, bacause query returns _id field
db.inventory.find({ type: "food", item:/^c/ }, { item: 1 })

// EXPLAIN
db.collection.find().explain()

// INSERT
// MongoDB always adds the _id field
db.users.insert({name: "sue", age: 26, status: "A"})
db.inventory.save({type: "book", item: "notebook", qty: 40})
// UPDATE
// by default mongo update single row. use {multi: true} to update all maches documents
db.users.update({ age: { $gt: 18 } }, { $set: { status: "A" } }, { multi: true })
// REPLACE
db.inventory.update(
    { type: "book", item : "journal" },
    { $set : { qty: 10 } },
    { upsert : true }
)
// DELETE
db.users.remove({ status: "D" })

// To set errors ignored write concern, specify w values of -1 to your driver.
// To set unacknowledged write concern, specify w values of 0 to your driver.
// To set acknowledged write concern, specify w values of 1 to your driver. DEFAULT.
// To set a journaled write concern, specify w values of 1 and set the journal or j option to true for your driver.
// To set replica acknowledged write concern, specify w values greater than 1 to your driver.

````
