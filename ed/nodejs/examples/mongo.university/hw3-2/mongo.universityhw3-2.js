let mongodb = require('mongodb');

mongodb.MongoClient.connect('mongodb://dbu:dbp@xmongo:27017/test', function (err, db) {
  if (err) console.error(err);

  var query = {};
  var cursor = db.collection('test').find(query);
  cursor.skip(6);
  cursor.limit(2);
  cursor.sort({'grade': 1});
      
  cursor.forEach(
      function(doc) {
          console.log(doc);
      },
      function(err) {
        console.log('ERROR', err);
      }
  );
});
