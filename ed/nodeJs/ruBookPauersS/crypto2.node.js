var mysql = require('mysql'),
    crypto = require('crypto');
var client = mysql.createClient({
    user: 'username',
    password: 'password'
});
client.query('USE databasenm');
var username = process.argv[2];
var password = process.argv[3];
client.query(
    'SELECT password, salt FROM user WHERE username = ?',
    [username],
    function(err, result, fields) {
        if (err) return console.log(err);
        var newhash = crypto.createHash('sha512')
            .update(result[0].salt + password)
            .digest('hex');
        if (result[0].password === newhash) {
            console.log("OK, you're cool.");
        } else {
            console.log("Your password is wrong. Try again.");
        }
        client.end();
    }
);

// node password.js Michael apple*frk13*