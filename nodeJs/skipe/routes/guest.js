exports.registration = function (req, res) {
    req.checkBody('email', 'Invalid email').isEmail();
    req.checkBody('sname', 'Invalid sname').matches(/^[\w\s\_\-\=\+@]+$/); // or .is()
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
                    res.json({errors: [{param: 'email', msg: 'This email address has already exists.'}]});
                    return;
                }
                collection.insert({
                    email: req.param('email'),
                    sname: req.param('sname'),
                    password: require('crypto').randomBytes(10).toString('hex'),
                }, function (err, docs) {
                    if (err) {
                        res.json({errors: err});
                        return;
                    }
                    res.render('registrationEmail', {email: docs[0].email, password: docs[0].password}, function(err, html) {
                        if (err) {
                            res.json({errors: err});
                            return;
                        }
                        global.mail.sendMail({
                            to: docs[0].email,
                            subject: '☺ Your Skipe address, '+docs[0].email+', has been created!',
                            html: html,
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
