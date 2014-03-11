exports.registration = function (req, res) {
    req.checkBody('email', 'Invalid email').isEmail();
    req.checkBody('sname', 'Invalid sname').matches(/^[\w\s\_\-\=\+@]+$/); // or .is()
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    var mongodb = require('mongodb');
    var mongoServer = new mongodb.Server('localhost', 27017, {auto_reconnect: true});
    var mongoDB = new mongodb.Db('skipe', mongoServer);
    mongoDB.open(function(err, db) {
        if (err) {
            console.log(err);
        } else {
            db.collection('user', function(err, collection) {
                if (err) {
                    console.log(err);
                } else {
                    collection.insert({
                        email: req.param('email'),
                        sname: req.param('sname'),
                    }, function(err, docs) {
                        if (err) {
                            console.log(err);
                        } else {
                            res.json({success: docs});
                        }
                    });
                }
            });
        }
    });
};
