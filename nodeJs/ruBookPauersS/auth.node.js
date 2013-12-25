var express = require('express')
    , flash = require('connect-flash')
    , passport = require('passport')
    , LocalStrategy = require('passport-local').Strategy
    , http = require('http');
var mysql = require('mysql')
    , crypto = require('crypto');
// проверка аутентификации пользователя
function ensureAuthenticated(req, res, next) {
    if (req.isAuthenticated()) { return next(); }
    res.redirect('/login')
}
// сериализация пользовательских данных для сеанса
passport.serializeUser(function(user, done) {
    done(null, user.id);
});
// поиск пользователя в базе данных MySQL
passport.deserializeUser(function(id, done) {
    var client = mysql.createClient({
        user : 'username',
        password: 'password'
    });
    client.query('USE databasenm');
    client.query(
        'SELECT username, password FROM user WHERE userid = ?',
        [id],
        function(err, result, fields) {
            var user = {
                id : id,
                username : result[0].username,
                password : result[0].password
            };
            done(err, user);
            client.end();
        }
    );
});
// настройка локальной стратегии аутентификации
// на пользовательскую запись в MySQL
passport.use(new LocalStrategy(function(username, password, done) {
    var client = mysql.createClient({
        user : 'username',
        password: 'password'
    });
    client.query('USE nodetest2');
    client.query(
        'SELECT userid, password, salt FROM user WHERE username = ?',
        [username],
        function(err, result, fields) {
            // ошибка базы данных
            if (err) {
                return done(err);
            // имя пользователя не найдено
            } else if (result.length == 0) {
                return done(null, false, {message: ' Unknown user ' + username});
                // проверка пароля
            } else {
                var newhash = crypto.createHash('sha512')
                    .update(result[0].salt + password)
                    .digest('hex');
                // если пароль соответствует
                if (result[0].password === newhash) {
                    var user = {
                        id : result[0].userid,
                        username : username,
                        password : newhash
                    };
                    return done(null, user);
                    // или если пароль не соответствует
                } else {
                    return done(null, false, {message: 'Invalid password'});
                }
            }
            client.end();
        }
    );
}));
var app = express();
app.configure(function(){
    app.set('views', __dirname + '/views');
    app.set('view engine', 'jade');
    app.use(express.favicon());
    app.use(express.logger('dev'));
    app.use(express.bodyParser());
    app.use(express.methodOverride());
    app.use(express.cookieParser('keyboard cat'));
    app.use(express.session());
    app.use(passport.initialize());
    app.use(passport.session());
    app.use(flash());
    app.use(app.router);
    app.use(express.static(__dirname + '/public'));
});
app.get('/', function(req, res){
    res.render('index', { title: 'authenticate', user: req.user });
});
app.get('/admin', ensureAuthenticated, function(req, res){
    res.render('admin', { title: 'authenticate', user: req.user });
});
app.get('/login', function(req, res){
    var username = req.user ? req.user.username : '';
    res.render('login', { title: 'authenticate', username: username, message: req.flash('error') });
});
app.post(
    '/login',
    passport.authenticate('local', { failureRedirect: '/login', failureFlash: true }),
    function(req, res) {
        res.redirect('/admin');
    }
);
http.createServer(app).listen(3000);
console.log("Express server listening on port 3000");