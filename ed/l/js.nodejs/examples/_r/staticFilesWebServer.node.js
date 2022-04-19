var http = require('http'),
    path = require('path'),
    fs = require('fs'),
    base = '/home/examples/public_html';
http.createServer(function (req, res) {
    pathname = base + req.url;
    console.log(pathname);
    path.exists(pathname, function(exists) {
        if (!exists) {
            res.writeHead(404);
            res.write('Bad request 404\n');
            res.end();
        } else {
            res.setHeader('Content-Type', 'text/html');
            // код состояния 200 найден, значит, ошибок нет
            res.statusCode = 200;
            // создание потока чтения и направление его в канал
            var file = fs.createReadStream(pathname);
            file.on("open", function() {
                file.pipe(res);
            });
            file.on("error", function(err) {
                console.log(err);
            });
        }
    });
}).listen(8124);
console.log('Server running at 8124/');