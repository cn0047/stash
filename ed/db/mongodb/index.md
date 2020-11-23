INDEX
-

[reference](http://docs.mongodb.org/manual/reference/indexes/)

Types:

* B-tree
* Hash
* GEO
* Text
* TTL

Limitations:

````
Key value size - 1024 bytes
Key name size - 128 chars
Keys per index (fields in compound index) - 31
Indexes per collection - 64
````

````js
// MongoDB indexes may be ascending, (i.e. 1) or descending (i.e. -1)

// indexes list
db.collection.getIndexes();
db.system.indexes.find();
// Each index requires at least 8KB of data space.
db.inventory.find({ type: 'aston' });
db.inventory.ensureIndex( { type: 1 })
// index uses
db.inventory.find({ type: "food", item:/^c/ }, { item: 1, _id: 0 })
// index not uses, bacause query returns _id field
db.inventory.find({ type: "food", item:/^c/ }, { item: 1 })
// delete index
db.items.dropIndex({name : 1})
// Rebuild Indexes. This operation drops all indexes, including the _id index,
// and then rebuilds all indexes.
db.accounts.reIndex()
/*
{
  topics: ["whaling" , "allegory" , "revenge" ,
    "American" , "novel" , "nautical" , "voyage" , "Cape Cod"]
}
*/
db.volumes.ensureIndex({topics: 1});
db.volumes.findOne({topics: "voyage"}, {title: 1});
// Create a text Index
db.collection.ensureIndex({subject: "text", content: "text"});
db.collection.ensureIndex({"$**": "text"}, {name: "TextIndex"});
// Specify a Language for Text Index
db.quotes.ensureIndex({content: "text"}, {default_language "spanish"});
db.quotes.ensureIndex({quote : "text"}, {language_override: "idioma"} );
// The following example indexes any string value
// in the data of every field of every document in collection and names the index TextIndex:
db.collection.ensureIndex({"$**": "text"}, {name: "TextIndex"});
/*
{ _id: 1, idioma: "portuguese", quote: "A sorte protege os audazes"}
{ _id: 2, idioma: "spanish", quote: "Nada hay más surreal que la realidad."}
{ _id: 3, idioma: "english", quote: "is this a dagger which I see before me"}

// Control Search Results with Weights

{
    _id: 1,
    content: "This morning I had a cup of coffee.",
    about: "beverage",
    keywords: ["coffee"]
}
{
    _id: 2,
    content: "Who doesn't like cake?",
    about: "food",
    keywords: ["cake", "food", "dessert"]
}
*/
db.blog.ensureIndex(
    {
        content: "text",
        keywords: "text",
        about: "text"
    },
    {
        weights: {
            content: 10,
            keywords: 5,
        },
        name: "TextIndex"
    }
);
// Limit the Number of Entries Scanned
/*
{_id: 1, dept: "tech", description: "lime green computer"}
{_id: 2, dept: "tech", description: "wireless red mouse"}
{_id: 3, dept: "kitchen", description: "green placemat"}
{_id: 4, dept: "kitchen", description: "red peeler"}
{_id: 5, dept: "food", description: "green apple"}
{_id: 6, dept: "food", description: "red potato"}
*/
db.inventory.ensureIndex({dept: 1, description: "text"});
// Do not use a hashed index for floating point numbers.
db.active.ensureIndex({a: 'hashed'});
// Unique Indexes
db.members.ensureIndex({'user_id': 1}, {unique: true});
````

````js
// given:
db.foo.createIndex({ a:1, b:1 })
// when
db.foo.insert({a: [1,2], b: [1,2]})
// then
"writeError": "cannot index parallel arrays [b] [a]"
````

(Sparse Indexes)[http://docs.mongodb.org/manual/core/index-sparse/] - when
index key missed in some documents.

````js
// if indexed field das NULL it don not be at result.
{"_id" :ObjectId('523b6e32fb408eea0eec2647'), 'userid': 'newbie'}
{"_id" :ObjectId('523b6e61fb408eea0eec2648'), 'userid': 'abby', 'score': 82}
{"_id" :ObjectId('523b6e6ffb408eea0eec2649'), 'userid': 'nina', 'score': 90}
db.scores.ensureIndex({score: 1}, {sparse: true})
db.scores.find({score: {$lt: 90}})
{'_id' : ObjectId('523b6e61fb408eea0eec2648'), 'userid': 'abby', 'score': 82}
````

````js
// Create index in foreground:
// 1) fast;
// 2) block writes and reads in db (regardless of engine);
// Create in background:
// 1) slow;
// 2) don't block writes and reads;

// NON BLOCK DATA
db.people.ensureIndex({zipcode: 1}, {background: true})
// Drop Duplicates
db.accounts.ensureIndex({username: 1}, {unique: true, dropDups: true})
// Index has the name inventory.
db.products.ensureIndex({item: 1, quantity: -1} , {name: 'inventory'})
// (Specify a Language for Text Index)[http://docs.mongodb.org/manual/tutorial/specify-language-for-text-index/]
db.quotes.ensureIndex({content : "text"}, {default_language: "spanish"})
// Set index weight.
// The default weight is 1.
db.blog.ensureIndex(
    {content: "text", keywords: "text", about: "text"},
    {
        weights: {content: 10, keywords: 5, },
        name: "TextIndex"
    }
);
````

A covered query is a query in which:
* all the fields in the query are part of an index, and
* all the fields returned in the results are in the same index.

````js
// Use Indexes to Sort Query Results
/*
For example, given an index { a: 1, b: 1, c: 1, d: 1 },
if the query condition includes equality match conditions on a and b,
you can specify a sort on the subsets { c: 1 } or { c: 1, d: 1 }
*/
// THE sort() OPERATION WILL ABORT WHEN IT USES 32 MEGABYTES OF MEMORY.
// Collection has the following index: {a: 1, b: 1, c: 1, d: 1}
// Then following operations can use the index efficiently:
db.collection.find({a: 5}).sort({a: 1, b: 1});
db.collection.find({a: 5}).sort({b: 1, c: 1})
db.collection.find({a: 5, c: 4, b: 3}).sort({d: 1})

// given
db.foo.createIndex({ a: 1, b: 1, c: 1 })
// then
db.foo.explain('executionStats').find({ b: 3, c: 4 }) // COLLSCAN
db.foo.explain('executionStats').find({ a: 3 }) // IXSCAN
db.foo.explain('executionStats').find({ c: 1 }).sort({ a: 1, b: 1}) // IXSCAN
db.foo.explain('executionStats').find({ c: 1 }).sort({ a: -1, b: 1}) // COLLSCAN

// given
db.products.createIndex({price: -1})
// then
db.products.find({brand: 'ge'}).sort({price: 1}) // use price index
db.products.find({$and:[{price: {$gt: 30}}, {price: {$lt: 50}}]}).sort({brand: 1}) // use price index

// given
db.products.createIndex({category: 1, brand: 1})
// then
db.products.find({brand: 'ge'}).sort({category: 1, brand: -1}) // won't use index

db.collection.totalIndexSize()
````

[Text Search Languages](http://docs.mongodb.org/manual/reference/text-search-languages/)
