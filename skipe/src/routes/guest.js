var actions = {
    GET: {},
    POST: {},
};

actions.GET.getDefaultLogIn = function (req, res) {
    res.json(global.demoUser);
}

actions.POST.registration = function (req, res) {
    req.checkBody('email', res.__('invalidEmail')).isEmail();
    req.checkBody('sname', res.__('invalidScreenName')).matches(global.validator.pattern.sname);
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    global.mongo.collection('user', function (err, collection) {
        if (err) {
            res.json({errors: err});
        } else {
            collection.findOne({email: req.param('email')}, {_id: 1}, function (err, doc) {
                if (err) {
                    res.json({errors: err});
                    return;
                }
                if (doc) {
                    res.json({errors: [{param: 'email', msg: res.__('emailAlreadyExists')}]});
                    return;
                }
                collection.insert({
                    email    : req.param('email'),
                    sname    : req.param('sname'),
                    password : require('crypto').randomBytes(5).toString('hex'),
                }, function (err, docs) {
                    if (err) {
                        res.json({errors: err});
                        return;
                    }
                    var args = {
                        email    : docs[0].email,
                        sname    : docs[0].sname,
                        password : docs[0].password,
                        locale   : res.getLocale(),
                    };
                    res.render('guest/mailRegistration', args, function (err, html) {
                        if (err) {
                            res.json({errors: err});
                            return;
                        }
                        global.mail.sendMail({
                            to      : docs[0].email,
                            subject : res.__('guestMailSubjectRegistration', docs[0].email),
                            html    : html,
                        }, function (err, r) {
                            if (err) {
                                res.json({errors: err});
                            } else {
                                res.json({success: true});
                            }
                        });
                    });
                });
            });
        }
    });
};

function checkUserArguments(req, res, cb) {
    req.checkBody('type', res.__('invalidLoginType')).matches(/^(email|sname)$/);
    var e = req.validationErrors();
    switch (req.param('type')) {
        case 'email':
            req.checkBody('token', res.__('invalidEmail')).isEmail();
            var find = {email: req.param('token')};
            break;
        case 'sname':
            req.checkBody('token', res.__('invalidScreenName')).matches(global.validator.pattern.sname);
            var find = {sname: req.param('token')};
            break;
        default:
            e.push({param: 'token', msg: res.__('invalidParam', req.param('type'))});
    }
    if (e) {
        res.json({errors: e});
        return;
    }
    global.mongo.collection('user', function (err, collection) {
        if (err) {
            res.json({errors: err});
        } else {
            collection.findOne(find, function (err, doc) {
                if (err) {
                    res.json({errors: err});
                    return;
                }
                if (!doc) {
                    res.json({errors: [{param: 'token', msg: res.__('userNotFound', req.param('type'))}]});
                    return;
                }
                if (doc) {
                    cb(doc);
                }
            });
        }
    });
}

actions.POST.logIn = function (req, res) {
    req.checkBody('pass', res.__('invalidPassword')).matches(/^\w{10}$/);
    checkUserArguments(req, res, function (d) {
        if (d.password != req.param('pass')) {
            res.json({errors: [{param: 'pass', msg: res.__('wrongPassword')}]});
            return;
        }
        req.session.user = d;
        res.json({success: true});
    });
};

actions.POST.forgotPassword = function (req, res) {
    checkUserArguments(req, res, function (d) {
        global.mongo.collection('user', function (err, collection) {
            if (err) {
                res.json({errors: err});
            } else {
                var p = require('crypto').randomBytes(5).toString('hex');
                collection.update(
                    {_id: d._id},
                    {$set : {password: p}},
                    {safe: true},
                    function (err, result) {
                        if (err) {
                            res.json({errors: err});
                        } else {
                            if (result) {
                                res.render(
                                    'guest/mailForgotPassword',
                                    {
                                        email    : d.email,
                                        sname    : d.sname,
                                        password : p,
                                        locale   : res.getLocale(),
                                    },
                                    function (err, html) {
                                        if (err) {
                                            res.json({errors: err});
                                            return;
                                        }
                                        global.mail.sendMail({
                                            to: d.email,
                                            subject: res.__('guestMailSubjectForgotPassword'),
                                            html: html,
                                        }, function (err, r) {
                                            if (err) {
                                                res.json({errors: err});
                                            } else {
                                                res.json({success: true});
                                            }
                                        });
                                    }
                                );
                            }
                        }
                    }
                );
            }
        });
    });
};

exports.go = function (req, res) {
    if (req.method in actions && req.param('action') in actions[req.method]) {
        actions[req.method][req.param('action')](req, res);
    } else {
        res.render('app', {title: new Date()});
    }
};
