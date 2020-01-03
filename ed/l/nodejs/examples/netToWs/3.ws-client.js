/**
 * @example node ed/nodejs/examples/netToWs/3.ws-client.js
 *          And now: touch /tmp/test
 */

const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:5430');

ws.on('open', function open() {
  console.log('Established new [WebSocket] connection for CLIENT.');
});

ws.on('message', function incoming(data) {
  console.log('Received: %s', data);
});
