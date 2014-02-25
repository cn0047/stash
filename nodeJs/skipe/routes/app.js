exports.go = function (req, res) {
    res.render('app', {date: new Date()});
};
