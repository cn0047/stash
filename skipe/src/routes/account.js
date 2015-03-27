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

/**
 * Returns chat id.
 *
 * Activates present chat (if it exists directly with target user)
 * or creates new.
 *
 * @todo Security.
 */
actions.POST.startChat = function (req, res) {
    var user = global.mongo.ObjectID(req.param('user'));
    var withUser = global.mongo.ObjectID(req.param('withUser'));
    var caption = req.param('userSname')
        + res.__('with')
        + req.param('withUserSname');
    var startNewChat = function (res) {
        // New chat.
        global.mongo.collection('chat', function (err, collection) {
            collection.insert({caption: caption}, function (err, documents) {
                if (err) {
                    res.json({errors: err});
                    return;
                }
                var chat = documents[0]._id;
                // Add target user to chat.
                global.mongo.collection('usersInChat', function (err, collection) {
                    var d = {
                        chat: {_id: chat, caption: caption},
                        user: {_id: withUser, sname: req.param('withUserSname')}
                    };
                    collection.insert(d, function (err, docs) {
                        if (err) {
                            res.json({errors: err});
                            return;
                        }
                        // Add me to chat.
                        global.mongo.collection('usersInChat', function (err, collection) {
                            var d = {
                                chat: {_id: chat, caption: caption},
                                user: {_id: user, sname: req.param('userSname')}
                            };
                            collection.insert(d, function (err, docs) {
                                if (err) {
                                    res.json({errors: err});
                                    return;
                                }
                                // All OK.
                                res.json(chat);
                            });
                        });
                    });
                });
            });
        });
    }
    global.mongo.collection('usersInChat', function (err, collection) {
        // Find my chats.
        collection.find(
            {'user._id': user},
            {_id: false, 'chat._id': true},
            function (err, cursor) {
                cursor.toArray(function (err, docs) {
                    var chats = docs.map(function (o) {
                        return global.mongo.ObjectID(o.chat._id);
                    });
                    // Find users with whom i want start chat in users for my present chats.
                    // Find only chats where i'm and user with whom i want start chat.
                    global.mongo.collection('usersInChat', function (err, collection) {
                        collection.aggregate(
                            {
                                $group: {
                                    _id: {chat: '$chat._id'},
                                    user: {$push: '$user._id'},
                                    count: {$sum: 1}
                                }
                            },
                            {
                                $match: {
                                    '_id.chat': {$in: chats},
                                    count: {$eq: 2},
                                    user: user,
                                    user: withUser
                                }
                            },
                            function (err, documents) {
                                if (documents.length === 0) {
                                    startNewChat(res);
                                } else {
                                    // Activate present chat.
                                    res.json(documents[0]._id.chat);
                                }
                            }
                        );
                    });
                });
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
