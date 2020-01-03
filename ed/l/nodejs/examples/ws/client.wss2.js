const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:5432/wss2');

ws.on('open', function open() {
  ws.send('Something from client.');
});

ws.on('message', function incoming(data) {
  console.log('Received: %s', data);
});
