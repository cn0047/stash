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

actions.GET.getUser = function (req, res) {
    res.json(req.session.user);
};

actions.GET.getAllContacts = function (req, res) {
    global.mongo.collection('user', function (err, collection) {
        collection.find(
            {},
            {sort: ['sname', 'asc']},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    res.json(docs);
                });
            }
        );
    });
};

actions.GET.getContacts = function (req, res) {
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
