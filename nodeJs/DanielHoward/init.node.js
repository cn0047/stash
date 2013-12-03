exports.initGET = function(req, pre, cb) {
    pre._GET = {};
    var urlparts = req.url.split('?');
    if (urlparts.length >= 2) {
        var query = urlparts[urlparts.length-1].split('&');
        for (var p=0; p < query.length; ++p) {
            var pair = query[p].split('=');
            pre._GET[pair[0]] = pair[1];
        }
    }
    cb();
};
exports.initPOST = function(req, pre, cb) {
    pre._POST = {};
    var body = '';
    req.on('data', function(chunk) {
        body += chunk;
        if (body.length > 1e6) {
            req.connection.destroy();
        }
    });
    req.on('end', function() {
        var pairs = body.split('&');
        for (var p=0; p < pairs.length; ++p) {
            var pair = pairs[p].split('=');
            pre._POST[pair[0]] = pair[1];
        }
        cb();
    });
};
exports.initCOOKIE = function(req, pre, cb) {
    pre._COOKIE = {};
    if (req.headers.cookie) {
        var cookies = req.headers.cookie.split(';');
        for (var c=0; c < cookies.length; ++c) {
            var pair = cookies[c].split('=');
            pre._COOKIE[pair[0]] = pair[1];
        }
    }
    cb();
};
exports.initREQUEST = function(req, pre, cb) {
    pre._REQUEST = {};
    if (pre._GET) {
        for (var k in pre._GET) {
            pre._REQUEST[k] = pre._GET[k];
        }
    }
    if (pre._POST) {
        for (var k in pre._POST) {
            pre._REQUEST[k] = pre._POST[k];
        }
    }
    if (pre._COOKIE) {
        for (var k in pre._COOKIE) {
            pre._REQUEST[k] = pre._COOKIE[k];
        }
    }
    cb();
};
/** All the sessions of all the users. */
var sessions = {};
exports.initSESSION = function(req, pre, cb) {
    if ((typeof pre._COOKIE['NODESESSID']) == 'undefined') {
        var pool = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
        var newid = '';
        for (var i = 0; i < 26; ++i) {
            var r = Math.floor(Math.random() * pool.length);
            newid += pool.charAt(r);
        }
        pre._COOKIE['NODESESSID'] = newid;
        sessions[pre._COOKIE['NODESESSID']] = {};
    }
    var id = pre._COOKIE['NODESESSID'];
    if ((typeof sessions[id]) == 'undefined') {
        sessions[id] = {};
    }
    pre._SESSION = sessions[id];
    cb();
}