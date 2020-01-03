const http = require('http');
const url = require('url');
const WebSocket = require('ws');

const server = http.createServer();
const wss1 = new WebSocket.Server({ noServer: true });
const wss2 = new WebSocket.Server({ noServer: true });

wss1.on('connection', function connection(ws) {
  console.log('[WSS1] Established new connection.');

  ws.on('message', function incoming(message) {
    console.log('[WSS1] Received: %s', message);
    ws.send('[WSS1] Something from server.');
  });
});

wss2.on('connection', function connection(ws) {
  console.log('[WSS2] Established new connection.');

  ws.on('message', function incoming(message) {
    console.log('[WSS2] Received: %s', message);
    ws.send('[WSS2] Something from server.');
  });
});

server.on('upgrade', (request, socket, head) => {
  const pathname = url.parse(request.url).pathname;
  console.log(`Upgrade with path: ${pathname}`);

  if (pathname === '/wss1') {
    wss1.handleUpgrade(request, socket, head, (ws) => {
      wss1.emit('connection', ws);
      console.log('Connected with path /wss1');
    });
  } else if (pathname === '/wss2') {
    wss2.handleUpgrade(request, socket, head, (ws) => {
      wss2.emit('connection', ws);
      console.log('Connected with path /wss2');
    });
  } else {
    socket.destroy();
  }
});

server.listen(5432, 'localhost');
