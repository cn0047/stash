// @example curl http://localhost:3000/file-upload -H "Content-Type: multipart/form-data" -F "file=@/Users/k/f.txt" -F "msg=MyFile"
const express = require('express')
const fileUpload = require('express-fileupload');

const app = express()
app.use(fileUpload());

app.post('/file-upload', function(req, res, next) {
  console.log(req.body.msg);
  console.log(req.files);
  res.send('ok');
  next();
});

app.listen(3000, function () {
  console.log('Example app listening on port 3000!')
})
