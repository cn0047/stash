var elasticsearch = require('elasticsearch');
var mysql = require('mysql');
var _ = require('underscore');

var es = new elasticsearch.Client({host: '127.0.0.1:9201', log: 'error'});
var db = mysql.createConnection({host: 'localhost', database: 'd', user: 'u', password: 'p'});
db.connect(function (err) {
  if (err) {
    console.error(err);
  }
});

loopDb();

function loop() {
  for (i = 0; 75000; i++) {
    console.log(i);
    getFromDb(i);
  }
}

function loopDb() {
  var sql = 'SELECT UserID AS userId FROM users LIMIT 150';
  db.query(sql, function (err, rows) {
    if (err) {
      console.error(err);
    }
    _.each(rows, function(r) {
      getFromES(r.userId);
    });
    // for (var i = 0; i < rows.length; i++) {
    //   var userId = rows[i].userId;
    //   console.log('DB', userId);
    //   getFromES(userId);
    //   sleep.sleep(1);
    // }
  });
}

function getFromDb(userId) {
  console.log('DB', userId);
  var sql = 'SELECT UserID AS userId, StatusID AS statusId FROM users WHERE UserID = ' + userId;
  db.query(sql, function (err, rows) {
    if (err) {
      console.error(err);
    }
    console.log(rows);
  });
}

function getFromES(userId) {
  console.log('ES', userId);
  es.search({
    index: 'ziipr',
    type: 'users',
    body: {
      query: {ids: {values: [userId]}}
    }
  }).then(function (resp) {
    var hits = resp.hits.hits;
    console.log(hits);
  }, function (err) {
    console.error(err.message);
  });
}
