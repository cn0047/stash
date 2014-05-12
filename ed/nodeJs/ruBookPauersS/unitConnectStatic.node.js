var connect = require('connect'),
    http = require('http'),
    __dirname = '/home/examples';
http.createServer(
    connect()
    .use(connect.logger())
    .use(connect.static(_dirname + '/public_html'), {redirect: true})
).listen(8124);