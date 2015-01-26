// nodemon ./snippets.js

var mongodb = require('mongodb');
var config  = require('./configs/main').config;

mongodb.MongoClient.connect(config.mongo.url, function (err, db) {
    if (err) {
        console.error(err);
    }
    /*
    db.collection('user', function(err, collection) {
        collection.find().toArray(function(err, docs) {
            console.log(docs);
        });
    });
    db.collection('contact', function(err, collection) {
        collection.find({"owner": "James Bond"}).toArray(function(err, docs) {
            console.log(docs);
        });
    });
    */
    db.collection('post', function(err, collection) {
        collection.find().toArray(function(err, docs) {
            docs.map(function (d) {
                db.dereference(d.chat, function(err, item) {
                    console.log(d);
                    console.log(item);
                });
            });
        });
    });
});
