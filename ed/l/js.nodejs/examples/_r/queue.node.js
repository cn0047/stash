var spawn = require('./child_process').spawn;
var net = require('net');
var client = new net.Socket();
client.setEncoding('utf8');
// соединение с к TCP-сервером
client.connect ('3000','examples.burningbird.net', function() {
    console.log('connected to server');
});
// запуск дочернего процесса
var logs = spawn('tail', [
    '-f',
    '/home/main/logs/access.log',
    '/home/tech/logs/access.log',
    '/home/shelleypowers/logs/access.log',
    '/home/green/logs/access.log',
    '/home/puppies/logs/access.log'
]);
// обработка данных дочернего процесса
logs.stdout.setEncoding('utf8');
logs.stdout.on('data', function(data) {
    // URL-адрес ресурса
    var re = /GET\s(\S+)\sHTTP/g;
    // тест на изображение
    var re2 = /\.gif|\.png|\.jpg|\.svg/;
    // извлечение URL-адреса, тест на изображение
    // сохранение в Redis в случае нахождения
    var parts = re.exec(data);
    console.log(parts[1]);
    var tst = re2.test(parts[1]);
    if (tst) {
        client.write(parts[1]);
    }
});
logs.stderr.on('data', function(data) {
    console.log('stderr: ' + data);
});
logs.on('exit', function(code) {
    console.log('child process exited with code ' + code);
    client.end();
});
