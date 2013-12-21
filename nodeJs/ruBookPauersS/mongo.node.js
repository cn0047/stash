/*
npm install mongodb
npm install mongoose

При пакетной вставке документов вам нужно установить параметр keepGoing
в true , чтобы вставка документов продолжалась, даже если одна из вставок завер-
шится неудачей. По умолчанию если при вставке происходит ошибка, приложение
останавливает свою работу.

find          Возвращает курсор со всеми документами, возвращенными запросом.
findOne       Возвращает указатель на первый документ, соответствующий запросу.
findAndRemove Находит, а затем удаляет документ.
findAndModify Находит документ, а затем выполняет некое действие (например, remove или upsert ).
update        Либо обновляет документ, либо обновляет со вставкой (добавляет документ, если такого документа еще не существует).
remove        Удаляет документ.
findAndModify Находит и изменяет либо удаляет документ (возвращая модифицированный или удаленный документ).
findAndRemove Находит и удаляет документ (возвращая удаленный документ).
*/

var mongodb = require('mongodb');
var server = new mongodb.Server('localhost',:27017, {auto_reconnect: true});
var db = new mongdb.Db('mydb', server);
db.createCollection('mycollection', function(err, collection{});
db.createCollection('mycollection', {safe : true}, function(err, collection{}); 