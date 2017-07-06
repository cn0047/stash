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

    var doc = {
        name: 'Uthapizza',
        image: 'img',
        category:'food' ,
        price: '$1,200.55"',
        description: 'Test',
        comments: [
            {
                rating: 3,
                comment: 'This is insane',
                author: 'Matt Daemon'
            }
        ]
    };
    var doc = {
        name: 'Promo1',
        image: 'imgP1',
        price: '$3.50"',
        description: 'Test promo'
    };
    var doc = {
        name: 'Lead1.1',
        image: 'imgL1',
        designation: 'designation1',
        abbr: 'L1',
        description: 'Test leadership'
    };
    Model.create(doc, function (err, d) {
        if (err) throw err;
        console.log('created!', d);

        var id = d._id;

        Model.findById(id, function (err, d) {
            if (err) throw err;
            // console.log('FOUND:', d, 'PRICE:', d.price.toFixed(2));
            console.log('FOUND:', d);
            db.collection('dishes').drop(function () {
                db.collection('promotions').drop(function () {
                    db.collection('leaderships').drop(function () {
                        db.close();
                    });
                });
            });
            process.exit();
        });

    });
});
