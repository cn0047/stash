const WebSocket = require('ws');

const wss = new WebSocket.Server({ port: 5432 });

wss.on('connection', function connection(ws) {
  console.log('Established new binary connection.');

  ws.on('message', function incoming(message) {
    console.log('Received message:');
    console.log(message);
    console.log(new Buffer(message, 'base64').toString());
    console.log(Buffer.from(message));
    const arrayBuffer = message.buffer.slice(
      message.byteOffset,
      message.byteOffset + message.byteLength
    );
    const float32Array = new Float32Array(arrayBuffer);
    console.log(float32Array);
  });
});
