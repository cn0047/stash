const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 5432 });

wss.on('connection', function connection(ws) {
  console.log('Established new connection.');

  ws.on('message', function incoming(message) {
    console.log('Received: %s', message);
    ws.send('Something from server.');
  });
});
