Aggregation
-

Additionally, map-reduce operations can have output sets that exceed the 16 megabyte output limitation of the aggregation pipeline.
The aggregation pipeline can use indexes to improve its performance during some of its stages.
In addition, the aggregation pipeline has an internal optimization phase.
Map-reduce operations can also output to a sharded collection

####Aggregation with the Zip Code Data Set
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

####Aggregation with User Preference Data.
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
// The $unwind operator separates each value in the likes array, and creates a new version of the source document for every element in the array.
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

####Map-Reduce Examples
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

####Perform Incremental Map-Reduce
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

[Aggregation Commands Comparison](http://docs.mongodb.org/manual/reference/aggregation-commands-comparison/)
[SQL to Aggregation Mapping Chart](http://docs.mongodb.org/manual/reference/sql-aggregation-comparison/)
[Aggregation Interfaces](http://docs.mongodb.org/manual/reference/operator/aggregation/interface/)
[Variables in Aggregation](http://docs.mongodb.org/manual/reference/aggregation-variables/)