var connect = require('connect'),
    http = require('http'),
    fs = require('fs'),
    crossroads = require('crossroads'),
    httpProxy = require('http-proxy'),
    base = '/home/examples/public_html';
// создание прокси-сервера, слушающего все запросы
httpProxy.createServer(function(req,res,proxy) {
    if (req.url.match(/^\/node\//))
        proxy.proxyRequest(req, res, {
            host: 'localhost',
            port: 8000
        });
    else
        proxy.proxyRequest(req,res, {
            host: 'localhost',
            port: 8124
        });
}).listen(9000);
// добавление маршрута для запросов динамического ресурса
crossroads.addRoute('/node/{id}/', function(id) {
    console.log('accessed node ' + id);
});
// динамический файловый сервер
http.createServer(function(req,res) {
    crossroads.parse(req.url);
    res.end('that\'s all!');
}).listen(8000);
// статический файловый сервер
http.createServer(connect()
    .use(connect.favicon())
    .use(connect.logger('dev'))
    .use(connect.static(base))
).listen(8124);



/*
/node/345
/example1.html
/node/800
/html5media/chapter2/example14.html
*/