const fs = require('fs');

function handler (req, res) {
  // @example '/client.html'
  const url = req.url;
  fs.readFile(__dirname + url, function (err, data) {
    res.end(data);
  });
}

const app = require('http').createServer(handler);
const io = require('socket.io').listen(app);
app.listen(8080);
console.log('Open in browser http://localhost:8080/client.html');

io.sockets.on('connection', function (socket) {
    socket.on('disconnect', function() {
      console.log('got disconnect.');
    });

    socket.on('hello', function (data) {
      console.log('Got data:', data);
      socket.emit('world', 'Hello world!');
    });
});
