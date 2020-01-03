var http = require('http'),
    url = require('url'),
    fs = require('fs'),
    mime = require('mime');
function processRange(res,ranges,len) {
    var start, end;
    // извлечение диапазона от начала до конца
    var rangearray = ranges.split('-');
    start = parseInt(rangearray[0].substr(6));
    end = parseInt(rangearray[1]);
    if (isNaN(start)) start = 0;
    if (isNaN(end)) end = len −1;
    // начальное значение выходит за пределы размера файла
    if (start > len - 1) {
        res.setHeader('Content-Range', 'bytes */' + len);
        res.writeHead(416);
        res.end();
    }
    // конечное значение не может выходить за пределы размера файла
    if (end > len - 1)
        end = len - 1;
    return {start:start, end:end};
}
http.createServer(function (req, res) {
    pathname = __dirname + '/public' + req.url;
    fs.stat(pathname, function(err, stats) {
        if (err) {
            res.writeHead(404);
            res.write('Bad request 404\n');
            res.end();
        } else if (stats.isFile()) {
            var opt={};
            // предположение об отсутствии диапазона
            res.statusCode = 200;
            var len = stats.size;
            // у нас есть запрос диапазона
            if (req.headers.range) {
                opt = processRange(res,req.headers.range,len);
                // настройка размера
                len = opt.end - opt.start + 1;
                // изменение кода состояния на частичный диапазон
                res.statusCode = 206;
                // установка заголовка
                var ctstr = 'bytes ' + opt.start + '-' + opt.end + '/' + stats.size;
                res.setHeader('Content-Range', ctstr);
            }
            console.log('len ' + len);
            res.setHeader('Content-Length', len);
            // тип контента
            var type = mime.lookup(pathname);
            res.setHeader('Content-Type', type);
            res.setHeader('Accept-Ranges','bytes');
            // создание потока чтения и направление его в канал
            var file = fs.createReadStream(pathname,opt);
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