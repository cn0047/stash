const fs = require('fs');
const app = require('http').createServer(function (req, res) {
  fs.readFile(__dirname + '/dist' + req.url, function (err, data) {
    res.end(data);
  });
});

app.listen(8080);

// Now open in browser:
// http://localhost:8080/index.html
