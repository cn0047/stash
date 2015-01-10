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

exports.go = function (req, res) {
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
