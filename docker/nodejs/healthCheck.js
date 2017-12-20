const http = require('http');

const app = http.createServer((req, res) => {
  res.end('It works from `Docker`!');
});

app.listen(3000);
