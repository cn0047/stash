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
// создание очереди
q = client.createQueue();
// вставка
q.query(
    'INSERT INTO nodetest2 (title, text, created) ' + 'values(?,?,NOW())',
    ['Title for 8', 'Text for 8']
);
// обновление
q.query(
    'UPDATE nodetest2 SET title = ? WHERE title = ?',
    ['New Title for 8','Title for 8']
);
q.execute();
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