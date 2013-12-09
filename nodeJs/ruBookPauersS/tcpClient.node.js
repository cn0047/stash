var net = require('net');
var client = new net.Socket();
client.setEncoding('utf8');
// установка соединения с сервером
client.connect ('8124','localhost', function () {
    console.log('connected to server');
    client.write('Who needs a browser to communicate?');
});
// подготовка к вводу данных с терминала
process.stdin.resume();
// отправка данных при их получении на сервер
process.stdin.on('data', function (data) {
    client.write(data);
});
// при получении ответных данных вывод их на консоль
client.on('data',function(data) {
    console.log(data);
});
// при закрытиии сервера
client.on('close',function() {
    console.log('connection is closed');
});

// Run: nodejs tcpClient.node.js
// then type any text, this text will be printed in tcpServer terminal.