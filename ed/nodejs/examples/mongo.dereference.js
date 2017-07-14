collection.find({}, function (err, cursor) {
    cursor.toArray(function (err, docs) {
        var count = docs.length - 1;
        for (i in docs) {
            (function (docs, i) {
                db.dereference(docs[i].ref, function(err, doc) {
                    docs[i].ref = doc;
                    if (i == count) {
                        (function (docs) {
                            console.log(docs);
                        })(docs);
                    }
                });
            })(docs, i)
        }
    });
});
