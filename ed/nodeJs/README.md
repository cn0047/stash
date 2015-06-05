Node JS
-

*v0.10.26*

````js
req.url    // Url string.
req.params // Parsed params from url.
````

#### Mongodb dereference
````js
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
````

####Redis Sessions
````js
var redisSessions = require('redis-sessions');
global.session = new redisSessions();
global.session.create(
    {
        app: 'skipe',
        id: d._id,
        ip: req.connection.remoteAddress,
        ttl: 7200,
        d: {sname: d.sname}
    },
    function(err, res) {
        if (err) {
            console.error(err);
        }
        global.user.sname = d.sname;
        global.user.token = res.token;
    }
);
global.session.set(
    {
        app: 'skipe',
        token: global.user.token,
        d: {sname: global.user.sname}
    },
    function(err, res) {
        if (err) {
            console.error(err);
        }
        console.log(res);
    }
);
````
