var express = require('express');
var bodyParser = require('body-parser');

var leaderRouter = express.Router();

leaderRouter.use(bodyParser.json());

leaderRouter.route('/')
    .all(function(req,res,next) {
          res.writeHead(200, { 'Content-Type': 'text/plain' });
          next();
    });

leaderRouter.route('/:leaderId')
    .all(function(req,res,next) {
          res.writeHead(200, { 'Content-Type': 'text/plain' });
          next();
    });

module.exports = leaderRouter;
