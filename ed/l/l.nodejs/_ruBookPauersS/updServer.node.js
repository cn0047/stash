var dgram = require('dgram');
var server = dgram.createSocket("udp4");
server.on ("message", function(msg, rinfo) {
    console.log("Message: " + msg + " от " + rinfo.address + ":"+ rinfo.port);
});
server.bind(8124);
