var dgram = require('dgram');
var client = dgram.createSocket("udp4");
// подготовка к вводу с терминала
process.stdin.resume();
process.stdin.on('data', function (data) {
    console.log(data.toString('utf8'));
    client.send(data, 0, data.length, 8124, "examples.burningbird.net",
    function (err, bytes) {
        if (err)
            console.log('error: ' + err);
        else
            console.log('successful');
    });
});
z