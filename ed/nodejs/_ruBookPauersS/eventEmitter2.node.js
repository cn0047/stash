var util = require('util');
var eventEmitter = require('events').EventEmitter;
var fs = require('fs');
function inputChecker (name, file) {
    this.name = name;
    this.writeStream = fs.createWriteStream('./' + file + '.txt', {
        'flags' : 'a',
        'encoding' : 'utf8',
        'mode' : 0666
    });
};
util.inherits(inputChecker,eventEmitter);
inputChecker.prototype.check = function check(input) {
    var command = input.toString().trim().substr(0,3);
    if (command == 'wr:') {
        this.emit('write',input.substr(3,input.length));
    } else if (command == 'en:') {
        this.emit('end');
    } else {
        this.emit('echo',input);
    }
};
// проверка нового объекта и обработка событий
var ic = new inputChecker('Shelley','output');
ic.on('write', function(data) {
    this.writeStream.write(data, 'utf8');
});
ic.on('echo', function( data) {
    console.log(this.name + ' wrote ' + data);
});
ic.on('end', function() {
    process.exit();
});
process.stdin.resume();
process.stdin.setEncoding('utf8');
process.stdin.on('data', function(input) {
    ic.check(input);
});

// А с помощью такого кода можно прослушивать только первое событие:
// ic.once(event, function);

// Когда количество слушателей события оказывается больше десяти, то по умолча-
// нию вы получаете предупреждение. Для изменения количества слушателей нужно
// их число передать функции setMaxListeners.
