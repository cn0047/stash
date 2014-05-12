/*
npm install db-mysql

npm install mysql
*/

var mysql = require('db-mysql');
var db = new mysql.Database({
    hostname: 'localhost',
    user: 'username',
    password: 'userpass',
    database: 'databasenm'
});
db.connect();
db.on('error', function(error) {
    console.log("CONNECTION ERROR: " + error);
});
// база данных подключена
db.on('ready', function(server) {
    // запрос, использующий выстроенные в цепочку методы
    // и вложенную функцию обратного вызова
    this.query()
        .select('*')
        .from('nodetest2')
        .where('id = 1')
        .execute(function(error, rows, columns) {
            if (error) {
              return console.log('ERROR: ' + error);
            }
            console.log(rows);
            console.log(columns);
        });
    // запрос, использующий непосредственную строку запроса и событие
    var qry = this.query();
    qry.execute('select * from nodetest2 where id = 1');
    qry.on('success', function(rows, columns) {
        console.log(rows); // вывод возвращенных строк
        console.log(columns); // вывод возвращенных столбцов
    });
    qry.on('error', function(error) {
        console.log('ERROR: ' + error);
    });
});

qry.execute('update nodetest2 set title = "This is a better title" where id = 1');

qry.execute('update nodetest2 set title = ? where id = ?', ["This was a better title", 1]);
qry.on('success', function(result) {
    console.log(result);
});
qry.on('error', function(error) {
    console.log('ERROR: ' + error);
});

qry.execute(
    'insert into nodetest2 (title, text,created) values(?,?,NOW())',
    ['Fourth Entry','Fourth entry in series'],
    function(err,result) {
        if (err) {
          console.log(err);
        } else {
            console.log(result);
            var qry2 = db.query();
            qry2.execute(
                'update nodetest2 set title = ? where id = ?',
                ['Better title', 4],
                function(err,result) {
                    if(err) {
                        console.log(err);
                    } else {
                        console.log(result);
                        var qry3 = db.query();
                        qry3.execute(
                            'delete from nodetest2 where id = ?',
                            [4],
                            function(err, result) {
                                if(err) {
                                    console.log(err);
                                } else {
                                    console.log(result);
                                }
                            }
                        );
                    }
                }
            );
        }
    }
);

db.on('ready', function(server) {
    // запрос, использующий выстроенные в цепочку методы
    // и вложенные функции обратного вызова
    var qry = this.query();
    qry.insert(
        'nodetest2',['title','text','created'],
       ['Fourth Entry', 'Fourth entry in series', 'NOW()']
    ).execute(function(err,result) {
        if (err) {
            console.log(err);
        } else {
            console.log(result);
            var qry2 = db.query();
            qry2.update('nodetest2')
                .set({title: 'Better title'})
                .where('id = ?',[4])
                .execute(function(err, result) {
                    if(err) {
                        console.log(err);
                    } else {
                        console.log(result);
                    }
                });
        }
    });
});