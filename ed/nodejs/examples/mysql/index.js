var mysql = require('mysql');

var db = mysql.createConnection({
  host: 'mysql-master',
  database: 'test',
  user: 'dbu',
  password: 'dbp'
});

db.connect(function (err) {
  if (err) {
    console.error(err);
  }
  db.query('SELECT NOW() AS dateNow', function (err, rows) {
    if (err) {
      console.error(err);
    }
    for (var i = 0; i < rows.length; i++) {
      console.log('Date now from DB is: %s', rows[i].dateNow);
    }
  });
});
