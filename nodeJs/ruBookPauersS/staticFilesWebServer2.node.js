var http = require('http'),
    path = require('path'),
    fs = require('fs'),
    mime = require('mime'),
    base = '/home/examples/public_html';
http.createServer(function (req, res) {
    pathname = base + req.url;
    console.log(pathname);
    fs.stat(pathname, function(err, stats) {
        if (err) {
            res.writeHead(404);
            res.write('Bad request 404\n');
            res.end();
        } else if (stats.isFile()) {
            // тип контента
            var type = mime.lookup(pathname);
            console.log(type);
            res.setHeader('Content-Type', type);
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
        } else {
            res.writeHead(403);
            res.write('Directory access is forbidden');
            res.end();
        }
    });
}).listen(8124);
console.log('Server running at 8124/');