/*
npm install mysql
*/
var mysql = require('mysql');
var sql_host = '192.168.2.20';
var sql_user = 'admin';
var sql_pass = 'passw0rd';
var link = mysql.createClient({host: sql_host, user: sql_user, password: sql_pass});

var sql_pre = 'myapp_';
var sql_stmt = 'CREATE TABLE `'+sql_pre+'users` (`id` int AUTO_INCREMENT KEY, `user` text)';
var created = false;
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        created = true;
    }
});
var sql_stmt = 'DROP TABLE `'+sql_pre+'users`';
var dropped = false;
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        dropped = true;
    }
});
var sql_stmt = 'SELECT user FROM `'.$sql_pre.'users`';
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        for (var r=0; r < rows.length; ++r) {
            var row = rows[r];
            var user = row['user'];
            console.log(user);
        }
    }
});
var sql_stmt = 'UPDATE `'+sql_pre+'users` SET `user`="jsmith" WHERE `user`="dhoward"';
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        var rows_updated = rows.affectedRows;
        console.log('Changed '+rows_updated+' rows.');
    }
});
var sql_pre = 'myapp_';
var sql_stmt = 'INSERT INTO `'+sql_pre+'users` (`id`, `user`) VALUES (0, "dhoward")';
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        var inserted_id = rows.insertId;
        console.log('Inserted row with id='+inserted_id+'.');
    }
});
var sql_stmt = 'DELETE FROM `'+sql_pre+'users` WHERE `user`="dhoward"';
link.query(sql_stmt, function(e, rows, f) {
    if (!e) {
        var rows_deleted = rows.affectedRows;
        console.log('Deleted '+rows_deleted +' rows.');
    }
});
var sql_stmt = 'SELECT id FROM `'+sql_pre+'users` WHERE `user`="dhoward"';
link.query(sql_stmt, function(e, rows, f) {
    if (!e && (rows.length === 1)) {
        var row = rows[0];
        var id= row['id'];
        console.log(id);
    } else {
        // an error occurred
    }
});