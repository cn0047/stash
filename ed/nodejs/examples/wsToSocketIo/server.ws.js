const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 8080 });

wss.on('connection', (ws) => {
  ws.send(JSON.stringify({event: 'foo', jsonrpc: '2.0', result: {code: 200, message: 'OK'}}));
});
