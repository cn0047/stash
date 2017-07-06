var express = require('express');
var morgan = require('morgan');

var dishRouter = require('./routes/dishRouter.js')();
var leaderRouter = require('./routes/leaderRouter.js');
var promoRouter = require('./routes/promoRouter.js');

var hostname = 'localhost';
var port = 3000;

var app = express();

app.use(morgan('dev'));

app.use('/dishes', dishRouter);
app.use('/leadership', leaderRouter);
app.use('/promotions', promoRouter);

app.use(express.static(__dirname + '/public'));

app.listen(port, hostname, function(){
  console.log(`Server running at http://${hostname}:${port}/`);
});

/*

curl 'http://localhost:3000/dishes'
curl -XPOST http://localhost:3000/dishes \
    -H 'Content-Type: application/json' -d '{"name": "newDish", "description": "newDesc"}'
curl -XDELETE http://localhost:3000/dishes/1
curl -XPUT http://localhost:3000/dishes/1 \
    -H 'Content-Type: application/json' -d '{"dishId": 1, "name": "newDishName", "description": "newDishDesc"}'

curl 'http://localhost:3000/leadership'
curl -XPOST http://localhost:3000/leadership \
    -H 'Content-Type: application/json' -d '{"name": "newDish", "description": "newDesc"}'
curl -XDELETE http://localhost:3000/leadership/1
curl -XPUT http://localhost:3000/leadership/1 \
    -H 'Content-Type: application/json' -d '{"dishId": 1, "name": "newDishName", "description": "newDishDesc"}'

curl 'http://localhost:3000/promotions'
curl -XPOST http://localhost:3000/promotions \
    -H 'Content-Type: application/json' -d '{"name": "newDish", "description": "newDesc"}'
curl -XDELETE http://localhost:3000/promotions/1
curl -XPUT http://localhost:3000/promotions/1 \
    -H 'Content-Type: application/json' -d '{"dishId": 1, "name": "newDishName", "description": "newDishDesc"}'

*/
