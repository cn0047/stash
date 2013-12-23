/*
npm install mysql
*/
var mysql = require('mysql');
var client = mysql.createClient({
    user: 'username',
    password: 'password'
});
client.query('USE databasenm');
// создание
client.query(
    'INSERT INTO nodetest2 ' + 'SET title = ?, text = ?, created = NOW()',
    ['A seventh item', 'This is a seventh item'],
    function(err, result) {
        if (err) {
            console.log(err);
        } else {
            var id = result.insertId;
            console.log(result.insertId);
            // обновление
            client.query(
                'UPDATE nodetest2 SET ' + 'title = ? WHERE ID = ?',
                ['New title', id],
                function (err, result) {
                    if (err) {
                        console.log(err);
                    } else {
                        console.log(result.affectedRows);
                        // удаление
                        client.query(
                            'DELETE FROM nodetest2 WHERE id = ?',
                            [id],
                            function(err, result) {
                                if(err) {
                                    console.log(err);
                                } else {
                                    console.log(result.affectedRows);
                                    // использование вместо вложенной функции обратного вызова
                                    // именованной функции
                                    getData();
                                }
                            }
                        );
                    }
                }
            );
        }
    }
);
// извлечение данных
function getData() {
    client.query(
        'SELECT * FROM nodetest2 ORDER BY id',
        function(err, result,fields) {
            if(err) {
                console.log(err);
            } else {
                console.log(result);
                console.log(fields);
            }
            client.end();
        }
    );
}

/*
id - int(11), primary key, not null
title - varchar(255), unique key, not null
text - text, nulls allowed,
created - datetime, nulls allowed

type          Тип данных поля.
allowNull     Значение параметра false допускает наличие значений null ; по умолчанию параметр имеет значение true .
unique        Значение true не позволяет иметь повторяющиеся значения; по умолчанию параметр имеет значение false .
primaryKey    Значение true задает первичный ключ.
autoIncrement Значение true приводит к автоматическому приращению значения поля.
*/
var Nodetest2 = sequelize.define('nodetest2', {
    id : {type: Sequelize.INTEGER, primaryKey: true},
    title : {type: Sequelize.STRING, allowNull: false, unique: true},
    text : Sequelize.TEXT,
    created : Sequelize.DATE
});
// синхронизация
Nodetest2.sync().error(function(err) {
    console.log(err);
});