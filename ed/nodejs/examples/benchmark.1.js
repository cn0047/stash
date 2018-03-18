const http = require('http');

http.createServer((req, res) => {
    if (req.url.startsWith('/v1/id/')) {
        if (req.method != 'GET') {
            res.end('{error:"501 Not Implemented"}');
            return;
        }

        const url = req.url;
        const p = url.lastIndexOf('/') + 1;
        const id = url.substring(p, url.length)

        res.end(JSON.stringify({Id: id, Msg: "OK."}));
    }
}).listen(8080, '0.0.0.0');
