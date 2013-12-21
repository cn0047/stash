var http = require('http');
var async = require('async');
var redis = require('redis');
var jade = require('jade');
// настройка Jade-шаблона
var layout = require('fs').readFileSync(__dirname + '/score.jade', 'utf8');
var fn = jade.compile(layout, {filename: __dirname + '/score.jade'});
// запуск Redis-клиента
var client = redis.createClient();
// выбор пятой базы данных
client.select(5);
// вспомогательная функция
function makeCallbackFunc(member) {
    return function(callback) {
        client.hgetall(member, function(err, obj) {
           callback(err,obj);
        });
    };
}
http.createServer(function(req,res) {
    // первый фильтр из запроса значка
    if (req.url === '/favicon.ico') {
        res.writeHead(200, {'Content-Type': 'image/x-icon'} );
        res.end();
        return;
    }
    // получение показателей, выстраивание в обратном порядке
    // только первых пяти результатов
    client.zrevrange('Zowie!',0,4, function(err,result) {
        var scores;
        if (err) {
            console.log(err);
            res.end('Top scores not currently available, please check back');
            return;
        }
        // создание массива функций обратного вызова для вызова Async.series
        var callFunctions = new Array();
        // обработка результатов с помощью makeCallbackFunc с помещением
        // только что возвращенных результатов в массив
        for (var i = 0; i < result.length; i++) {
           callFunctions.push(makeCallbackFunc(result[i]));
        }
        // использование Async-метода series для обработки
        // каждого обратного вызова по очереди и возвращения
        // конечного результата в виде массива объектов
        async.series(callFunctions, function (err, result) {
            if (err) {
                console.log(err);
                res.end('Scores not available');
                return;
            }
            // передача массива объектов движку шаблонов
            var str = fn({scores : result});
            res.end(str);
        });
    });
}).listen(3000);
console.log('Server running on 3000/');