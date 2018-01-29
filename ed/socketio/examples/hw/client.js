const io = require('socket.io-client');

const socket = io('http://localhost:8080');
socket.on('connect', (socket) => {
  console.log('Connected');
});

socket.emit('hello', {target: 'server', sender: 'server'});

socket.on('world', (data) => {
  console.log('Got:', data);
});
