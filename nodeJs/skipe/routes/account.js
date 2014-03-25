var actions = {
    POST: {},
};

actions.POST.init = function (req, res) {
    res.render('account', {title: 'skipe'});
};

exports.go = function (req, res) {
    if (req.session.user) {
        if (req.method in actions && req.param('action') in actions[req.method]) {
            actions[req.method][req.param('action')](req, res);
        }
        // Init user account after login.
        if (req.method == 'POST' && req.param('action') == undefined) {
            actions.POST.init(req, res);
        }
    } else {
        res.redirect('/#guest/login');
    }
};
