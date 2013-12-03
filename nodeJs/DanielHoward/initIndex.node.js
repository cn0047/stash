var initreq = require('./init.node.js');
exports.serve = function(req, res) {
    var pre = {};
    initGET(req, pre, function() {
        initPOST(req, pre, function() {
            initCOOKIE(req, pre, function() {
                initREQUEST(req, pre, function() {
                    initSESSION(req, pre, function() {
                        page(req, res, pre, function() {
                            var cookies = [];
                            for ( var c in pre._COOKIE) {
                                cookies.push(c + '=' + pre._COOKIE[c]);
                            }
                            res.setHeader('Set-Cookie', cookies);
                            res.writeHead(200, {'Content-Type': 'text/plain'});
                            res.end(res.content);
                        });
                    });
                });
            });
        });
    });
};
function page(req, res, pre, cb) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.end('admin/index.njs\n'+util.inspect(pre));
    cb();
}