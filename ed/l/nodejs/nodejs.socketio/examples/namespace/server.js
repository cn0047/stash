const fs = require('fs');

function handler (req, res) {
  // @example '/client1.html'
  // @example '/client2.html'
  const url = req.url;
  fs.readFile(__dirname + url, function (err, data) {
    res.end(data);
  });
}

const app = require('http').createServer(handler);
const io = require('socket.io').listen(app);
app.listen(8080);
console.log('Open in browser         🖥 http://localhost:8080/client1.html');
console.log('Open in another browser 🖥 http://localhost:8080/client2.html');

const nsp1 = io.of('/space1');
nsp1.on('connection', (socket) => {
  console.log('[space1] Someone connected.');

  socket.on('req', function (data) {
    console.log('[space1 req] ', data);
    socket.emit('res', 'I am server in space1.');
  });
});

const nsp2 = io.of('/space2');
nsp2.on('connection', (socket) => {
  console.log('[space2] Dude connected.');

  socket.on('req', function (data) {
    console.log('[space2 req] ', data);
    socket.emit('res', 'Yo dude, I am server in space2.');
  });
});
