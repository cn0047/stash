var redis = require("redis"),
    http = require('http');
var messageServer = http.createServer();
// прослушивание входящих запросов
messageServer.on('request', function (req, res) {
    // сначало фильтрация запросов на значок
    if (req.url === '/favicon.ico') {
        res.writeHead(200, {'Content-Type': 'image/x-icon'} );
        res.end();
        return;
    }
    // создание Redis-клиента
    var client = redis.createClient();
    client.on('error', function (err) {
        console.log('Error ' + err);
    });
    // установка на шестую базу данных
    client.select(6);
    client.lpop('images', function(err, reply) {
        if(err) {
            return console.error('error response ' + err);
        }
        // если это данные
        if (reply) {
            res.write(reply + '\n');
        } else {
            res.write('End of queue\n');
        }
        res.end();
    });
    client.quit();
});
messageServer.listen(8124);
console.log('listening on 8124');