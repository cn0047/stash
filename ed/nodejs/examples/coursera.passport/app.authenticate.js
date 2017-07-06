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

// var routes = require('./routes/index');
var users = require('./routes/users');
// var dishRouter = require('./routes/dish');
var dishRouter = require('./routes/dishRouter');
var promoRouter = require('./routes/promoRouter');
// var leaderRouter = require('./routes/leaderRouter');

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
// uncomment after placing your favicon in /public
//app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')));

app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());

// passport config
app.use(passport.initialize());

app.use(express.static(path.join(__dirname, 'public')));

// app.get('/', function(req,res,next){
//         res.end('It works!');
// });

// app.use('/', routes);
app.use('/users', users);
app.use('/dishes',dishRouter);
app.use('/promotions',promoRouter);
// app.use('/leadership',leaderRouter);

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

curl -XPOST 'http://localhost:3000/users/register' \
    -H 'Content-Type: application/json' -d \
    '{"username": "usr", "password": "pwd", "firstname": "si", "lastname": "ple"}'

curl -XPOST 'http://localhost:3000/users/register' \
    -H 'Content-Type: application/json' -d \
    '{"username": "admin", "password": "pwd", "firstname": "nosi", "lastname": "nople"}'

// go to mongo and assign admin role
db.users.update({username: "admin"}, { $set: {admin: true} })

curl -XPOST 'http://localhost:3000/users/login' \
    -H 'Content-Type: application/json' -d \
    '{"username": "usr", "password": "pwd"}'

curl -XPOST 'http://localhost:3000/users/login' \
    -H 'Content-Type: application/json' -d \
    '{"username": "admin", "password": "pwd"}'

export tokenU="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyIkX18iOnsic3RyaWN0TW9kZSI6dHJ1ZSwic2VsZWN0ZWQiOnt9LCJnZXR0ZXJzIjp7fSwiX2lkIjoiNTk1ZGY1YzdjODE2MTM3ZDQzMzFiZWE4Iiwid2FzUG9wdWxhdGVkIjpmYWxzZSwiYWN0aXZlUGF0aHMiOnsicGF0aHMiOnsiZmlyc3RuYW1lIjoiaW5pdCIsImxhc3RuYW1lIjoiaW5pdCIsImFkbWluIjoiaW5pdCIsIl9fdiI6ImluaXQiLCJ1c2VybmFtZSI6ImluaXQiLCJoYXNoIjoiaW5pdCIsInNhbHQiOiJpbml0IiwiX2lkIjoiaW5pdCJ9LCJzdGF0ZXMiOnsiaWdub3JlIjp7fSwiZGVmYXVsdCI6e30sImluaXQiOnsiX192Ijp0cnVlLCJmaXJzdG5hbWUiOnRydWUsImxhc3RuYW1lIjp0cnVlLCJhZG1pbiI6dHJ1ZSwidXNlcm5hbWUiOnRydWUsImhhc2giOnRydWUsInNhbHQiOnRydWUsIl9pZCI6dHJ1ZX0sIm1vZGlmeSI6e30sInJlcXVpcmUiOnt9fSwic3RhdGVOYW1lcyI6WyJyZXF1aXJlIiwibW9kaWZ5IiwiaW5pdCIsImRlZmF1bHQiLCJpZ25vcmUiXX0sInBhdGhzVG9TY29wZXMiOnt9LCJlbWl0dGVyIjp7ImRvbWFpbiI6bnVsbCwiX2V2ZW50cyI6e30sIl9ldmVudHNDb3VudCI6MCwiX21heExpc3RlbmVycyI6MH19LCJpc05ldyI6ZmFsc2UsIl9kb2MiOnsiZmlyc3RuYW1lIjoiIiwibGFzdG5hbWUiOiJwbGUiLCJhZG1pbiI6ZmFsc2UsIl9fdiI6MCwidXNlcm5hbWUiOiJ1c3IiLCJoYXNoIjoiNDc4YWQ0MmUwNGM5ZjhlMGNmZjBhZTliNWExM2Q1YzQzZjczZDkwNDE5OTIxMzQyN2MyNWM4OTBiODQyMDNjMjlmZDIxMGFkODIyNjNjOTIzMDAyOWQyMGI2ZTQxNzM4YjM2MjZmMjZmODFhODNkNzJmOGNlMmI1ZGI2OTBjMGVlN2YyMDI2MmUxMWNhZmFkNjFmZmI2MzRjZTdkOWUyYzc0ZjkyMmJkYzhjNzYyNzI4M2ZmNTdjYzIwNjQ1NWVjODA0YmU1Mzc2MjgzZTEyYTg0MTE1MDM5YTk2MDhhY2NiNTJmYmZkNDAxNGJmMjcwZTlkZDU5MDdhYjkyN2Q0MmNkYzRlOTdmMThkNWEyOGM2Yjk0NDhjNzU2ZjU1YTgzNDAyZmZjYTNmNjVlMTVlODY0MzAwMjBhMGUyY2ZhOWYxMzY0OTA5YWFkOGVlYjE3MjJmN2Y0ODVkM2E3NmM2Nzc3NTI5NWQ1YzI5OWE3ZTJkNjZiMTA1MjJjZjBhYTMyNzdlNDJmNWM3OGMxNjI1YmMwNTQzM2ExOTcwZTNlZTM3MWQ0YTQ1MTdiMjNiMDRiMzA0NjgwYmQ4MjRhNzIxMGEyOGI3YTYzMGNkZWNiZTU2NjM0NGY2Y2FjZWU0ZDU5YTBhODE3MGYzYjRlOGQxZGNkZDNiZmYxNjhkMTQ3MGViZmJjZWUyMzMzM2VhMDRiYWJlNGNlODViNzdjOTY3NGNiMjZiY2Y5ZTJhYWNkNDYyZjI1OWQ2MDY0OGViMWY3Y2U4ZGVlMGJiMmU2MDEzZTFiYWVjNDI1YTQyOGYwNjU1YzdiMjhmNjQ2YzdiMDMzM2I3MzJmODQ4MGQ1MTE0ZjI3ZDk5NmU3ODA1NmNmOWIwZDdiNTVlZGZlMTdmZjljNTYwYzhmZjU1ODQ3NzVhYjM4NzgzOTYzYjEzMTAyYWNhYmUzNTk3ODEzYmQyMzFlZTdjOThmNWIyZmZjMGE0YTNmOGZiMzU0YTM3YmJjOTQxM2E5YzA4MjdlNTg0NzE5MTg3YmI4YmEzNGUzYmY2YWVmNDgyZjIzMmIwMGQ1NDg0M2E4MTlmNzhlODU4OWFkMTZmYzJkMWY1MWI5Y2QwNmZkNTRhNzlhM2U2NGY3YmU2M2QyYzc4ZTBiNmUyNmUyYjM2MjRmMzFiMmFmZDJiMjFkZTk5NmYxNTFiZGFiY2I1N2EzNjBjZDQyOTcxOWQxYzI1Yzc0ODA5ZTk1OWZmMmE5ODQwNzQxNjI0ZDYzMGY5ZjMzMGRlNmEyNzJmZTg2YzAxYzU3Zjg1NjMzODBlMGFjY2U2NmM0ZTA0ODBlOTIzMDRjYzZkMmI1YzczYzQ5MmJhZDFiZDM3OWFiNjdlYSIsInNhbHQiOiJmZTFmZjMwNjUwZDIxZjIwYjE4MzM3ZmU4MjIzZTEyMDUyYWVlMzUzYTgxNWExNmZlZGJhZDlmNjYyZGRkYmVhIiwiX2lkIjoiNTk1ZGY1YzdjODE2MTM3ZDQzMzFiZWE4In0sIiRpbml0Ijp0cnVlLCJpYXQiOjE0OTkzMzAxOTksImV4cCI6MTQ5OTMzMzc5OX0.6PbAUYNMwZee-SnFjdXCnHypE6T2bXmasbvDcrlR6nA"
export tokenA="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyIkX18iOnsic3RyaWN0TW9kZSI6dHJ1ZSwic2VsZWN0ZWQiOnt9LCJnZXR0ZXJzIjp7fSwiX2lkIjoiNTk1ZGY2MDZjODE2MTM3ZDQzMzFiZWE5Iiwid2FzUG9wdWxhdGVkIjpmYWxzZSwiYWN0aXZlUGF0aHMiOnsicGF0aHMiOnsiZmlyc3RuYW1lIjoiaW5pdCIsImxhc3RuYW1lIjoiaW5pdCIsImFkbWluIjoiaW5pdCIsIl9fdiI6ImluaXQiLCJ1c2VybmFtZSI6ImluaXQiLCJoYXNoIjoiaW5pdCIsInNhbHQiOiJpbml0IiwiX2lkIjoiaW5pdCJ9LCJzdGF0ZXMiOnsiaWdub3JlIjp7fSwiZGVmYXVsdCI6e30sImluaXQiOnsiX192Ijp0cnVlLCJmaXJzdG5hbWUiOnRydWUsImxhc3RuYW1lIjp0cnVlLCJhZG1pbiI6dHJ1ZSwidXNlcm5hbWUiOnRydWUsImhhc2giOnRydWUsInNhbHQiOnRydWUsIl9pZCI6dHJ1ZX0sIm1vZGlmeSI6e30sInJlcXVpcmUiOnt9fSwic3RhdGVOYW1lcyI6WyJyZXF1aXJlIiwibW9kaWZ5IiwiaW5pdCIsImRlZmF1bHQiLCJpZ25vcmUiXX0sInBhdGhzVG9TY29wZXMiOnt9LCJlbWl0dGVyIjp7ImRvbWFpbiI6bnVsbCwiX2V2ZW50cyI6e30sIl9ldmVudHNDb3VudCI6MCwiX21heExpc3RlbmVycyI6MH19LCJpc05ldyI6ZmFsc2UsIl9kb2MiOnsiZmlyc3RuYW1lIjoibm9zaSIsImxhc3RuYW1lIjoibm9wbGUiLCJhZG1pbiI6dHJ1ZSwiX192IjowLCJ1c2VybmFtZSI6ImFkbWluIiwiaGFzaCI6IjgxODY2MmRjZTVkNDJlZmY3MTM1OTg2MzNmNmFiYjAwMjAzMDdiNmZiZDIyYTE2ZGEzYTlkYWI1NTY2Y2ZkNmYyM2QwNTRlNTBlZjJhZDE3MzA2Y2ViOTY0NWY4Y2E5MDg2MjlkZTJhMGQ5MDI4ODdiMDUxZTk1ZmVmNDQ3ZjU0NTM5YjIyZjI2MWM2ZGNhZGFmZWJmNzQxYzQ3MTA0N2VlOTEzNGMwYWMyMGZlYWRmZTUzYzdlNTU2YjEzMWMyMTQ1YjM1MWE5NDM1Nzk2YzNhNTE5OTEzOTg2OTcxYWJiYmMzMDFlMmIwMmUyMWEzZjJlMmJjNDMzM2Q5NzY0NWJlZTY3NDMyZjk3OGMwMDY3ZDgxOGE2MjFiMjRkZWFlM2FkMmQ5NTIzYzBkYjkyZjI0Y2FiNzEzNjhkZjNmOThhOGIwMjk5Njk3ZWM5NWVjOTQwOTA0ZmUyNmRiYTNjYjViYmJjYTI3NTE1YzliYzE3NzAyYWEyNDUxZDI3ZWE4Y2EyZjE2MjBkNzBhOGQ0YTQ4YjIyNGNjN2M3NmYxZWE3ZDMxNTRiZDMxMmZlNTFjMDkzMDQ4M2VhZWU5Yjg0ODFkMjkyMTVkMjIyYjUyNDQ1ZDMzNDkzODE3NjdmZDQ3YWYzYmIwMzRiMTVjZWNjODBlNzFmNWFhN2FhYjE2ZDJhMTliNWEzYWVkNDUzNzE0Zjk2MTRiOGMwMWU4NjI5OGZkYmVmMGYyNzU4MDQ5YmE5M2Y5ZmYxNDY4YTUyYjhlMjhiM2U5MDM5Yzk2MDc0NmE1N2E2ZGI1YmQ2N2E2NzI5N2JiMWQyNjhkMWViN2I1NzNiMTU1MzY5ZDBjYTFlYjAyMjllNjcyYTBmYzc2YjNhNjYwY2M5YWU5Yzc1NjJjZDY3Zjg2YTk1OTVjMjA4NjBlMmM3YmM1NDJkNmExMTdiZTUwNDg0MDkyYTk4ODQ5ZTcxM2M1OGU4ODdkYTZmYWY4Y2YzN2MyZTA5ZWM5OWNlOTVhMjk2N2QwZmQwN2IyMWUxMjg2NGI5YWY4NDk3ZjU0Nzk0ZjdmZjM3YjA0ODJiY2ZhYTQ2OGQ3YWIwZTRkMDY1NjdlZmFjYzA2Y2NiY2Q1YjhhMDNhMjNkOGY3MmZjMGQ2Y2YzMzNkZGYxYzI4MTU2OTIxMDliMjE4MTg1NDQxZGM2YmQ1Njk0NjhmZDI5MTY4MGU0NWU2NTVjZmM0YjI5ZDkxYzZlOTkwMzViMjNmNThjMmRkMjdlNzlhNGEwNWFjOTk2MmJkMzI3NzRhMGMyMTE2YzVkYWE0Mjc1NTdmZDQ4Y2U3YzdlZWE3ZDllZGExOTU4M2Y0OTQyYzg3YjFmY2Q1ZDAyNDk2ODEzNzEiLCJzYWx0IjoiM2FiMjgxOWFiZmI0M2VhMTBiZTIwOTVlZDVkZDg5NDkxODc5ZTk5MWI3ODYwZmUxMjNhMTlkOGFkOWExYjU2YyIsIl9pZCI6IjU5NWRmNjA2YzgxNjEzN2Q0MzMxYmVhOSJ9LCIkaW5pdCI6dHJ1ZSwiaWF0IjoxNDk5MzMwMjIyLCJleHAiOjE0OTkzMzM4MjJ9.yJmCETzL_fzQwATVGN4i1fb4e-K2IBav6_ehUsLKgNE"

===

curl 'http://localhost:3000/dishes' \
    -H 'x-access-token: '$tokenU

curl -XPOST 'http://localhost:3000/dishes' \
    -H 'x-access-token: '$tokenA \
    -H 'Content-Type: application/json' -d \
    '{"name": "Pasta", "image": "pasta img", "category": "food", "price": "$19.75", "description": "It pasta"}'

curl -XPOST 'http://localhost:3000/dishes' \
    -H 'x-access-token: '$tokenA \
    -H 'Content-Type: application/json' -d \
    '{"name": "pizza", "image": "pizza img", "category": "food", "price": "$19.75", "description": "It pizza"}'

===

curl -XPOST 'http://localhost:3000/dishes/595dfdd3f135340c6d617bc8/comments' \
    -H 'x-access-token: '$tokenU \
    -H 'Content-Type: application/json' -d \
    '{"rating": 5, "comment": "pasta is nice!"}'

===

curl 'http://localhost:3000/promotions' \
    -H 'x-access-token: '$tokenU

curl -XPOST http://localhost:3000/promotions \
    -H 'x-access-token: '$tokenU \
    -H 'Content-Type: application/json' -d \
    '{"name": "p1-1", "image": "pi1", "price": "1.20", "description": "p1d"}'

curl -XPOST http://localhost:3000/promotions \
    -H 'x-access-token: '$tokenA \
    -H 'Content-Type: application/json' -d \
    '{"name": "p11", "image": "pi1", "price": "1.20", "description": "p1d"}'

curl http://localhost:3000/promotions/595d78cad41a8828fa50e957

curl -XPUT http://localhost:3000/promotions/595d78cad41a8828fa50e957 \
    -H 'Content-Type: application/json' -d \
    '{"name": "p1UPDATED", "image": "pi1", "price": "1.20", "description": "p1d"}'

===

curl 'http://localhost:3000/users' \
    -H 'x-access-token: '$tokenU

*/
