// var events = require('events');
// var em = new events.EventEmitter();
var eventEmitter = require('events').EventEmitter;
var counter = 0;
var em = new eventEmitter();
setInterval(function() { em.emit('timed', counter++); }, 3000);
em.on('timed', function(data) {
    console.log('timed ' + data);
});
// При запуске приложения на консоль выводится сообщение для события timed до
// тех пор, пока приложение не завершит работу.

/*
util.inherits(someobj, EventEmitter);
someobj.prototype.somemethod = function() { this.emit('event'); };
someobjinstance.on('event', function() { });
*/