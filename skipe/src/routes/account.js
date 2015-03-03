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

actions.POST.addPost = function (req, res) {
    global.mongo.collection('post', function (err, collection) {
        collection.insert({
            chat: {
              '$ref': 'chat',
              '$id': global.mongo.ObjectID(req.param('chat')),
              '$db': 'skipe'
            },
            user: req.param('user'),
            date: req.param('date'),
            text: req.param('text'),
        }, function (err, docs) {
            if (err) {
                res.json({errors: err});
                return;
            }
            res.json({success: true});
        });
    });
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
        collection.find({'chat.$id': global.mongo.ObjectID(req.param('chat'))}, function (err, cursor) {
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

/**
 * @todo Refactor dereference.
 */
actions.GET.getChats = function (req, res) {
    global.mongo.collection('usersInChat', function (err, collection) {
        collection.find({'user.$id': global.mongo.ObjectID(req.param('user'))}, function (err, cursor) {
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

/**
 * @todo Refactor dereference.
 */
actions.GET.getUsersInChat = function (req, res) {
    global.mongo.collection('usersInChat', function (err, collection) {
        var args = {
            'chat.$id': global.mongo.ObjectID(req.param('chat')),
            'user.$id': {$ne: global.mongo.ObjectID(req.param('user'))}
        };
        collection.find(args, function (err, cursor) {
            cursor.toArray(function (err, docs) {
                var count = docs.length - 1;
                for (i in docs) {
                    (function (docs, i) {
                        global.mongo.dereference(docs[i].user, function(err, doc) {
                            docs[i].user = doc;
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
