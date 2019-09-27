var Widget = require('../models/widget.js');
// индексный список виджетов в каталоге /widgets/
exports.index = function(req, res) {
    Widget.find({}, function(err, docs) {
        console.log(docs);
        res.render('widgets/index', {title : 'Widgets', widgets : docs});
    });
};
// вывод формы нового виджета
exports.new = function(req, res) {
    console.log(req.url);
    var filePath = require('path').normalize(__dirname + "/../public/widgets/new.html");
    res.sendfile(filePath);
};
// добавление виджета
exports.create = function(req, res) {
    var widget = {
        sn : req.body.widgetsn,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice),
        desc: req.body.widgetdesc
    };
    var widgetObj = new Widget(widget);
    widgetObj.save(function(err, data) {
        if (err) {
            res.send(err);
        } else {
            console.log(data);
            res.render('widgets/added', {title: 'Widget Added', widget: widget});
        }
    });
};
// вывод виджета
exports.show = function(req, res) {
    var sn = req.params.sn;
    Widget.findOne({sn : sn}, function(err, doc) {
        if (err)
            res.send('There is no widget with sn of ' + sn);
        else
            res.render('widgets/show', {title : 'Show Widget', widget : doc});
    });
};
// удаление виджета
exports.destroy = function(req, res) {
    var sn = req.params.sn;
    Widget.remove({sn : sn}, function(err) {
        if (err) {
            res.send('There is no widget with sn of ' + sn);
        } else {
            console.log('deleted ' + sn);
            res.send('deleted ' + sn);
        }
    });
};
// вывод формы редактирования
exports.edit = function(req, res) {
    var sn = req.params.sn;
    Widget.findOne({sn : sn}, function(err, doc) {
        console.log(doc);
        if(err)
            res.send('There is no widget with sn of ' + sn);
        else
            res.render('widgets/edit', {title : 'Edit Widget', widget : doc});
    });
};
// обновление виджета
exports.update = function(req, res) {
    var sn = req.params.sn;
    var widget = {
        sn : req.body.widgetsn,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice),
        desc : req.body.widgetdesc
    };
    Widget.update({sn : sn}, widget, function(err) {
        if (err)
            res.send('Problem occured with update' + err)
        else
            res.render('widgets/added', {title: 'Widget Edited', widget : widget})
    });
};