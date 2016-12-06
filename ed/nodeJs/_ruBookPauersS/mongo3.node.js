var mongodb = require('mongodb');
var server = new mongodb.Server('localhost', 27017, {auto_reconnect: true});
var db = new mongodb.Db('exampleDb', server);
// открытие соединения с базой данных
db.open(function(err, db) {
    if(!err) {
        // доступ к коллекции виджетов или создание этой коллекции
        db.collection('widgets', function(err, collection) {
            // удаление всех виджет-документов
            collection.remove(null,{safe : true}, function(err, result) {
                if (!err) {
                    // создание четырех записей
                    var widget1 = {
                        id: 1,
                        title : 'First Great widget',
                        desc : 'greatest widget of all',
                        price : 14.99,
                        type: 'A'
                    };
                    var widget2 = {
                        id: 2,
                        title : 'Second Great widget',
                        desc : 'second greatest widget of all',
                        price : 29.99, type: 'A'
                    };
                    var widget3 = {
                        id: 3,
                        title: 'third widget',
                        desc: 'third widget',
                        price : 45.00, type: 'B'
                    };
                    var widget4 = {
                        id: 4,
                        title: 'fourth widget',
                        desc: 'fourth widget',
                        price: 60.00, type: 'B'
                    };
                    collection.insert(
                        [widget1,widget2,widget3,widget4],
                        {safe : true},
                        function(err, result) {
                            if(err) {
                                console.log(err);
                            } else {
                                // возвращение всех документов
                                collection.find().toArray(function(err, docs) {
                                    console.log(docs);
                                    // закрытие базы данных
                                    db.close();
                                });
                            }
                        }
                    );
                }
            });
        });
    }
});

// возвращение всех документов, имеющих тип А
collection.find({type:'A'},{fields:{type:0}}).toArray(function(err, docs) {
    if(err) {
        console.log(err);
    } else {
        console.log(docs);
        // закрытие базы данных
        db.close();
    }
});

// возвращение только одного документа
collection.findOne({id:1},{fields:{title:1}}, function(err, doc) {
    if (err) {
        console.log(err);
    } else {
        console.log(doc);
        // закрытие базы данных
        db.close();
    }
});
