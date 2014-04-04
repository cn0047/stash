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