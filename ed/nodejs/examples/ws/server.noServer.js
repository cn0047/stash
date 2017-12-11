const WebSocket = require('ws');
const { createServer } = require('http');

const server = createServer((req, res) => {});
server.listen(5432);

server.on('upgrade', async (request, socket, head) => {
  const wss = new WebSocket.Server({ perMessageDeflate: false, noServer: true });

  wss.on('connection', function connection(ws) {
    console.log('Established new connection.');

    ws.on('message', function incoming(message) {
      console.log('Received: %s', message);
      ws.send('Something from server.');
    });
  });

  wss.handleUpgrade(request, socket, head, (ws) => wss.emit('connection', ws));
});
