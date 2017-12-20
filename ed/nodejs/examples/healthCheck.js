const http = require('http');

const app = http.createServer((req, res) => {
  res.end('It works!');
});

app.listen(8080);
