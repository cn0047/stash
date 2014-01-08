var connect = require('connect');
var http = require('http');
var app = connect()
    .use(connect.favicon())
    .use(connect.logger())
    .use(function(req,res) {
        res.end('Hello World\n');
    });
http.createServer(app).listen(8124);



var connect = require('connect')
    , http = require('http');
var app = connect()
    .use(connect.logger('dev'))
    .use(connect.cookieParser())
    .use(function(req, res, next) {
        console.log('tracking ' + req.cookies.username);
        next();
    })
    .use(connect.static('/home/examples/public_html'));
http.createServer(app).listen(8124);
console.log('Server listening on port 8124');