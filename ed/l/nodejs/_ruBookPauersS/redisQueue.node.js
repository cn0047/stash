var net = require('net');
var redis = require('redis');
var server = net.createServer(function(conn) {
    console.log('connected');
    // создание Redis-клиента
    var client = redis.createClient();
    client.on('error', function(err) {
        console.log('Error ' + err);
    });
    // очередью изображений является шестая база данных
    client.select(6);
    // прослушивание входящих данных
    conn.on('data', function(data) {
        console.log(data + ' from ' + conn.remoteAddress + ' ' +
        conn.remotePort);
        // сохранение данных
        client.rpush('images',data);
    });
}).listen(3000);
server.on('close', function(err) {
    client.quit();
});
console.log('listening on port 3000');