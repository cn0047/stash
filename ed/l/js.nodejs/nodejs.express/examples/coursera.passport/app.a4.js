var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var mongoose = require('mongoose');
var passport = require('passport');
var authenticate = require('./authenticate');

var config = require('./config');

mongoose.connect(config.mongoUrl);
var db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', function () {
    // we're connected!
    console.log("Connected correctly to server");
});

var users = require('./routes/users');
var dishRouter = require('./routes/dishRouter');
var promoRouter = require('./routes/promoRouter');
var leaderRouter = require('./routes/leaderRouter');
var favoriteRouter = require('./routes/favoriteRouter');

var app = express();

// Secure traffic only
app.all('*', function(req, res, next){
  console.log('req start: ',req.secure, req.hostname, req.url, app.get('port'));
  if (req.secure) {
    return next();
  };
  res.redirect('https://'+req.hostname+':'+app.get('secPort')+req.url);
});

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());

// passport config
app.use(passport.initialize());

app.use(express.static(path.join(__dirname, 'public')));

app.use('/users', users);
app.use('/dishes',dishRouter);
app.use('/promotions',promoRouter);
app.use('/leadership',leaderRouter);
app.use('/favorites',favoriteRouter);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  var err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// error handlers
// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
  app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.json({
      message: err.message,
      error: err
    });
  });
}

// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
  res.status(err.status || 500);
  res.json({
    message: err.message,
    error: {}
  });
});

module.exports = app;

/*

token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyIkX18iOnsic3RyaWN0TW9kZSI6dHJ1ZSwic2VsZWN0ZWQiOnsiaGFzaCI6MCwic2FsdCI6MH0sImdldHRlcnMiOnt9LCJfaWQiOiI1OTVlNDRkZjI5OTk4NzNmYjBkMjVmOGMiLCJ3YXNQb3B1bGF0ZWQiOmZhbHNlLCJhY3RpdmVQYXRocyI6eyJwYXRocyI6eyJmaXJzdG5hbWUiOiJpbml0IiwibGFzdG5hbWUiOiJpbml0IiwiYWRtaW4iOiJpbml0IiwiX192IjoiaW5pdCIsInVzZXJuYW1lIjoiaW5pdCIsIk9hdXRoSWQiOiJpbml0IiwiT2F1dGhUb2tlbiI6ImluaXQiLCJfaWQiOiJpbml0In0sInN0YXRlcyI6eyJpZ25vcmUiOnt9LCJkZWZhdWx0Ijp7fSwiaW5pdCI6eyJfX3YiOnRydWUsImZpcnN0bmFtZSI6dHJ1ZSwibGFzdG5hbWUiOnRydWUsImFkbWluIjp0cnVlLCJ1c2VybmFtZSI6dHJ1ZSwiT2F1dGhJZCI6dHJ1ZSwiT2F1dGhUb2tlbiI6dHJ1ZSwiX2lkIjp0cnVlfSwibW9kaWZ5Ijp7fSwicmVxdWlyZSI6e319LCJzdGF0ZU5hbWVzIjpbInJlcXVpcmUiLCJtb2RpZnkiLCJpbml0IiwiZGVmYXVsdCIsImlnbm9yZSJdfSwicGF0aHNUb1Njb3BlcyI6e30sImVtaXR0ZXIiOnsiZG9tYWluIjpudWxsLCJfZXZlbnRzIjp7fSwiX2V2ZW50c0NvdW50IjowLCJfbWF4TGlzdGVuZXJzIjowfX0sImlzTmV3IjpmYWxzZSwiX2RvYyI6eyJmaXJzdG5hbWUiOiIiLCJsYXN0bmFtZSI6IiIsImFkbWluIjpmYWxzZSwiX192IjowLCJ1c2VybmFtZSI6IkphbWVzamFtZXNqYW1lcyBCb25kYm9uZGJvbmQiLCJPYXV0aElkIjoiNTEyMzMxNTIyMjk2MTk5IiwiT2F1dGhUb2tlbiI6IkVBQVk0amVvWkNLMHNCQU1ENlVIbEttMHFhb0FaQzYyS3lTRUlMankxRmYwMGlNY3pwT08wVGZqY1BJeEJtcjc0OGJ3R1dqU3ZiNXFLbzZPbXBubENFeU9WUmhlcTdQRWRqWkFYZjlGOEF0M0xES1FGWDdxSVFkbFpDdDN2ZU9OMkNmSDZNYm5aQjliU2JxZzZUU0JtRlkxNTRaQkZsT1YyTnY5ZUdjTkZKS0xBWkRaRCIsIl9pZCI6IjU5NWU0NGRmMjk5OTg3M2ZiMGQyNWY4YyJ9LCIkaW5pdCI6dHJ1ZSwiaWF0IjoxNDk5MzY1MjQwLCJleHAiOjE0OTkzNjg4NDB9.IcojpnMvL4p3oECH_6DbbGiMThRZMkIkFIgoxL_ampA"

curl -XGET -k 'https://localhost:3443/favorites' \
    -H 'x-access-token: '$token

curl -XPOST -k 'https://localhost:3443/favorites' \
    -H 'x-access-token: '$token \
    -H 'Content-Type: application/json' -d \
    '{"_id": "595dfdd3f135340c6d617bc8"}'

curl -XPOST -k 'https://localhost:3443/favorites' \
    -H 'x-access-token: '$token \
    -H 'Content-Type: application/json' -d \
    '{"_id": "595dfee9f15dd90f6c2e05d5"}'

curl -XDELETE -k 'https://localhost:3443/favorites/595dfdd3f135340c6d617bc8' \
    -H 'x-access-token: '$token

curl -XDELETE -k 'https://localhost:3443/favorites' \
    -H 'x-access-token: '$token

*/
