const express = require('express');
const expressValidator = require('express-validator');

const routes  = require('./routes/index');
const mongo = require('./services/mongo');

mongo.init();

var app = express();
app.listen(process.env.PORT);
app.use(expressValidator());
app.use('/', routes);
