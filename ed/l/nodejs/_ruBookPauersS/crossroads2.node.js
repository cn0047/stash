var crossroads = require('crossroads'),
    http = require('http');
var typeRoute = crossroads.addRoute('/{type}/{id}');
function onTypeAccess(type,id) {
    console.log('access ' + type + ' ' + id);
};
typeRoute.matched.add(onTypeAccess);
http.createServer(function(req,res) {
    crossroads.parse(req.url);
    res.end('processing');
}).listen(8124);