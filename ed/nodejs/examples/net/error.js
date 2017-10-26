const net = require('net');

const client = net.connect({ host: 'h', port: 0 }, () => console.log(200));

client.on('close', function () {
  console.log('1', arguments);
});
client.on('end', function () {
  console.log('2', arguments);
});
client.on('lookup', function () {
  console.log('3', arguments);
});
client.on('timeout', function () {
  console.log('4', arguments);
});
client.on('error', function () {
  console.log('5', arguments);
});
