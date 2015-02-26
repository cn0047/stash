// nodemon ./app.js localhost 3000

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

mongodb.MongoClient.connect(config.mongo.url, function (err, db) {
    if (err) {
        console.error(err);
    }
    global.mongo = db;
    global.mongo.ObjectID = mongodb.ObjectID;
});
global.localesDirectory = './public/nls';
global.availableLocales = [];
fs.readdir(global.localesDirectory, function (err, files) {
    var locales = [];
    files.forEach(function (file) {
        global.availableLocales.push(file.replace('.json', ''));
    });
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

var app = express();
app.use(express.static('./public'));
app.use(express.urlencoded());
app.use(express.json());
app.use(express.cookieParser());
app.use(express.session({
    secret: config.sessionSecret,
}));
app.use(expressValidator());
app.use(i18n.init);
app.set('views', './views');
app.set('view engine', 'jade');
app.use(function (req, res, next) {
    next();
});
app.all('/account/:action?/chat/:chat/user/:user?', require('./routes/account').go);
app.all('/account/:action?/chat/:chat?', require('./routes/account').go);
app.all('/account/:action?/user/:user?', require('./routes/account').go);
app.all('/account/:action?', require('./routes/account').go);
app.all('/guest/:action?', require('./routes/guest').go);
app.all('*', require('./routes/guest').go);

app.listen(3000);
console.log('Listening on port 3000...');
