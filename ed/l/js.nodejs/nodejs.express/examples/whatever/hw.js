const express = require('express');
const app = express();

app.listen('8080');

app.get('/hw', (req, res) => {
  res.send('hello world');
});
