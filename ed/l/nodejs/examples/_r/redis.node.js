var net = require('net');
var redis = require('redis');
var server = net.createServer(function(conn) {
    console.log('connected');
    // создание Redis-клиента
    var client = redis.createClient();
    client.on('error', function(err) {
        console.log('Error ' + err);
    });
    // пятая база данных, являющаяся базой данных показателей игры
    client.select(5);
    conn.on('data', function(data) {
        console.log(data + ' from ' + conn.remoteAddress + ' ' + conn.remotePort);
        try {
            var obj = JSON.parse(data);
            // добавление или перезапись показателя
            client.hset(obj.member, "first_name", obj.first_name, redis.print);
            client.hset(obj.member, "last_name", obj.last_name, redis.print);
            client.hset(obj.member, "score", obj.score, redis.print);
            client.hset(obj.member, "date", obj.date, redis.print);
            // добавление показателя для игры Zowie!
            client.zadd("Zowie!", parseInt(obj.score), obj.member);
        } catch(err) {
           console.log(err);
        }
    });
    conn.on('close', function() {
        console.log('клиент закрыл подключение');
        client.quit();
    });
}).listen(8124);
console.log('listening on port 8124');