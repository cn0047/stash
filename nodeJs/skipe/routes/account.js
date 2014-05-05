var actions = {
    POST: {},
    GET: {},
};

actions.POST.init = function (req, res) {
    res.render('app', {title: 'skipe'});
};

actions.GET.initLayouts = function (req, res) {
    res.json({
        menu: {
            "<span class='glyphicon glyphicon-new-window'></span>": "activate",
            "<span class='glyphicon glyphicon-time'></span>": "recentChats",
            "<span class='glyphicon glyphicon-bookmark'></span>": "chatsBookmarks",
            "divider1": "divider",
            "<img src='/i/Circle_Green.png'>": "online",
            "<img src='/i/Circle_Orange.png'>": "outOfPlace",
            "<img src='/i/Circle_Red.png'> ": "doNotDisturb",
            "<img src='/i/Circle_Grey.png'> ": "invisible",
            "<img src='/i/Circle_Grey.png'> ": "offline",
            "divider2": "divider",
            "<span class='glyphicon glyphicon-wrench'></span>": "settings",
            "divider3": "divider",
            "<span class='glyphicon glyphicon-transfer'></span>": "changeAccount",
            "<span class='glyphicon glyphicon-log-out'></span>": "logOut",
        },
        locales: global.availableLocales,
    });
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
        res.redirect('/#guest/login');
    }
};
