// nodemon --watch ./app.js localhost 3000

/**
 * @todo Move available locales to front.
 * @todo Fix i18n.
 */
var express          = require('express');
var expressValidator = require('express-validator');
var mongodb          = require('mongodb');
var mailer           = require('nodemailer');
var i18n             = require('i18n');
var fs               = require('fs');

var config           = require('./configs/main').config;

var mongoServer = new mongodb.Server(config.mongo.host, config.mongo.port, config.mongo.options);
var mongoDB = new mongodb.Db(config.mongo.base, mongoServer);
var app = express();
global.localesDirectory = './public/nls';
global.availableLocales = [];
fs.readdir(global.localesDirectory, function (err, files) {
    var locales = [];
    files.forEach(function (file) {
        global.availableLocales.push(file.replace('.json', ''));
    });
});
mongoDB.open(function (err, db) {
    if (err) {
        console.error(err);
    } else {
        global.mongo = db;
    }
});
global.mail = mailer.createTransport(config.mail.type, {
    service: config.mail.service,
    auth: {user: config.mail.user, pass: config.mail.password}
});
global.validator = {
    pattern: {
        sname: /^[\w\s\_\-\=\+@]+$/,
    }
};
global.demoUser = config.demoUser;
i18n.configure({
    cookie        : 'locale',
    locales       : global.availableLocales,
    directory     : global.localesDirectory,
    defaultLocale : config.defaultLocale,
    updateFiles   : false,
});

app.configure(function () {
    app.use(express.static('./public'));
    app.use(express.bodyParser());
    app.use(express.cookieParser());
    app.use(express.session({secret: config.sessionSecret}));
    app.use(expressValidator());
    app.use(i18n.init);
    app.set('views', './views');
    app.set('view engine', 'jade');
});

app.use(function (req, res, next) {
    next();
});
app.all('/account/:action?', require('./routes/account').go);
app.all('/guest/:action?', require('./routes/guest').go);
app.all('*', require('./routes/guest').go);

app.listen(3000);
console.log('Listening on port 3000...');
