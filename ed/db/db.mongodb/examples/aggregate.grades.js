/*
Document:
{
    "_id" : ObjectId("50b59cd75bed76f46522c392"),
    "student_id" : 10,
    "class_id" : 5,
    "scores" : [
        {
            "type" : "exam",
            "score" : 69.17634380939022
        },
        {
            "type" : "quiz",
            "score" : 61.20182926719762
        },
        {
            "type" : "homework",
            "score" : 73.3293624199466
        },
        {
            "type" : "homework",
            "score" : 15.206314042622903
        },
        {
            "type" : "homework",
            "score" : 36.75297723087603
        },
        {
            "type" : "homework",
            "score" : 64.42913107330241
        }
    ]
}

Task: calculate the class with the best average student performance.
*/

// AVG for student in class:
db.grades.aggregate([
    {$unwind: "$scores"},
    {$match: {"scores.type": {$ne: "quiz"}}},
    {$group: {_id: {sId: "$student_id", cId: "$class_id"}, as: {$avg: "$scores.score"}}},
    {$sort: {"_id.sId": -1, "_id.cId": -1}},
]).pretty()

// AVG classes:
db.grades.aggregate([
    {$unwind: "$scores"},
    {$match: {"scores.type": {$ne: "quiz"}}},
    {$group: {_id: {sId: "$student_id", cId: "$class_id"}, as: {$avg: "$scores.score"}}},
    {$group: {_id: "$_id.cId", avSc: {$avg: "$as"}}},
    {$sort: {"avSc": -1}},
]).pretty()
