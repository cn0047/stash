var mongoose = require('mongoose');
var express = require('express');
var morgan = require('morgan');

var dishRouter = require('./dishRouter.js');

var url = 'mongodb://localhost:27017/conFusion';
var url = 'mongodb://dbu:dbp@localhost:27017/conFusion';
mongoose.connect(url);
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', function () {
    // we're connected!
    console.log("Connected correctly to server");
});

var hostname = 'localhost';
var port = 3000;
var app = express();
app.use(morgan('dev'));
app.use('/dishes', dishRouter);
app.listen(port, hostname, function(){
  console.log(`Server running at http://${hostname}:${port}/`);
});

/*
curl 'http://localhost:3000/dishes'

curl -XPOST 'http://localhost:3000/dishes' \
    -H 'Content-Type: application/json' -d \
    '{"name": "Pasta", "image": "pasta img", "category": "food", "price": "$19.75", "description": "It pasta"}'

curl 'http://localhost:3000/dishes/595cf05e53a386738ac11422'

curl -XDELETE http://localhost:3000/dishes/595cf05e53a386738ac11422

curl -XPUT http://localhost:3000/dishes/595cf05e53a386738ac11422 \
    -H 'Content-Type: application/json' -d \
    '{"name": "Pasta 2", "image": "pasta img 2", "category": "food 2", "price": "$19.75", "description": "It pasta 2"}'

curl 'http://localhost:3000/dishes/595cf05e53a386738ac11422/comments'

curl -XPOST 'http://localhost:3000/dishes/595cf05e53a386738ac11422/comments' \
    -H 'Content-Type: application/json' -d \
    '{"rating": 3, "comment": "This is insane", "author": "Matt Daemon"}'

curl -XPUT 'http://localhost:3000/dishes/595cf05e53a386738ac11422/comments/595cf1f51513b97937759516' \
    -H 'Content-Type: application/json' -d \
    '{"rating": 3, "comment": "This is insane in most cases", "author": "Matt Daemon"}'

*/
