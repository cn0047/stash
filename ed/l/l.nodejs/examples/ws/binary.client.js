const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:5432/binary');

ws.on('open', function open() {
  const array = new Float32Array(5);
  for (var i = 0; i < array.length; ++i) {
    array[i] = i / 2;
  }
  ws.send(array);
});
