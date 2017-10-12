/**
 * @example touch /tmp/test && node ed/nodejs/examples/netToWs/1.net-server.js /tmp/test
 *          Now in next terminal tab: telnet localhost 5431
 *          And now: touch /tmp/test
 */

const fs = require('fs');
const net = require('net');

const filename = process.argv[2];

if (!filename) {
  throw Error('No target filename was specified.');
}

const server = net.createServer(function(connection) {
  // reporting
  console.log('Subscriber connected.');
  connection.write(JSON.stringify({ type: 'watching', file: filename }) + '\n');

  // watcher setup
  let watcher = fs.watch(filename, function() {
    const m = JSON.stringify({ type: 'changed', file: filename, timestamp: Date.now() });
    console.log(`Sent message: ${m}.`);
    connection.write(`${m}\n`);
  });

  // cleanup
  connection.on('close', function() {
    console.log('Subscriber disconnected.');
    watcher.close();
  });
});

server.listen(5431, function() {
  console.log('Listening for subscribers...');
});
