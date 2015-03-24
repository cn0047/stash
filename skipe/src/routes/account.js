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

/**
 * @todo Security.
 */
actions.POST.addPost = function (req, res) {
    global.mongo.collection('post', function (err, collection) {
        collection.insert(
            {
                chat: global.mongo.ObjectID(req.param('chat')),
                user: req.param('user'),
                date: req.param('date'),
                text: req.param('text'),
            },
            function (err, docs) {
                if (err) {
                    res.json({errors: err});
                    return;
                }
                res.json({success: true});
            }
        );
    });
};

/**
 * @todo Security.
 */
actions.POST.addContact = function (req, res) {
    global.mongo.collection('contact', function (err, collection) {
        collection.insert(
            {
                owner: global.mongo.ObjectID(req.param('owner')),
                user: {_id: req.param('userId'), sname: req.param('sname')},
            },
            function (err, docs) {
                if (err) {
                    res.json({errors: err});
                    return;
                }
                res.json({success: true});
            }
        );
    });
};

actions.GET.getUser = function (req, res) {
    res.json(req.session.user);
};

/**
 * @todo Security.
 */
actions.GET.getAllContacts = function (req, res) {
    global.mongo.collection('contact', function (err, collection) {
        collection.find(
            {owner: global.mongo.ObjectID(req.param('user'))},
            {'user._id': true, _id: false},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    var myContacts = docs.map(function (o) {
                        return global.mongo.ObjectID(o.user._id);
                    });
                    // Exclude myself.
                    myContacts.push(global.mongo.ObjectID(req.param('user')));
                    // All available contacts without myself and my contacts.
                    global.mongo.collection('user', function (err, collection) {
                        collection.find(
                            {_id: {$nin: myContacts}},
                            {_id: true, sname: true},
                            {sort: ['sname', 'asc']},
                            function (err, cursor) {
                                cursor.toArray(function (err, docs) {
                                    res.json(docs);
                                });
                            }
                        );
                    });
                });
            }
        );
    });
};

/**
 * @todo Security.
 */
actions.GET.getMyContacts = function (req, res) {
    global.mongo.collection('contact', function (err, collection) {
        collection.find(
            {owner: global.mongo.ObjectID(req.param('user'))},
            {sort: ['user.sname', 'asc']},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    res.json(docs);
                });
            }
        );
    });
};

/**
 * @todo Security.
 */
actions.GET.getPosts = function (req, res) {
    global.mongo.collection('post', function (err, collection) {
        collection.find(
            {'chat': global.mongo.ObjectID(req.param('chat'))},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    res.json(docs);
                });
            }
        );
    });
};

actions.GET.getChats = function (req, res) {
    req.checkParams('user', res.__('invalidUser')).isMongoId();
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    global.mongo.collection('usersInChat', function (err, collection) {
        collection.find(
            {'user._id': global.mongo.ObjectID(req.param('user'))},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    res.json(docs);
                });
            }
        );
    });
};

actions.GET.getUsersInChat = function (req, res) {
    req.checkParams('chat', res.__('chatUser')).isMongoId();
    req.checkParams('user', res.__('invalidUser')).isMongoId();
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    global.mongo.collection('usersInChat', function (err, collection) {
        var args = {
            'chat._id': global.mongo.ObjectID(req.param('chat')),
            'user._id': {$ne: global.mongo.ObjectID(req.param('user'))}
        };
        collection.find(args, function (err, cursor) {
            cursor.toArray(function (err, docs) {
                res.json(docs);
            });
        });
    });
};

actions.GET.getContactInfo = function (req, res) {
    req.checkParams('user', res.__('invalidUser')).isMongoId();
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    global.mongo.collection('user', function (err, collection) {
        collection.findOne(
            {'_id': global.mongo.ObjectID(req.param('user'))},
            {sname: true, email: true},
            function (err, doc) {
                res.json(doc);
            }
        );
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
