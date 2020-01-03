module.exports = function favicon(path, options){
    var options = options || {}
        , path = path || __dirname + '/../public/favicon.ico'
        , maxAge = options.maxAge || 86400000;
    return function favicon(req, res, next){
        if ('/favicon.ico' == req.url) {
            if (icon) {
                res.writeHead(200, icon.headers);
                res.end(icon.body);
            } else {
                fs.readFile(path, function(err, buf){
                    if (err) return next(err);
                    icon = {
                        headers: {
                            'Content-Type': 'image/x-icon'
                            , 'Content-Length': buf.length
                            , 'ETag': '"' + utils.md5(buf) + '"'
                            , 'Cache-Control': 'public, max-age=' + (maxAge / 1000)
                        },
                        body: buf
                    };
                    res.writeHead(200, icon.headers);
                    res.end(icon.body);
                });
            }
        } else {
             next();
        }
    };
};



http.createServer(connect()
    .use(connect.favicon('/public_html/favicon.ico'))
    .use(connect.logger())
    .use(connect.static(_dirname + '/public_html'))
).listen(8124);