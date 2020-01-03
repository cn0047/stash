var redis = require('redis');
// главная страница
exports.index = function(req, res){
    res.render('index', { title: 'Express' });
};
// статистика
exports.stats = function(req, res){
    var client = redis.createClient();
    client.select(2);
    // Redis-транзакция для сбора данных
    client.multi()
        .smembers('ip')
        .hgetall('myurls')
        .exec(function(err, results) {
            var ips = results[0];
            var urls = results[1];
            res.render('stats',{ title: 'Stats', ips : ips, urls : urls});
            client.quit();
        });
};