// Aggregation $graphLookup 2
/*
Tree:
lexus  -> category  ->  suv
audi   -> model     ->  6
                    ->  7
                    ->  8    ->  A8
bmw    -> series    ->  x
*/

db.vendor.insertMany([
  {_id: 1, name: "lexus", root: 1},
  {_id: 2, name: "audi", root: 2},
  {_id: 3, name: "bmw", root: 3},
]);
db.vendor.drop();

db.cars.insertMany([
   {_id: 1, name: "category", children: [10]},
   {_id: 2, name: "model", children: [26,27,28]},
   {_id: 3, name: "series", children: [30]},
   {_id: 10, name: "suv", children: []},
   {_id: 26, name: "6", children: []},
   {_id: 27, name: "7", children: []},
   {_id: 28, name: "8", children: [281]},
   {_id: 30, name: "x", children: []},
   {_id: 281, name: "A8", children: []},
]);
db.cars.drop();

db.vendor.aggregate([
  {$match: {name: "audi"}},
  {
    $graphLookup: {
      from: "cars",
      startWith: "$root",
      connectFromField: "children",
      connectToField: "_id",
      maxDepth: 99,
      depthField: "height",
      as: "tree"
    }
  }
]).pretty();

// Result:
{
  "_id" : 2,
  "name" : "audi",
  "root" : 2,
  "tree" : [
    {
      "_id" : 2,
      "name" : "model",
      "children" : [
        26,
        27,
        28
      ],
      "height" : NumberLong(0)
    },
    {
      "_id" : 26,
      "name" : "6",
      "children" : [ ],
      "height" : NumberLong(1)
    },
    {
      "_id" : 27,
      "name" : "7",
      "children" : [ ],
      "height" : NumberLong(1)
    },
    {
      "_id" : 28,
      "name" : "8",
      "children" : [
        281
      ],
      "height" : NumberLong(1)
    },
    {
      "_id" : 281,
      "name" : "A8",
      "children" : [ ],
      "height" : NumberLong(2)
    },
  ]
}
