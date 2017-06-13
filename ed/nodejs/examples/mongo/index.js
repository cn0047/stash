let mongodb = require('mongodb');

mongodb.MongoClient.connect('mongodb://dbu:dbp@xmongo:27017/test', function (err, db) {
  if (err) {
    console.error(err);
  }

  global.mongo = db;
  global.mongo.ObjectID = mongodb.ObjectID;

  db.collection('test', function (err, collection) {
    if (err) {
      console.error(err);
    }
    collection.findOne({'code': 200}, function (err, doc) {
      if (err) {
        console.error(err);
      }
      console.log(doc);
    });
  });
});
