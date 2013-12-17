var connect = require('connect')
    , http = require('http');
// очистка всех данных сеанса
function clearSession(req, res, next) {
    if ('/clear' == req.url) {
        req.session = null;
        res.statusCode = 302;
        res.setHeader('Location', '/');
        res.end();
    } else {
        next();
    }
}
// отслеживание пользователя
function trackUser(req, res, next) {
    req.session.ct = req.session.ct || 0;
    req.session.username = req.session.username || req.cookies.username;
    console.log(
        req.cookies.username + ' requested ' +
        req.session.ct++ + ' resources this session'
    );
    next();
}
// cookie-файлы и сеанс
var app = connect()
    .use(connect.logger('dev'))
    .use(connect.cookieParser('mumble'))
    .use(connect.cookieSession({key : 'tracking'}))
    .use(clearSession)
    .use(trackUser);
// статический сервер
app.use(connect.static('/home/examples/public_html'));
// запуск сервера и прослушивание
http.createServer(app).listen(8124);
console.log('Server listening on port 8124');