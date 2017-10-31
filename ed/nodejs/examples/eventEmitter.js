// The EventListener calls all listeners synchronously in the order in which they were registered.
const EventEmitter = require('events');

class MyEmitter extends EventEmitter {}

const myEmitter = new MyEmitter();
myEmitter.on('event', (e) => {
  console.log('An event occurred!', e);
});

myEmitter.emit('event', {code: 204});
