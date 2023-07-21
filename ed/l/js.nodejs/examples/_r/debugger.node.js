httpProxy.createServer(function(req,res,proxy) {
    debugger;
    if (req.url.match(/^\/node\//))
        proxy.proxyRequest(req, res, {host: 'localhost', port: 8000 });
    else
        proxy.proxyRequest(req, res, {host: 'localhost', port: 8124 });
}).listen(9000);

