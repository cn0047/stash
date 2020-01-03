const handler = (req, res) => {
  fs.readFile(__dirname+'/client.socketIo.html', (err, data) => {
    res.end(data);
  });
};

const app = require('http').createServer(handler);
const io = require('socket.io').listen(app);
const fs = require('fs');

app.listen(8080);

io.sockets.on('connection', (socket) => {
  socket.emit('foo', {jsonrpc: '2.0', result: {code: 200, message: 'OK'}});
});
