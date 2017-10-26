const fs = require('fs');

function handler (req, res) {
  // @example '/room1.html'
  // @example '/room2.html'
  const url = req.url;
  fs.readFile(__dirname + url, function (err, data) {
    res.end(data);
  });
}

const app = require('http').createServer(handler);
const io = require('socket.io').listen(app);
app.listen(8080);
console.log('Open in browser         🖥 http://localhost:8080/room1.html');
console.log('Open in another browser 🖥 http://localhost:8080/room2.html');

io.sockets.on('connection', (socket) => {
  socket.on('join-room', (room) => {
    socket.join(room);
  });

  socket.emit('msg', 'M1');
  socket.to(socket.room).emit('msg', 'M1.2');
  socket.broadcast.to(socket.room).emit('msg', 'M2');
  socket.broadcast.to('room1').emit('msg', 'M3 for room1'); // ✅
  socket.broadcast.to('room2').emit('msg', 'M3 for room2'); // ✅
  io.to('room1').emit('msg', 'M4 for room1.');
  io.to('room2').emit('msg', 'M4 for room2.');
});
