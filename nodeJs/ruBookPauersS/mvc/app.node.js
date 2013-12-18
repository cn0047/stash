var express = require('express')
    ,routes = require('./routes')
    ,map = require('./maproutecontroller')
    ,http = require('http')
    ,app = express();
app.configure(function(){
    app.use(express.favicon());
    app.use(express.logger('dev'));
    app.use(express.staticCache({maxObjects: 100, maxLength: 512}));
    app.use(express.static(__dirname + '/public'));
    app.use(express.bodyParser());
    app.use(express.methodOverride());
    app.use(app.router);
    app.use(express.directory(__dirname + '/public'));
    app.use(function(req, res, next){
        throw new Error(req.url + ' not found');
    });
    app.use(function(err, req, res, next) {
        console.log(err);
        res.send(err.message);
    });
});
app.configure('development', function(){
    app.use(express.errorHandler());
});
app.get('/', routes.index);
var prefixes = ['widgets'];
// отображение маршрута на контроллер
prefixes.forEach(function(prefix) {
    map.mapRoute(app, prefix);
});
http.createServer(app).listen(3000);
console.log("Express server listening on port 3000");
