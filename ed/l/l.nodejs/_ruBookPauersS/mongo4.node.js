/*
$inc       Приращение значения поля на конкретную величину.
$set       Устанавливает поле, как было показано в предыдущем примере.
$unset     Удаление поля
$push      Добавление значения к массиву, если поле является массивом (с преобразованием поля в массив, если поле им не является).
$pushAll   Добавление нескольких значений к массиву.
$addToSet  Добавление к массиву, только если поле является массивом.
$pull      Удаление значения из массива.
$pullAll   Удаление из массива нескольких значений.
$rename    Переименование поля.
$bit       Выполнение поразрядной операции.
*/
var mongodb = require('mongodb');
var server = new mongodb.Server('localhost', 27017, {auto_reconnect: true});
var db = new mongodb.Db('exampleDb', server);
// открытие соединения с базой данных
db.open(function(err, db) {
    if(!err) {
        // доступ к коллекции виджетов или создание этой коллекции
        db.collection('widgets',function(err, collection) {
            // обновление
            collection.update(
                {id:4},
                {$set : {title: 'Super Bad Widget'}},
                {safe: true},
                function(err, result) {
                    if (err) {
                        console.log(err);
                    } else {
                        console.log(result);
                        // запрос на обновленнную запись
                        collection.findOne({id:4}, function(err, doc) {
                            if(!err) {
                                console.log(doc);
                                // закрытие базы данных
                                db.close();
                            }
                        });
                    }
                }
            );
        });
    }
});

// обновление
collection.findAndModify(
    {id:4},
    [[ti]],
    {$set : {title: 'Super Widget', desc: 'A really great widget'}},
    {new: true},
    function(err, doc) {
        if (err) {
            console.log(err);
        } else {
            console.log(doc);DB
        }
        db.close();
    }
);