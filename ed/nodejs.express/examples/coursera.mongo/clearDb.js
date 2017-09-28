var mongoose = require('mongoose'),
    assert = require('assert');

var Model = require('./models/dishes');
var Model = require('./models/promotions');
var Model = require('./models/leadership');

// Connection URL
var url = 'mongodb://localhost:27017/conFusion';
var url = 'mongodb://dbu:dbp@xmongo:27017/conFusion';
var url = 'mongodb://dbu:dbp@localhost:27017/conFusion';
mongoose.connect(url);
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', function () {
    // we're connected!
    console.log("Connected correctly to server");

db.collection('dishes').drop(function () {
    db.collection('promotions').drop(function () {
        db.collection('leaderships').drop(function () {
            db.collection('users').drop(function () {
                db.close();
                console.log("Done.");
                process.exit();
            });
        });
    });
});
});
