// nodemon ./app.js localhost 3000

var express = require('express');
var app = express();

app.configure(function(){
    app.use(express.static(__dirname + '/public'));
    app.set('views', __dirname + '/views');
    app.set('view engine', 'jade');
});

app.get('*', require('./routes/app').go);

app.listen(3000);
console.log('Listening on port 3000');
