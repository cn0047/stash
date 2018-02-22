Operators
-

#### [Operators](http://docs.mongodb.org/manual/reference/operator/query/)
````js
// OR
// all the clauses in the $or expression must be supported by indexes.
// Otherwise, MongoDB will perform a collection scan.
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
// TEXT (text search)
db.articles.find( { $text: { $search: "coffee" } } )
// contain the words bake or coffee but do not contain the term cake:
db.articles.find( { $text: { $search: "bake coffee -cake" } } )
// WHERE
db.myCollection.find( { $where: "this.credits == this.debits" } );
db.myCollection.find( { $where: function() { return obj.credits == obj.debits; } } );
// ALL - Equivalent to $and Operation, array contains all specified elements.
{ tags: { $all: [ "ssl" , "security" ] } }
// ELEMMATCH
db.scores.find({ results: { $elemMatch: { $gte: 80, $lt: 85 } } })
// SIZE - array has certain size
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
//  first matching document
db.products.update( { sku: "abc123" }, { $inc: { quantity: 5 } } );
// update all matching documents
db.records.update( { age: 20 }, { $inc: { age: 1 } }, { multi: true } );
// 1st matching document
db.products.update( { sku: "abc123" }, { $inc: { quantity: -2, sales: 2 } } );
// $mul
db.products.update({ _id: 1 }, { $mul: { price: 1.25 } })
// $rename
// renames the field nickname to alias, and the field cell to mobile
db.students.update( { _id: 1 }, { $rename: { 'nickname': 'alias', 'cell': 'mobile' } } )
// $setOnInsert
// during replace when insert
db.products.update({ _id: 1 }, { $setOnInsert: { defaultQty: 100 } }, { upsert: true })
// $set
db.products.update({ sku: "abc123" }, { $set: { quantity: 500, instock: true, "details.make": "ZYX" } })
// $unset
db.products.update( { sku: "unknown" }, { $unset: {quantity: "", instock: ""} } )
// $min
// if stored document has lowScore lower than 150 -  nothing happens, otherwise lowScore will set to 150
db.scores.update( { _id: 1 }, { $min: { lowScore: 150 } } )
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
db.inventory.aggregate([{$project: {
    item: 1, qty: 1, result: {$and: [{$gt: ["$qty", 100 ]}, {$lt: ["$qty", 250]}]}
}}])
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
