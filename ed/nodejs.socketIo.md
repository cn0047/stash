socket.io
-
1.7.3

````
emit - will send to all the clients
broadcast - will send the message to all except sender
volatile messages - use this type in case when some message may be lost.
````

## Server

Supports `multiplexing` - use 1 connection instead of two WebSocket\'s.

Namespace:

````
# Custom namespaces
var nsp = io.of('/my-namespace');
nsp.on('connection', (socket) => { console.log('someone connected'); });
const chat = io.of('/chat');
# On the client side:
var socket = io('/my-namespace');
````

Room:

````
# send message into particular room.
adminNamespace.to('level1').emit('an event', { some: 'data' });

# default room
io.on('connection', function(socket){
  socket.on('say to someone', function(id, msg){
    socket.to('room1').to('room2').emit('hello');
    # or
    socket.broadcast.to(id).emit('my message', msg);
  });
});

# rooms this client is in
socket.rooms

socket.join('room 237', () => {});
````

`socket.clients` - Gets a list of client IDs connected to this namespace (across all nodes if applicable).

## Client

````
socket.on('connect', () => {
  console.log(socket.id); // 'G5p5...'
});
````

## Debug

````
DEBUG=* node yourfile.js
# or in the browser:
localStorage.debug = '*';
````
