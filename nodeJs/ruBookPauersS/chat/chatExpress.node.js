var express = require('express'),
    sio = require('socket.io'),
    http = require('http'),
    app = express();
var server = http.createServer(app);
app.configure(function () {
    app.use(express.static(__dirname + '/public'));
    app.use(app.router);
});
app.get('/', function (req, res) {
    res.send('hello');
});
var io = sio.listen(server);
server.listen(8124);
io.sockets.on('connection', function (socket) {
    socket.on('addme',function(username) {
        socket.username = username;
        socket.emit('chat', 'SERVER', 'You have connected');
        socket.broadcast.emit('chat', 'SERVER', username + ' is on deck');
    });
    socket.on('sendchat', function(data) {
        io.sockets.emit('chat', socket.username, data);
    });
    socket.on('disconnect', function() {
        io.sockets.emit('chat', 'SERVER', socket.username + ' has left the building');
    });
});