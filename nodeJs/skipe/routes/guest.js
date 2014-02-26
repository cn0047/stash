exports.registration = function (req, res) {
    req.checkBody('email', 'Invalid email').isEmail();
    req.checkBody('sname', 'Invalid sname').is(/^[\w\s\_\-\=\+@]+$/);
    var e = req.validationErrors();
    if (e) {
        res.json({errors: e});
        return;
    }
    // console.log(req.param('email'));
    // MONGO
};
