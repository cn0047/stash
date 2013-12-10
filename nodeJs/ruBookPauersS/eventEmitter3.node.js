var util = require('util');
var eventEmitter = require('events').EventEmitter;
var fs = require('fs');
exports.inputChecker = inputChecker;
function inputChecker(name, file) {
    this.name = name;
    this.writeStream = fs.createWriteStream('./' + file + '.txt', {
        'flags' : 'a',
        'encoding' : 'utf8',
        'mode' : 0666
    });
};
inputChecker.prototype.check = function check(input) {
    var self = this;
    var command = input.toString().trim().substr(0,3);
    if (command == 'wr:') {
        self.emit('write',input.substr(3,input.length));
    } else if (command == 'en:') {
        self.emit('end');
    } else {
        self.emit('echo',input);
    }
};
