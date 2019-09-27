/**
 *@example After `npm start` just open in browser url: http://localhost:8080
 */

var app = require('http').createServer(handler)
    , io = require('socket.io').listen(app)
    , fs = require('fs');

app.listen(8080);

function handler (req, res) {
    fs.readFile(__dirname+'/index.html', function (err, data) {
        res.end(data);
    });
}

io.sockets.on('connection', function (socket) {
    socket.on('edit', function (data) {
        socket.broadcast.emit('update', data);
    });
});
