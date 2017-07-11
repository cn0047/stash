var connect = require('connect');

function logger(req, res, next) {
    console.log('%s %s', req.method, req.url);
    next();
}
function hello(req, res) {
    res.setHeader('Content-Type', 'text/plain');
    res.end('hello world');
}

var app = connect()
    .use(logger)
    .use(hello)
    .listen(3000)
;

// curl http://localhost:3000/
