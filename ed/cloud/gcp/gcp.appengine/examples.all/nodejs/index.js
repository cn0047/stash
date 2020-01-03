const http = require('http');

const handler = (req, res) => {
  res.end('nodejs - ok');
};

const app = http.createServer(handler);

const PORT = process.env.PORT || 8080;
app.listen(PORT);
