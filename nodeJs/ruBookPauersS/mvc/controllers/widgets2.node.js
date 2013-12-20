var widgets = [{
    id : 1,
    name : "The Great Widget",
    price : 1000.00,
    desc: "A widget of great value"
}]
// индексный список виджетов в каталоге /widgets/
exports.index = function(req, res) {
    res.render('widgets/index', {title : 'Widgets', widgets : widgets});
};
// вывод формы для нового виджета
exports.new = function(req, res) {
    var filePath = require('path').normalize(__dirname + "/../public/widgets/new.html");
    res.sendfile(filePath);
};
// добавление виджета
exports.create = function(req, res) {
    // генерирование id виджета
    var indx = widgets.length + 1;
    // добавление виджета
    widgets[widgets.length] = {
        id : indx,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice),
        desc : req.body.widgetdesc
    };
    // вывод на консоль и подтверждение добавления в адрес пользователя
    console.log(widgets[indx-1]);
    res.render('widgets/added', {title: 'Widget Added', widget : widgets[indx-1]});
};
// показ виджета
exports.show = function(req, res) {
    var indx = parseInt(req.params.id) - 1;
    if (!widgets[indx])
        res.send('There is no widget with id of ' + req.params.id);
    else
        res.render('widgets/show', {title : 'Show Widget', widget : widgets[indx]});
};
// удаление виджета
exports.destroy = function(req, res) {
    var indx = req.params.id - 1;
    delete widgets[indx];
    console.log('deleted ' + req.params.id);
    res.send('deleted ' + req.params.id);
};
// вывод формы редактирования
exports.edit = function(req, res) {
    var indx = parseInt(req.params.id) - 1;
    res.render('widgets/edit', {
        title : 'Edit Widget',
        widget : widgets[indx]
    });
};
// обновление виджета
exports.update = function(req, res) {
    var indx = parseInt(req.params.id) - 1;
    widgets[indx] = {
        id : indx + 1,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice),
        desc : req.body.widgetdesc
    };
    console.log(widgets[indx]);
    res.render('widgets/added', {title: 'Widget Edited', widget : widgets[indx]});
};
