// nodemon ./app.js localhost 3000

var express          = require('express');
var expressValidator = require('express-validator');
var mongodb          = require('mongodb');
var mailer           = require('nodemailer');

var config           = require('./configs/main').config;

var mongoServer = new mongodb.Server(config.mongo.host, config.mongo.port, config.mongo.options);
var mongoDB = new mongodb.Db(config.mongo.base, mongoServer);
var app = express();
mongoDB.open(function (err, db) {
    if (err) {
        console.error(err);
    } else {
        global.mongo = db;
    }
});
global.mail = mailer.createTransport('SMTP', {
    service: 'Gmail',
    auth: {user: config.mail.user, pass: config.mail.password}
});
global.validator = {
    patterns: {
        sname: /^[\w\s\_\-\=\+@]+$/,
    }
};

app.configure(function () {
    app.use(express.static(__dirname + '/public'));
    app.use(express.bodyParser());
    app.use(express.cookieParser());
    app.use(express.session({secret: '10QW3456789023456789QWERTY'}));
    app.use(expressValidator());
    app.set('views', __dirname + '/views');
    app.set('view engine', 'jade');
});

app.all('/account/:action?', require('./routes/account').go);
app.all('/guest/:action?', require('./routes/guest').go);
app.get('*', require('./routes/app').go);

app.listen(3000);
console.log('Listening on port 3000');
