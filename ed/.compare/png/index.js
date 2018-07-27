const http = require('http');

const handler = (req, res) => {
    console.log('[%s] New request', new Date());
    const t = (req.url === '/?busy') ? 7000 : 0;
    setTimeout(() => res.end('[nodejs] It works!'), t);
};

const app = http.createServer(handler);

app.listen(8000);
// curl 'http://localhost:8000/?busy'
