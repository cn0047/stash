socket.io
-
1.7.3

````
emit - will send to all the clients
broadcast - will send the message to all except sender
volatile messages - use this type in case when some message may be lost.
````

````
# Default room
io.on('connection', function(socket){
  socket.on('say to someone', function(id, msg){
    socket.broadcast.to(id).emit('my message', msg);
  });
});

# Custom namespaces
var nsp = io.of('/my-namespace');
# On the client side:
var socket = io('/my-namespace');
````
