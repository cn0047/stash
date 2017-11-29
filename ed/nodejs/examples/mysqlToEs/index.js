/*

docker run -it --rm -p 9200:9200 --name es elasticsearch:latest

docker run -it --rm -p 3307:3306 --name xmysql --hostname xmysql \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp mysql:latest

*/

const elasticsearch = require('elasticsearch');
const mysql = require('mysql');

const es = new elasticsearch.Client({host: '127.0.0.1:9200', log: 'error'});
const db = mysql.createConnection({host: 'localhost', port: 3307, database: 'test', user: 'dbu', password: 'dbp'});
db.connect(function (err) {
  if (err) {
    console.error(err);
  }
});

loopDb();


function loopDb() {
  const sql = 'SELECT * FROM user LIMIT 250';
  const query = db.query(sql);
  query.on('result', (row) => {
    db.pause();
    putUserIntoES(row);
    console.log('*');
    db.resume();
  });
}

const putUserIntoES = (data) => {
  const doc = {
    index: 'test',
    type: 'users',
    id: data.id,
    body: {
      doc: {
        first_name: data.first_name,
        last_name: data.last_name,
      }
    }
  };
  es.update(doc, (error, response) => {
    if (error) {
      if (error.displayName === 'NotFound' && error.status === 404) {
        doc.body = doc.body.doc;
        return es.create(doc, (err, response) => {
          if (err) throw err;
          console.log('📗');
        });
      }
      throw error;
    }
    console.log('✅');
  });
};

/**
 * @deprecated
 * @param userId
 */
function getFromDb(userId) {
  console.log('DB', userId);
  const sql = 'SELECT UserID AS userId, StatusID AS statusId FROM users WHERE UserID = ' + userId;
  db.query(sql, function (err, rows) {
    if (err) {
      console.error(err);
    }
    console.log(rows);
  });
}

/**
 * @deprecated
 * @param userId
 */
function getFromES(userId) {
  console.log('ES', userId);
  es.search({
    index: 'ziipr',
    type: 'users',
    body: {
      query: {ids: {values: [userId]}}
    }
  }).then(function (resp) {
    const hits = resp.hits.hits;
    console.log(hits);
  }, function (err) {
    console.error(err.message);
  });
}
