/**
 * @example node ed/nodejs/examples/netToWs/2.net-client.ws-server.js
 *          And now: touch /tmp/test
 */

const net = require('net');
const netClient = net.connect({port: 5431});

const WebSocket = require('ws');
const wss = new WebSocket.Server({ port: 5430 });

wss.on('connection', function connection(ws) {
  console.log('Established new [WebSocket] connection.');

  netClient.on('data', function(data) {
    console.log(`Sent data into [WebSocket] which received from [NET], Data: ${data}`);
    ws.send(data);
  });
});
