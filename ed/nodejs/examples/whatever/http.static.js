const http = require('http');
const fs = require('fs');

const port = 3000;
const app = http.createServer((req,res) => {
    res.writeHead(200);
    switch (req.url) {
      case '/favicon.ico':
        return;
        break;
      case '/':
        return;
        break;
      case '/hw':
        req.url = '/hw.js';
        break;
      case '/ntf':
        req.url = '/../../../f/f.html/examples/notification.html';
        break;
      case '/ntf2':
        req.url = '/../../../f/f.html/examples/notification2.html';
        break;
    }
    res.end(fs.readFileSync(__dirname + req.url));
});

app.listen(port);
