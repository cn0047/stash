const http = require('http');
const fs = require('fs');

const port = 3000;
const app = http.createServer((req,res) => {
    res.writeHead(200);
    if (req.url === '/') req.url = '/index.html';
    res.end(fs.readFileSync(__dirname + req.url));
});

app.listen(port);
