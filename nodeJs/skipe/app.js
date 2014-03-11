// nodemon ./app.js localhost 3000

var express = require('express');
var expressValidator = require('express-validator');
var mongodb = require('mongodb');
var mongoServer = new mongodb.Server('localhost', 27017, {auto_reconnect: true});
var mongoDB = new mongodb.Db('skipe', mongoServer);
var app = express();

app.configure(function () {
    app.use(express.static(__dirname + '/public'));
    app.use(express.bodyParser());
    app.use(expressValidator());
    app.set('views', __dirname + '/views');
    app.set('view engine', 'jade');
    app.set('mongoDB', mongoDB);
});

app.get('*', require('./routes/app').go);
app.post('/guest/registration', require('./routes/guest').registration);

app.listen(3000);
console.log('Listening on port 3000');
