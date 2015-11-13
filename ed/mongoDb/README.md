mongo
-

*MongoDB shell version: 2.4.6*

#### Introduction
````
sudo service mongodb start|stop|restart

sudo rm /var/lib/mongodb/mongod.lock
sudo service mongodb restart

mongo
````

MongoDB implements locks on a per-database basis for most read and write operations.
<br>[Which operations lock the database](http://docs.mongodb.org/manual/faq/concurrency/#which-operations-lock-the-database)

````js
// current database
db
show dbs
use mydb

// insert a collection
j = {name : "mongo"}
db.testData.insert(j)
db.test.insert({x :MaxKey})
db.test.insert({x :MinKey})

show collections
db.testData.find()
db.getCollection('_foo').find()
db.testData.findOne()
db.testData.find().limit(3)
db.testData.find({ name : "mongo" })
// iterate over the cursor with a loop
var c = db.testData.find()
while (c.hasNext()) printjson(c.next())
// print 4th item in list
printjson(c[4])
// print(tojson(c[4]));
var myCursor =  db.inventory.find( { type: 'food' } );
myCursor.forEach(printjson);
// cursor to array
var myCursor = db.inventory.find( { type: 'food' } );
var documentArray = myCursor.toArray();
var myDocument = documentArray[3];
// or
var myCursor = db.inventory.find( { type: 'food' } );
var myDocument = myCursor[3];
// or
myCursor.toArray() [3];
// search all documents with age exactly 18
db.users.find({age: 18})
// greater than condition
db.users.find({age:{$gt: 18}}).sort({age; 1})
db.users.find({age:{$gte: 18}}).sort({age; 1})
db.users.find({age:{$ne: 18}}) // not equal
//
db.records.find( { "user_id": { $lte: 42} }, { "name": 1, "email": 1} )
// switch on/off fields
db.users.find({age: {$gt: 18}}, {name: 1, address: 1, _id: 0}).limit(5)
// 'food' or 'snacks'
db.inventory.find({ type: { $in: [ 'food', 'snacks' ] } })
db.inventory.find({ type: { $nin: [ 'food', 'snacks' ] } })
// food and a less than ($lt) price
db.inventory.find( { type: 'food', price: { $lt: 9.95 } } )

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

// select count(user_id) from users
db.users.count( { user_id: { $exists: true } } )
db.users.distinct( "status" )

// copy all objects from one collection to another
db.collection.copyTo(newCollection)

// EXPLAIN
db.collection.find().explain()
// hint - Forces MongoDB to use a specific index for a query.
db.inventory.find( { type: 'food' } ).hint( { type: 1 } ).explain()
db.inventory.find( { type: 'food' } ).hint( { type: 1, name: 1 } ).explain()

// BSON-document size limit = 16MB
// In MongoDB, operations are atomic at the document level. No single write operation can change more than one document.

// INSERT
// MongoDB always adds the _id field
db.users.insert({name: "sue", age: 26, status: "A"})
// save() - replaces an existing document with the same _id
db.inventory.save({type: "book", item: "notebook", qty: 40})
db.inventory.save({_id: 10, type: "misc", item: "placard"} )

// UPDATE
// update() - modify existing data or modify a group of documents
// by default mongo update single row. use {multi: true} to update all maches documents
db.users.update({ age: { $gt: 18 } }, { $set: { status: "A" } }, { multi: true })
db.inventory.update({ type : "book" }, { $inc : { qty : -1 } }, { multi: true } )
db.students.update(
    { _id: 1 },
    { $push: { scores: {
        $each : [{ attempt: 3, score: 7 },
        { attempt: 4, score: 4 } ],
        $sort: { score: 1 },
        $slice: -3
    } } }
)
db.users.update(
    { },
    { $unset: { join_date: "" } },
    { multi: true }
)
book = {
    _id: 123456789,
    title: "MongoDB: The Definitive Guide",
    author: [ "Kristina Chodorow", "Mike Dirolf" ],
    published_date: ISODate("2010-09-24"),
    pages: 216,
    language: "English",
    publisher_id: "oreilly",
    available: 3,
    checkout: [ { by: "joe", date: ISODate("2012-10-15") } ]
}
db.books.findAndModify ({
    query: {
        _id: 123456789,
        available: { $gt: 0 }
    },
    update: {
        $inc: { available: -1 },
        $push: { checkout: { by: "abc", date: new Date() } }
    }
})
// REPLACE
db.inventory.update(
    { type: "book", item : "journal" },
    { $set : { qty: 10 } },
    { upsert : true }
)

// DELETE
db.users.remove({ status: "D" })
// delete documents but don't delete indexes
// to remove data and indexes use method drop()

// remove 1 document
db.inventory.remove( { type : "food" }, 1 )

db.users.drop()

// To set errors ignored write concern, specify w values of -1 to your driver.
// To set unacknowledged write concern, specify w values of 0 to your driver.
// To set acknowledged write concern, specify w values of 1 to your driver. DEFAULT.
// To set a journaled write concern, specify w values of 1 and set the journal or j option to true for your driver.
// To set replica acknowledged write concern, specify w values greater than 1 to your driver.

db.runCommand( { getLastError: 1, j: "true" } )
````

#### [Operators](http://docs.mongodb.org/manual/reference/operator/query/)
````js
// OR
// all the clauses in the $or expression must be supported by indexes. Otherwise, MongoDB will perform a collection scan.
db.inventory.find({ $or: [{ qty: { $gt: 100 } }, { price: { $lt: 9.95 } } ] })
// 'food' and either the qty has a value greater than ($gt) 100 or price is less than ($lt) 9.95
db.inventory.find( { type: 'food', $or: [ { qty: { $gt: 100 } }, { price: { $lt: 9.95 } } ] } )
// AND
db.inventory.find( { $and: [ { price: { $ne: 1.99 } }, { price: { $exists: true } } ] } )
// NOT
db.inventory.find( { price: { $not: { $gt: 1.99 } } } )
// NOR
// selects the documents that fail all the query expressions
db.inventory.find( { $nor: [ { price: 1.99 }, { sale: true } ] } )
// MOD
db.inventory.find( { qty: { $mod: [ 4, 0 ] } } )
// REGEX
db.collection.find( { field: /acme.*corp/i } );
db.collection.find( { field: { $regex: 'acme.*corp', $options: 'i' } } );
// TEXT
db.articles.find( { $text: { $search: "coffee" } } )
// contain the words bake or coffee but do not contain the term cake:
db.articles.find( { $text: { $search: "bake coffee -cake" } } )
// WHERE
db.myCollection.find( { $where: "this.credits == this.debits" } );
db.myCollection.find( { $where: function() { return obj.credits == obj.debits; } } );
// ALL - Equivalent to $and Operation.
{ tags: { $all: [ "ssl" , "security" ] } }
// ELEMMATCH
db.scores.find({ results: { $elemMatch: { $gte: 80, $lt: 85 } } })
// SIZE
db.collection.find( { field: { $size: 2 } } );
// $ (projection)
db.students.find( { semester: 1, grades: { $gte: 85 } }, { "grades.$": 1 } )
// $elemMatch
db.schools.find( { zipcode: "63109" }, { students: { $elemMatch: { school: 102, age: { $gt: 10} } } } )
// $meta
db.collection.find( {}, { score: { $meta: "textScore" } } )
// $slice (projection)
db.posts.find( {}, { comments: { $slice: 5 } } )
db.posts.find( {}, { comments: { $slice: [ 20, 10 ] } } ) // [ skip , limit ]
// $inc
db.products.update( { sku: "abc123" }, { $inc: { quantity: 5 } } ); //  first matching document
db.records.update( { age: 20 }, { $inc: { age: 1 } }, { multi: true } ); // update all matching documents
db.products.update( { sku: "abc123" }, { $inc: { quantity: -2, sales: 2 } } ); // 1st matching document
// $mul
db.products.update({ _id: 1 }, { $mul: { price: 1.25 } })
// $rename
db.students.update( { _id: 1 }, { $rename: { 'nickname': 'alias', 'cell': 'mobile' } } ) // renames the field nickname to alias, and the field cell to mobile
// $setOnInsert
db.products.update({ _id: 1 }, { $setOnInsert: { defaultQty: 100 } }, { upsert: true }) // during replace when insert
// $set
db.products.update({ sku: "abc123" }, { $set: { quantity: 500, instock: true, "details.make": "ZYX" } })
// $unset
db.products.update( { sku: "unknown" }, { $unset: {quantity: "", instock: ""} } )
// $min
db.scores.update( { _id: 1 }, { $min: { lowScore: 150 } } )
// if stored document has lowScore lower than 150 -  nothing happens, otherwise lowScore will set to 150
// $max
db.scores.update( { _id: 1 }, { $max: { highScore: 950 } } )
// $currentDate
db.users.update(
    { _id: 1 },
    {
        $currentDate: {lastModified: true, lastModifiedTS: { $type: "timestamp" }},
        $set: { status: "D" }
    }
)
db.users.update({}, {$currentDate: {created: true, updated: true } })
// $
db.students.update( { _id: 1, grades: 80 }, { $set: { "grades.$" : 82 } } )
// $ operator acts as a placeholder for the first match
db.students.update( { _id: 4, "grades.grade": 85 }, { $set: { "grades.$.std" : 6 } } )
// $addToSet
db.inventory.update({ _id: 1 }, { $addToSet: { tags: "accessories"  } })
// $pop
db.students.update( { _id: 1 }, { $pop: { scores: -1 } } // -1 first, 1 last
// $pullAll
db.survey.update( { _id: 1 }, { $pullAll: { scores: [ 0, 5 ] } } )
// $pull
db.cpuinfo.update({ flags: "msr" }, { $pull: { flags: "msr" } }, { multi: true })
db.profiles.update( { _id: 1 }, { $pull: { votes: { $gte: 6 } } } )
// $pushAll
db.collection.update( { field: value }, { $pushAll: { field1: [ value1, value2, value3 ] } } );
// $push
db.students.update({ _id: 1 }, { $push: { scores: 89 } })
db.students.update({ name: "joe" }, { $push: { scores: { $each: [ 90, 92, 85 ] } } })
// $each
db.students.update({ name: "joe" }, { $push: { scores: { $each: [ 90, 92, 85 ] } } })
// $slice
db.students.update( { _id: 1 }, { $push: { scores: {$each: [ 80, 78, 86 ], $slice: -5 } } } )
// $sort
db.students.update({ _id: 2 }, { $push: { tests: { $each: [ 40, 60 ], $sort: 1 } } })
// $position
db.students.update( { _id: 1 }, { $push: { scores: {$each: [ 20, 30 ], $position: 2 } } } )
// $bit
db.switches.update({ _id: 1 }, { $bit: { expdata: { and: NumberInt(10) } } } )
db.switches.update({ _id: 2 }, { $bit: { expdata: { or: NumberInt(5) } } } )
db.switches.update({ _id: 3 }, { $bit: { expdata: { xor: NumberInt(5) } } } )
// $isolated
db.foo.update({ status : "A" , $isolated : 1 }, { $inc : { count : 1 } }, { multi: true } )
````

#### Aggregation Pipeline Operators
````js
// $geoNear
// ...
// $out
db.books.aggregate([{ $group : { _id : "$author", books: { $push: "$title" } } }, { $out : "authors" }])
// $and
db.inventory.aggregate([{$project: {item: 1, qty: 1, result: {$and: [{$gt: ["$qty", 100 ]}, {$lt: ["$qty", 250]}]}}}])
````

#### Set Operators (Aggregation)
````js
// $allElementsTrue
{ $allElementsTrue: [ [ true, 1, "someString" ] ] } // true
// $anyElementTrue
{ $anyElementTrue: [ [ true, false ] ] } // true
{ $anyElementTrue: [ [ ] ] } // false
// $setDifference
{ $setDifference: [ [ "a", "b", "a" ], [ "b", "a" ] ] } // [ ]
// $setEquals
{ $setEquals: [ [ "a", "b", "a" ], [ "b", "a" ] ] } // true
// $setIntersection
{ $setIntersection: [ [ "a", "b", "a" ], [ "b", "a" ] ] } // [ "b", "a" ]
// $setIsSubset
{ $setIsSubset: [ [ "a", "b", "a" ], [ "b", "a" ] ] } // true
// $setUnion
{ $setUnion: [ [ "a", "b", "a" ], [ "b", "a" ] ] } // [ "b", "a" ]
````

####[Comparison Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-comparison/)
````js
$cmp
$eq
$gt
$gte
$lt
$lte
$ne
````

####[Arithmetic Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-arithmetic/)
````js
$add
$divide
$mod
$multiply
$subtract
````

####[String Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-string/)
````js
$concat
$strcasecmp
$substr
$toLower
$toUpper
````

####[Text Search Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-text-search/)
````js
$meta
````

####[Array Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-array/)
````js
$size
````

####[Aggregation Variable Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-projection/)
````js
$let
$map
````

####[Aggregation Literal Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-literal/)
````js
$literal // does not evaluate the expression
````

####[Date Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-date/)
````js
$dayOfMonth
$dayOfWeek
$dayOfYear
$hour
$millisecond // between 0 and 999
$minute
$month
$second
$week
$year
````

####[Conditional Aggregation Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-conditional/)
````js
$cond // $cond: { if: { $gte: [ "$qty", 250 ] }, then: 30, else: 20 }
$ifNull
````

####[Group Accumulator Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-group/)
````js
$addToSet
$avg
$first
$last
$max
$min
$push
$sum
````

####[Query Modifiers](http://docs.mongodb.org/manual/reference/operator/query-modifier/)
````js
$comment
$explain
$hint
$maxScan
$maxTimeMS
$max
$min
$orderby
$query
$returnKey
$showDiskLoc
$snapshot
$natural
````

####[Database Commands](http://docs.mongodb.org/manual/reference/command/)!!!

####[mongo Shell Methods](http://docs.mongodb.org/manual/reference/method/)!!!

[Geospatial Query Operators](http://docs.mongodb.org/manual/reference/operator/query-geospatial/#geospatial-query-operators)

[MongoDB CRUD Reference](http://docs.mongodb.org/manual/reference/crud/#mongodb-crud-reference)

[SQL to MongoDB Mapping Chart](http://docs.mongodb.org/manual/reference/sql-comparison/#sql-to-mongodb-mapping-chart)

In general, use embedded data models when: you have one-to-one or one-to-many model
((link)[http://docs.mongodb.org/manual/core/data-model-design/]).
<br>For model many-to-many use relationships with document references.

````js
// Model Tree Structures with Parent References
db.categories.insert( { _id: "MongoDB", parent: "Databases" } )
db.categories.insert( { _id: "dbm", parent: "Databases" } )
db.categories.insert( { _id: "Databases", parent: "Programming" } )
db.categories.insert( { _id: "Languages", parent: "Programming" } )
db.categories.insert( { _id: "Programming", parent: "Books" } )
db.categories.insert( { _id: "Books", parent: null } )
// parent
db.categories.findOne( { _id: "MongoDB" } ).parent
// index
db.categories.ensureIndex( { parent: 1 } )

// Model Tree Structures with Child References
db.categories.insert( { _id: "MongoDB", children: [] } )
db.categories.insert( { _id: "dbm", children: [] } )
db.categories.insert( { _id: "Databases", children: [ "MongoDB", "dbm" ] } )
db.categories.insert( { _id: "Languages", children: [] } )
db.categories.insert( { _id: "Programming", children: [ "Databases", "Languages" ] } )
db.categories.insert( { _id: "Books", children: [ "Programming" ] } )

// Manual References
original_id = ObjectId()
db.places.insert({
    "_id": original_id,
    "name": "Broadway Center",
    "url": "bc.example.net"
})
db.people.insert({
    "name": "Erin",
    "places_id": original_id,
    "url":  "bc.example.net/Erin"
})
// DBRef
{ "$ref" : <value(collection)>, "$id" : <value(_id)>, "$db" : <value> }
{
    "_id" : ObjectId("5126bbf64aed4daf9e2ab771"),
    // .. application fields
    "creator" : {
        "$ref" : "creators",
        "$id" : ObjectId("5126bc054aed4daf9e2ab772"),
        "$db" : "users"
    }
}
oId = ObjectId();
db.address.insert({'_id' : oId, country: 'UK', city: 'London'});
db.user.insert({name : 'Bond', address: {'$ref': 'address', '$id': oId, '$db': 'user'}});

ObjectId("507f191e810c19729de860ea").getTimestamp();
ObjectId("507f191e810c19729de860ea").str;
ObjectId("507f191e810c19729de860ea").valueOf();
ObjectId("507f191e810c19729de860ea").toString();

var mydate1 = new Date();
var mydate2 = ISODate();
mydate1.toString();
````

[Development patterns](http://docs.mongodb.org/manual/tutorial/#development-patterns)

````js
// Limit Number of Elements in an Array after an Update.
/*
{
    _id: 1,
    scores: [
        {attempt: 1, score: 10},
        {attempt: 2, score:8}
    ]
}
*/
db.students.update(
    {_id: 1},
    {
        $push: {scores: {
            $each : [
                {attempt: 3, score: 7},
                {attempt: 4, score: 4}
            ],
            $sort: {score: 1},
            $slice: -3
        }}
    }
);
/*
Result:
{
    "_id" : 1,
    "scores" : [
        {"attempt" : 3, "score" : 7},
        {"attempt" : 2, "score" : 8},
        {"attempt" : 1, "score" : 10}
    ]
}
*/
````

#### Dump
````js
mongodump --host mongodb1.example.net --port 3017 --username user --password pass --out /opt/backup/mongodump-2013-10-24
mongorestore --host mongodb1.example.net --port 3017 --username user --password pass /opt/backup/mongodump-2013-10-24/
````
