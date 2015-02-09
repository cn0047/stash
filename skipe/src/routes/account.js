var actions = {
    POST: {},
    GET: {},
};

actions.POST.init = function (req, res) {
    res.render('app', {title: 'skipe'});
};

actions.POST.logOut = function (req, res) {
    req.session.destroy();
    res.json({});
};

actions.GET.getUser = function (req, res) {
    res.json(req.session.user);
};

actions.GET.getContacts = function (req, res) {
    global.mongo.collection('contact', function (err, collection) {
        collection
        .find({owner: global.mongo.ObjectID(req.param('user'))})
        .toArray(function (err, docs) {
            res.json(docs);
        });
    });
};

actions.GET.getPosts = function (req, res) {
    global.mongo.collection('post', function (err, collection) {
        collection.find({}, function (err, cursor) {
            cursor.toArray(function (err, docs) {
                var count = docs.length - 1;
                for (i in docs) {
                    (function (docs, i) {
                        global.mongo.dereference(docs[i].chat, function(err, doc) {
                            docs[i].chat = doc;
                            if (i == count) {
                                (function (docs) {
                                    res.json(docs);
                                })(docs);
                            }
                        });
                    })(docs, i)
                }
            });
        });
    });
};

actions.GET.getChats = function (req, res) {
    global.mongo.collection('usersInChat', function (err, collection) {
        collection.find({user: global.mongo.ObjectID(req.param('user'))}, function (err, cursor) {
            cursor.toArray(function (err, docs) {
                var count = docs.length - 1;
                for (i in docs) {
                    (function (docs, i) {
                        global.mongo.dereference(docs[i].chat, function(err, doc) {
                            docs[i].chat = doc;
                            if (i == count) {
                                (function (docs) {
                                    res.json(docs);
                                })(docs);
                            }
                        });
                    })(docs, i)
                }
            });
        });
    });
};

exports.go = function (req, res) {
    /**
     * @todo Delete it.
     */
    req.session.user = global.demoUser;
    if (req.session.user) {
        req.session.user = req.session.user;
        if (req.method in actions && req.param('action') in actions[req.method]) {
            actions[req.method][req.param('action')](req, res);
        }
        if (req.param('action') == undefined) {
            actions.POST.init(req, res);
        }
    } else {
        res.redirect('/#login');
    }
};
