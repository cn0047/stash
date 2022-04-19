/*
npm install node-validator
*/

var check = require('validator').check,
    sanitize = require('validator').sanitize;
try {
    check(email).isEmail();
} catch (err) {
    console.log(err.message); // Неверный электронный адрес
}

try {
    check(email, "Please enter a proper email").isEmail();
} catch (err) {
    console.log(err.message); // Введите, пожалуйста, правильный адрес
}

var newstr = sanitize(str).xss(); // предупреждение XSS-атаки

var email = 'shelleyp@burningbird.net';
var email2 = 'this is a test';
var str = '<SCRIPT SRC=http://ha.ckers.org/xss.js></SCRIPT>';
try {
    check(email).isEmail();
    check(email2).isEmail();
} catch (err) {
    console.log(err.message);
}
var newstr = sanitize(str).xss();
console.log(newstr);


app.get('/somepage', function (req, rest) {
    req.check('zip', 'Please enter zip code').isInt(6);
    req.sanitize('newdata').xss();
});