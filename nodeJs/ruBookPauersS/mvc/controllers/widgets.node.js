// добавление виджета
app.post('/widgets/add', function(req, res) {
    var indx = widgets.length + 1;
    widgets[widgets.length] = {
        id : indx,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice)
    };
    console.log(widgets[indx-1]);
    res.send('Widget ' + req.body.widgetname + ' added with id ' + indx);
});

// Эта функция превращается в функцию widgets.create :
// добавление виджета
exports.create = function(req, res) {
    var indx = widgets.length + 1;
    widgets[widgets.length] = {
        id : indx,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice)
    };
    console.log(widgets[indx-1]);
    res.send('Widget ' + req.body.widgetname + ' added with id ' + indx);
};

var widgets = [{
    id : 1,
    name : "The Great Widget",
    price : 1000.00
}];
// индексированный список виджетов в каталоге /widgets/
exports.index = function(req, res) {
    res.send(widgets);
};
// вывод формы для нового виджета
exports.new = function(req, res) {
    res.send('displaying new widget form');
};
// добавление виджета
exports.create = function(req, res) {
    var indx = widgets.length + 1;
    widgets[widgets.length] = {
        id : indx,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice)
    };
    console.log(widgets[indx-1]);
    res.send('Widget ' + req.body.widgetname + ' added with id ' + indx);
};
// вывод виджета
exports.show = function(req, res) {
    var indx = parseInt(req.params.id) - 1;
    if (!widgets[indx])
      res.send('There is no widget with id of ' + req.params.id);
    else
       res.send(widgets[indx]);
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
    res.send('displaying edit form');
};
// обновление виджета
exports.update = function(req, res) {
    var indx = parseInt(req.params.id) - 1;
    widgets[indx] = {
        id : indx,
        name : req.body.widgetname,
        price : parseFloat(req.body.widgetprice)
    };
    console.log(widgets[indx]);
    res.send ('Updated ' + req.params.id);
};
