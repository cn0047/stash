Aggregation
-

[Accumulators](https://docs.mongodb.com/manual/meta/aggregation-quick-reference/#accumulators)
[Aggregation Commands Comparison](http://docs.mongodb.org/manual/reference/aggregation-commands-comparison/)
[SQL to Aggregation Mapping Chart](http://docs.mongodb.org/manual/reference/sql-aggregation-comparison/)
[Aggregation Interfaces](http://docs.mongodb.org/manual/reference/operator/aggregation/interface/)
[Variables in Aggregation](http://docs.mongodb.org/manual/reference/aggregation-variables/)

The aggregation pipeline can use indexes to improve its performance during some of its stages.
In addition, the aggregation pipeline has an internal optimization phase.

Result is limited by size = 16MB.

For performance (follow this flow):
1. $match as early as possible
2. $project as early as possible
3. $sort as early as possible
4. use index

#### Pipeline Aggregation Stages

* $collStats - statistics regarding a collection.
* $indexStats - returns statistics regarding the use of each index for the collection.
* $count - returns a count of the number of documents at this stage of the aggregation pipeline.
* $project - like select in SQL
* $match - filters the document stream.
* $unwind - from document with array to documents with only one element.
* $sort
* $sortByCount - groups incoming documents based on the value of a specified expression.
* $group
* $limit
* $skip
* $out - must be the last stage in the pipeline.
* $redact - can be used to implement field level redaction.
* $sample
* $geoNear
* $lookup
* $facet - processes multiple aggregation pipelines.
* $bucket - categorizes incoming documents into groups.
* $bucketAuto
* $addFields - adds new fields to documents.
* $replaceRoot - replaces a document with the specified embedded document.
* $graphLookup - performs a recursive search on a collection.

`$group: {"_id": "all", ...}`

````js
// Add field1 to field2:
// given
db.aggtest.insert({a: 1, b: 1});
db.aggtest.insert({a: 2, b: 2});
// when
db.aggtest.aggregate({ $project: { aPlusB: { $sum: [ "$a", "$b" ] } } });
// then result will be
// { "_id" : ObjectId("5a7be85d460df66d8243b967"), "aPlusB" : 2 }
// { "_id" : ObjectId("5a7be861460df66d8243b968"), "aPlusB" : 4 }

// count 😀
db.device_file.aggregate([
    {$group: {_id: null, 'count': {$sum: 1}}}
])
````

#### [Group Accumulator Operators](http://docs.mongodb.org/manual/reference/operator/aggregation-group/)
````js
$addToSet - put grouped documents into one (only distinct)
$avg
$first - `{"$group": {_id: "$smt", sample: {"$first": "$item"}}}`
$last
$max
$min
$push - put grouped documents into one
$sum
````

#### Aggregation with the Zip Code Data Set
````js
// Return States with Populations above 10 Million.
/*
{
    "_id": "10280",
    "city": "NEW YORK",
    "state": "NY",
    "pop": 5574,
    "loc": [
        -74.016323,
        40.710537
    ]
}
*/
db.zipcodes.aggregate(
    {
        $group: {_id: "$state", totalPop: {$sum : "$pop"}}
    },
    {
        $match: {totalPop: {$gte :10*1000*1000}}
    }
);
/*
SELECT state, SUM(pop) AS totalPop
FROM zipcodes
GROUP BY state
HAVING totalPop >= (10*1000*1000)
*/

db.zipcodes.aggregate(
    {
        $group: {_id: {state: "$state", city: "$city"}, pop: {$sum : "$pop"}}
    },
    {
        $group: {_id: "$_id.state", avgCityPop: {$avg : "$pop"}}
    }
);

// Return Largest and Smallest Cities by State.
db.zipcodes.aggregate(
    {
        $group: {_id: {state: "$state", city: "$city" }, pop: {$sum: "$pop"}}
    },
    {$sort: {pop: 1}},
    {
        $group: {
            _id:          "$_id.state",
            biggestCity:  {$last: "$_id.city"},
            biggestPop:   {$last: "$pop"},
            smallestCity: {$first: "$_id.city"},
            smallestPop:  {$first: "$pop"}
        }
    },
    // the following $project is optional, and
    // modifies the output format.
    {
        $project: {
            _id:          0,
            state:        "$_id",
            biggestCity:  {name: "$biggestCity",  pop: "$biggestPop"},
            smallestCity: {name: "$smallestCity", pop: "$smallestPop"}
        }
    }
);
````

#### Aggregation with User Preference Data.
````js
/*
{
    _id : "jane",
    joined : ISODate("2011-03-02"),
    likes : ["golf", "racquetball"]
}
{
    _id : "joe",
    joined : ISODate("2012-07-02"),
    likes : ["tennis", "golf", "swimming"]
}
*/
// Normalize and Sort Documents.
db.users.aggregate([
    {$project: {name: {$toUpper: "$_id"}, _id: 0}},
    {$sort: {name: 1}}
]);
// Return Usernames Ordered by Join Month.
db.users.aggregate([
    {
        $project: {
            month_joined: {$month: "$joined"},
            name: "$_id",
            _id: 0
        }
    },
    {$sort: {month_joined: 1}}
]);
// Return Total Number of Joins per Month.
db.users.aggregate([
    {$project: {month_joined: {$month: "$joined"}}} ,
    {$group: {_id: {month_joined: "$month_joined"}, number: {$sum: 1}}},
    {$sort: {"_id.month_joined": 1}}
]);
// Return the Five Most Common “Likes”.
db.users.aggregate([
    {$unwind: "$likes"},
    {$group: {_id: "$likes", number: {$sum: 1}}},
    {$sort: {number: -1}}, // sort in reverse order.
    {$limit: 5}
]);
// The $unwind operator separates each value in the likes array,
// and creates a new version of the source document for every element in the array.
// Example:
{
    _id: "jane",
    joined: ISODate("2011-03-02"),
    likes: ["golf", "racquetball"]
}
// The $unwind operator would create the following documents:
{
    _id: "jane",
    joined: ISODate("2011-03-02"),
    likes: "golf"
}
{
    _id: "jane",
    joined: ISODate("2011-03-02"),
    likes: "racquetball"
}
````

#### Map-Reduce Examples

Map-reduce operations can have output sets that exceed the 16 megabyte output limitation
of the aggregation pipeline.
Map-reduce operations can also output to a sharded collection.

Reduce operations can run in parallel across shards.

````js
var mapFunction1 = function () {
    emit(this.cust_id, this.price);
};
var reduceFunction1 = function (keyCustId, valuesPrices) {
    return Array.sum(valuesPrices);
};
db.ordersCollection.mapReduce(
    mapFunction1,
    reduceFunction1,
    {out: "map_reduce_example"}
);
````

Example 2.

````js
db.sites.insert({url: "www.google.com", date: new Date(), duration_seconds: 5 , scope: ["search", "open-source"]});
db.sites.insert({url: "www.no-fucking-idea.com", date: new Date(), duration_seconds: 13 , scope: ["tech"]});
db.sites.insert({url: "www.google.com", date: new Date(), duration_seconds: 1 , scope: ["search", "tech"]});
db.sites.insert({url: "www.no-fucking-idea.com", date: new Date(), duration_seconds: 69 , scope: ["tech"]});
db.sites.insert({url: "www.no-fucking-idea.com", date: new Date(), duration_seconds: 256 , scope: ["tech"]});
db.sites.insert({url: "www.google.com", date: new Date(), duration_seconds: 1 , scope: ["search", "music"]});
db.sites.insert({url: "www.youtube.com", date: new Date(), duration_seconds: 1256 , scope: ["music", "video"]});
db.sites.insert({url: "stackoverflow.com", date: new Date(), duration_seconds: 256 , scope: ["tech", "open-source"]});
db.sites.insert({url: "www.github.com", date: new Date(), duration_seconds: 256 , scope: ["tech", "open-source"]});
db.sites.insert({url: "stackoverflow.com", date: new Date(), duration_seconds: 256 , scope: ["tech", "open-source", "music"]});
// db.sites.insert({url: "www.youtube.com", date: new Date(), duration_seconds: 196 , scope: ["sport"]});
````

````js
var map = function(){
  emit(this.url, 1);
}
var reduce = function(key, values){
  var res = 0;
  values.forEach(function(v){ res += 1});
  return {count: res};
}
db.sites.mapReduce(map, reduce, { out: "mapped_urls" });
db.mapped_urls.find()
/*
Result:
{ "_id" : "stackoverflow.com", "value" : { "count" : 2 } }
{ "_id" : "www.github.com", "value" : 1 }
{ "_id" : "www.google.com", "value" : { "count" : 3 } }
{ "_id" : "www.no-fucking-idea.com", "value" : { "count" : 3 } }
{ "_id" : "www.youtube.com", "value" : 1 }
*/

var map = function(){
  for (var i in this.scope) {
    emit(this.scope[i], 1);
  }
}
var reduce = function(key, values){
  var res = 0;
  values.forEach(function(v){res += 1});
  return {count: res};
}
db.sites.mapReduce(map, reduce, { out: "mapped_urls_2" });
db.mapped_urls_2.find()
/*
Result:
{ "_id" : "music", "value" : { "count" : 3 } }
{ "_id" : "open-source", "value" : { "count" : 4 } }
{ "_id" : "search", "value" : { "count" : 3 } }
{ "_id" : "tech", "value" : { "count" : 7 } }
{ "_id" : "video", "value" : 1 }
*/
````

#### Perform Incremental Map-Reduce
````js
var mapFunction = function () {
    var key = this.userid;
    var value = {
        userid: this.userid,
        total_time: this.length,
        count: 1,
        avg_time: 0
    };
    emit( key, value );
};
var reduceFunction = function (key, values) {
    var reducedObject = {
        userid: key,
        total_time: 0,
        count:0,
        avg_time:0
    };
    values.forEach(function (value) {
        reducedObject.total_time += value.total_time;
        reducedObject.count += value.count;
    });
    return reducedObject;
};
var finalizeFunction = function (key, reducedValue) {
    if (reducedValue.count > 0) {
        reducedValue.avg_time = reducedValue.total_time / reducedValue.count;
    }
    return reducedValue;
};
db.sessions.mapReduce(
    mapFunction,
    reduceFunction,
    {
        out: {reduce: "session_stat"},
        finalize: finalizeFunction
    }
);
db.sessions.mapReduce(
    mapFunction,
    reduceFunction,
    {
        query: {ts: {$gt: ISODate('2011-11-05 00:00:00')}},
        out: {reduce: "session_stat"},
        finalize: finalizeFunction
    }
);
````
