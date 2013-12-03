var http = require('http');
var static = require('node-static');
var file = new static.Server();
var url = require('url');
var showx5 = require('./showX5.node.js');
http.createServer(function (req, res) {
    if (url.parse(req.url).pathname == '/showx5.php') {
        showx5.serve(req, res);
    } else {
        file.serve(req, res);
    }
}).listen(1337, '127.0.0.1');
console.log('Server running at http://127.0.0.1:1337/');

// Type in browser:
// http://localhost/showx5.php?x=22
// http://localhost:1337/showx5.php?x=22