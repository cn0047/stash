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
    var content = '';
    var x = parseInt(pre._REQUEST['x']);
    x += 5;
    content += '<html><head></head><body>';
    content += 'The value of x plus 5 is '+x+'.';
    content += '</body></html>';
    res.writeHead(200, {'Content-Type': 'text/html'})
    res.end(content);
    cb();
}