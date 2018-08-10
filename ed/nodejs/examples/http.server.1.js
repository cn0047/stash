const http = require('http');

const handler = (req, res) => {
  if (req.method === 'POST') {
    let body = '';
    req.on('data', chunk => body += chunk.toString());
    req.on('end', () => {
        const post = JSON.parse(body);
        console.log(post)
    });
  }
};

const app = http.createServer(handler);
app.listen(8000);
