var redis = require('redis');
module.exports = function getStats() {
    return function getStats(req, res, next) {
        // создание Redis-клиента
        var client = redis.createClient();
        client.on('error', function (err) {
            console.log('Error ' + err);
        });
        // Установка на вторую базу данных
        client.select(2);
        // добавление IP к набору
        client.sadd('ip',req.socket.remoteAddress);
        // приращение значения счетчика обращений к ресурсам
        client.hincrby('myurls',req.url, 1);
        client.quit();
        next();
    }
}