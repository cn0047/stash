var app = require('http').createServer(handler)
    , io = require('socket.io').listen(app)
    , fs = require('fs');
var counter;
app.listen(8124);
function handler (req, res) {
    console.log(__dirname)
    fs.readFile(__dirname + '/socket.html', function (err, data) {
        if (err) {
            res.writeHead(500);
            return res.end('Error loading socket.html');
        }
        res.writeHead(200);
        res.end(data);
    });
}
io.sockets.on('connection', function (socket) {
    socket.counter = 1;
    socket.emit('news', { news: 'world' });
    socket.on('echo', function (data) {
        if (socket.counter <= 50) {
            data.back+=socket.counter;
            socket.counter++;
            socket.emit('news', {news: data.back});
        }
    });
});