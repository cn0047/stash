var mysql = require('mysql');
var queues = require('mysql-queues');
// подключение к базе данных
var client = mysql.createClient({
    user: 'username',
    password: 'password'
});
client.query('USE databasenm');
// связывание очередей с запросом
// использование отладки
queues(client, true);
// создание транзакции
var trans = client.startTransaction();
// вставка
trans.query(
    'INSERT INTO nodetest2 (title, text, created) ' + 'values(?,?,NOW())',
    ['Title for 8', 'Text for 8'],
    function(err,info) {
        if (err) {
            trans.rollback();
        } else {
            console.log(info);
            // обновление
            trans.query(
                'UPDATE nodetest2 SET title = ? WHERE title = ?',
                ['Better Title for 8','Title for 8'],
                function(err,info) {
                    if(err) {
                        trans.rollback();
                    } else {
                        console.log(info);
                        trans.commit();
                    }
                }
            );
        }
    }
);
trans.execute();
// выборки не произойдет, пока не завершатся предыдущие запросы
client.query(
    'SELECT * FROM nodetest2 ORDER BY ID',
    function(err, result, fields) {
        if (err) {
            console.log(err);
        } else {
            // будут показаны все записи, включая самые новые
            console.log(result);
            client.end();
        }
    }
);